package runners

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/erikgeiser/promptkit/selection"
	"github.com/groundsec/secbutler/pkg/data"
	"github.com/groundsec/secbutler/pkg/logger"
	"github.com/groundsec/secbutler/pkg/utils"
	"golang.org/x/exp/slices"
)

var installScriptTpl = `
function banner {
  echo "                   __          __  __"
  echo "   ________  _____/ /_  __  __/ /_/ /__  _____"
  echo "  / ___/ _ \/ ___/ __ \/ / / / __/ / _ \/ ___/"
  echo " (__  )  __/ /__/ /_/ / /_/ / /_/ /  __/ /"    
  echo "/____/\___/\___/_.___/\__,_/\__/_/\___/_/"	
}

function log {

}
`

func getGroupedOptions(groupingChoice string) ([]string, error) {
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

func generateScript(groupingChoice string, chosenGroups []string) {
	startInstallToolsFunc := "function install_tools {"
	endInstallToolsFunc := "\n}"
	requirements := []string{}
	installScript := installScriptTpl
	installScript = strings.Join([]string{installScript, startInstallToolsFunc}, "\n")

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
			comment := fmt.Sprintf("\n  # Install %s", tool.Name)
			installScript = strings.Join([]string{installScript, comment, trimmedInstallCmd}, "\n")
			for _, req := range tool.Requirements {
				if !slices.Contains(requirements, req) {
					requirements = append(requirements, req)
				}
			}
		}
	}
	installScript = strings.Join([]string{installScript, endInstallToolsFunc}, "\n")

	fmt.Println(installScript)
}

func GenerateToolsInstallScript() {
	// The user can generate the install script by grouping for 'category', 'subcategory' and 'tool'
	groupingOptions := []string{"Category", "Subcategory", "Tool"}

	groupingSelection := selection.New("Select the grouping for the installation script:", groupingOptions)

	groupingChoice, err := groupingSelection.RunPrompt()
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	groupedOptions, err := getGroupedOptions(groupingChoice)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}

	chosenGroups := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select the groups/tools you want:",
		Options: groupedOptions,
	}
	survey.AskOne(prompt, &chosenGroups)
	generateScript(groupingChoice, chosenGroups)
}
