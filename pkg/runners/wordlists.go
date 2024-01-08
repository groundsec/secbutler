package runners

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/groundsec/secbutler/pkg/data"
	"github.com/groundsec/secbutler/pkg/logger"
	"github.com/groundsec/secbutler/pkg/utils"
	"golang.org/x/exp/slices"
)

var startInstallWordlistsFunc = `function install_wordlists {
  bold "Installing wordlists"`
var endInstallWordlistsFunc = "\n}"

var endInstallWordlistTpl = `
function main {
  banner
  check_requirements
  mkdir -p /usr/share/wordlists && cd /usr/share/wordlists
  install_wordlists
}

main $@
`

func generateWordlistsScript(chosenWordlists []string) {
	requirements := []string{}
	installWordlist := utils.StartInstallTpl
	installWordlist = strings.Join([]string{installWordlist, startInstallWordlistsFunc}, "\n")

	// Get list of wordlists
	for _, wordlist := range data.Wordlists {
		if slices.Contains(chosenWordlists, wordlist.Name) {

			// Update tpl
			lines := strings.Split(wordlist.InstallCmd, "\n")
			for i, line := range lines {
				lines[i] = fmt.Sprintf("  %s", strings.TrimSpace(line))
			}
			lines = utils.RemoveEmptyStrings(lines)
			trimmedInstallCmd := strings.Join(lines, "\n")
			commentLine := fmt.Sprintf("\n  # Install %s", wordlist.Name)
			infoLine := fmt.Sprintf("  info \"Installing %s\"", wordlist.Name)
			installWordlist = strings.Join([]string{installWordlist, commentLine, infoLine, trimmedInstallCmd}, "\n")
			for _, req := range wordlist.Requirements {
				formattedReq := fmt.Sprintf("\"%s\"", req)
				if !slices.Contains(requirements, formattedReq) {
					requirements = append(requirements, formattedReq)
				}
			}
		}
	}
	installWordlist = strings.Join([]string{installWordlist, "\n\n  success \"Installation completed.\"", endInstallWordlistsFunc}, "")

	// Add the 'WORDLISTS' variable to your current shell
	shellFile := utils.GetShellRc()
	installWordlist = strings.Join([]string{installWordlist, fmt.Sprintf("echo \"export WORDLISTS=\"/usr/share/wordlists\"\" >> %s", shellFile)}, "\n\n")
	installWordlist = strings.Join([]string{installWordlist, fmt.Sprintf("source %s", shellFile)}, "\n\n")

	// Add check requirements function
	installWordlist = strings.Join([]string{installWordlist, utils.StartCheckRequiremenstFunc}, "\n\n")

	requirementsLine := fmt.Sprintf("  REQUIREMENTS=(%s)", strings.Join(requirements, " "))
	installWordlist = strings.Join([]string{installWordlist, requirementsLine, utils.EndCheckRequirementsFunc}, "\n")

	installWordlist = strings.Join([]string{installWordlist, endInstallWordlistTpl}, "\n")

	// Write the content to the file, creating it or overwriting if it already exists.
	installWordlistsFilePath := filepath.Join(utils.UserHomeDir(), utils.MainDirName, "install_wordlists.sh")
	err := os.WriteFile(installWordlistsFilePath, []byte(installWordlist), 0644)
	if err != nil {
		logger.Fatal(err)
	}

	// Change the file permissions to make it executable.
	err = os.Chmod(installWordlistsFilePath, 0755)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info(fmt.Sprintf("Script '%s' correctly created. You can run it to install all the selected wordlists.", installWordlistsFilePath))
}

func GenerateWordlistsInstallScript() {
	var wordlists []string
	for _, w := range data.Wordlists {
		wordlists = append(wordlists, w.Name)
	}

	chosenWordlists := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select the wordlists you want:",
		Options: wordlists,
	}
	err := survey.AskOne(prompt, &chosenWordlists)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	generateWordlistsScript(chosenWordlists)
}
