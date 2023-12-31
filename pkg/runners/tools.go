package runners

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/erikgeiser/promptkit/selection"
	"github.com/groundsec/secbutler/pkg/data"
	"github.com/groundsec/secbutler/pkg/logger"
	"github.com/groundsec/secbutler/pkg/utils"
	"golang.org/x/exp/slices"
)

var checkToolFunc = `
function check_tool {
  if ! command -v $1 &> /dev/null; then
    error "'$1' not installed (or not in PATH)"
  else
    success "'$1' correctly installed"
  fi
}

`

var startInstallToolsFunc = `function install_tools {
  bold "Installing tools"`
var endInstallToolsFunc = "\n}"

var endInstallScriptTpl = `
function main {
  banner
  check_requirements
  mkdir -p $HOME/.secbutler/tools && cd $HOME/.secbutler/tools
  install_tools
}

main $@
`

func getToolsGroupedOptions(groupingChoice string) ([]string, error) {
	var groupedOptions []string
	for _, tool := range data.SecTools {
		var matchingEl string
		switch groupingChoice {
		case "Category":
			matchingEl = tool.Category
		case "Subcategory":
			matchingEl = tool.Subcategory
		case "Tool":
			matchingEl = tool.Name
		default:
			return nil, fmt.Errorf("invalid grouping option %s", groupingChoice)
		}
		if !slices.Contains(groupedOptions, matchingEl) {
			groupedOptions = append(groupedOptions, matchingEl)
		}
	}
	return groupedOptions, nil
}

func generateToolsScript(groupingChoice string, chosenGroups []string) {
	requirements := []string{}
	installScript := utils.StartInstallTpl
	installScript = strings.Join([]string{installScript, checkToolFunc, startInstallToolsFunc}, "\n")

	// Get list of tools
	for _, tool := range data.SecTools {
		var matchingEl string
		switch groupingChoice {
		case "Category":
			matchingEl = tool.Category
		case "Subcategory":
			matchingEl = tool.Subcategory
		case "Tool":
			matchingEl = tool.Name
		default:
			logger.Fatalf("Invalid grouping option %s", groupingChoice)
		}

		if slices.Contains(chosenGroups, matchingEl) {
			// Update tpl
			lines := strings.Split(tool.InstallCmd, "\n")
			for i, line := range lines {
				lines[i] = fmt.Sprintf("  %s", strings.TrimSpace(line))
			}
			lines = utils.RemoveEmptyStrings(lines)
			trimmedInstallCmd := strings.Join(lines, "\n")
			commentLine := fmt.Sprintf("\n  # Install %s", tool.Name)
			infoLine := fmt.Sprintf("  info \"Installing %s\"", tool.Name)
			checkOutputLine := fmt.Sprintf("  check_tool \"%s\"", tool.Name)
			installScript = strings.Join([]string{installScript, commentLine, infoLine, trimmedInstallCmd, checkOutputLine}, "\n")
			for _, req := range tool.Requirements {
				formattedReq := fmt.Sprintf("\"%s\"", req)
				if !slices.Contains(requirements, formattedReq) {
					requirements = append(requirements, formattedReq)
				}
			}
		}
	}
	installScript = strings.Join([]string{installScript, "\n\n  success \"Installation completed.\"", endInstallToolsFunc}, "")

	// Add check requirements function
	installScript = strings.Join([]string{installScript, utils.StartCheckRequiremenstFunc}, "\n\n")

	requirementsLine := fmt.Sprintf("  REQUIREMENTS=(%s)", strings.Join(requirements, " "))
	installScript = strings.Join([]string{installScript, requirementsLine, utils.EndCheckRequirementsFunc}, "\n")

	installScript = strings.Join([]string{installScript, endInstallScriptTpl}, "\n")

	// Write the content to the file, creating it or overwriting if it already exists.
	installToolsFilePath := filepath.Join(utils.UserHomeDir(), utils.MainDirName, "install_tools.sh")
	err := os.WriteFile(installToolsFilePath, []byte(installScript), 0644)
	if err != nil {
		logger.Fatal(err)
	}

	// Change the file permissions to make it executable.
	err = os.Chmod(installToolsFilePath, 0755)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info(fmt.Sprintf("Script '%s' correctly created. You can run it to install all the selected tools.", installToolsFilePath))
}

func GenerateToolsInstallScript() {
	// The user can generate the install script by grouping for 'category', 'subcategory' and 'tool'
	groupingOptions := []string{"Category", "Subcategory", "Tool"}

	groupingSelection := selection.New("Select the grouping for the installation script:", groupingOptions)

	groupingChoice, err := groupingSelection.RunPrompt()
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	groupedOptions, err := getToolsGroupedOptions(groupingChoice)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}

	chosenGroups := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select the groups/tools you want:",
		Options: groupedOptions,
	}
	err = survey.AskOne(prompt, &chosenGroups)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	generateToolsScript(groupingChoice, chosenGroups)
}
