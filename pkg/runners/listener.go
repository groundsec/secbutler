package runners

import (
	"fmt"
	"strings"

	"github.com/erikgeiser/promptkit/selection"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/groundsec/secbutler/pkg/data"
	"github.com/groundsec/secbutler/pkg/logger"
)

func listenerFormatter(listenerName string, port string) {
	for _, listener := range data.RevShellListeners {
		if listener.Name == listenerName {
			listenerTpl := listener.ListenerTpl

			reverseShellListener := strings.ReplaceAll(listenerTpl, "{{PORT}}", port)
			fmt.Printf("%s Reverse Shell Listener:\n\n%s\n\n", listenerName, reverseShellListener)
			return
		}
	}

	logger.Error("Unable to find selected reverse shell listener")
}

func GetReverseShellListener() {

	// Get the reverse shells for the chosen OS
	var listenerOptions []string
	for _, listener := range data.RevShellListeners {
		listenerOptions = append(listenerOptions, listener.Name)

	}

	listenerSelection := selection.New("Select the reverse shell listener you want:", listenerOptions)

	listenerChoice, err := listenerSelection.RunPrompt()
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

	listenerFormatter(listenerChoice, portValue)
}
