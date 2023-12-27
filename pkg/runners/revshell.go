package runners

import (
	"fmt"
	"strings"

	"github.com/erikgeiser/promptkit/selection"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/groundsec/secbutler/pkg/data"
	"github.com/groundsec/secbutler/pkg/logger"
	"golang.org/x/exp/slices"
)

func revshellFormatter(revshellName string, host string, port string) {
	for _, revshell := range data.RevShells {
		if revshell.Name == revshellName {
			shellTpl := revshell.ShellTpl

			reverseShell := strings.ReplaceAll(shellTpl, "{{HOST}}", host)
			reverseShell = strings.ReplaceAll(reverseShell, "{{PORT}}", port)
			fmt.Printf("%s Reverse Shell:\n\n%s\n\n", revshellName, reverseShell)
			return
		}
	}

	logger.Error("Unable to find selected reverse shell")
}

func GetReverseShell() {
	osSelection := selection.New("Select the OS for the reverse shell:", []string{"Linux", "Windows", "Mac"})

	osChoice, err := osSelection.RunPrompt()
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}

	// Get the reverse shells for the chosen OS
	var revshellOptions []string
	for _, revshell := range data.RevShells {
		if slices.Contains(revshell.OS, osChoice) {
			revshellOptions = append(revshellOptions, revshell.Name)
		}
	}

	revshellSelection := selection.New("Select the reverse shell you want:", revshellOptions)

	revshellChoice, err := revshellSelection.RunPrompt()
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}

	// Select Host
	hostInput := textinput.New("Type the HOST value:")
	hostInput.InitialValue = "10.10.10.10"
	hostInput.Placeholder = "The HOST field cannot be empty"

	hostValue, err := hostInput.RunPrompt()
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}

	// Select Port
	portInput := textinput.New("Type the PORT value:")
	portInput.InitialValue = "9001"
	portInput.Placeholder = "The PORT field cannot be empty"

	portValue, err := portInput.RunPrompt()
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}

	revshellFormatter(revshellChoice, hostValue, portValue)
}
