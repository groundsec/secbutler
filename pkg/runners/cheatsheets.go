package runners

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/go-git/go-git/v5"
	"github.com/groundsec/secbutler/pkg/data"
	"github.com/groundsec/secbutler/pkg/logger"
	"github.com/groundsec/secbutler/pkg/utils"
)

func checkCheatsheet(folderName string) bool {
	homeDir := utils.UserHomeDir()
	// Define SecButler dirs
	cheatsheetDir := filepath.Join(homeDir, utils.MainDirName, utils.CheatsheetsDirName, folderName)
	if _, err := os.Stat(cheatsheetDir); os.IsNotExist(err) {
		return false
	}
	return true
}

// findReadmeFiles searches for README.md files in a given directory and its subdirectories.
// It returns a slice of strings containing the relative paths to each README.md file.
func findReadmeFiles(rootDir string) ([]string, error) {
	var readmeFiles []string

	// Walk through the directory tree starting at rootDir.
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is named "README.md"
		if !info.IsDir() && strings.ToLower(info.Name()) == "readme.md" {
			// Get the relative path
			relPath, err := filepath.Rel(rootDir, path)
			if err != nil {
				return err
			}
			readmeFiles = append(readmeFiles, relPath)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return readmeFiles, nil
}

func ShowCheatsheet() {
	for _, cheatsheet := range data.Cheatsheets {
		isDownloaded := checkCheatsheet(cheatsheet.Name)
		if !isDownloaded {
			download := false
			prompt := &survey.Confirm{
				Message: fmt.Sprintf("%s is missing, do you want to download it?", cheatsheet.Name),
			}
			survey.AskOne(prompt, &download)
			if !download {
				logger.Error("You cannot continue without downloading the cheatsheet")
				os.Exit(1)
			}
			_, err := git.PlainClone(filepath.Join(utils.UserHomeDir(), utils.MainDirName, utils.CheatsheetsDirName, cheatsheet.Name), false, &git.CloneOptions{
				URL:      cheatsheet.Repository,
				Progress: os.Stdout,
			})
			if err != nil {
				logger.Fatalf("Error: %v", err)
			}
		}
	}

	// Select the cheatsheet
	cheatsheetToUse := data.Cheatsheets[0].Name
	if len(data.Cheatsheets) > 1 {
		var options []string
		for _, c := range data.Cheatsheets {
			options = append(options, c.Name)
		}
		prompt := &survey.Select{
			Message: "Select the cheatsheet:",
			Options: options,
		}
		err := survey.AskOne(prompt, &cheatsheetToUse)
		if err != nil {
			logger.Fatalf("Error: %v", err)
		}
	}

	cheatsheetDir := filepath.Join(utils.UserHomeDir(), utils.MainDirName, utils.CheatsheetsDirName, cheatsheetToUse)
	readmeFiles, err := findReadmeFiles(cheatsheetDir)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}

	var readmeToShow string
	prompt := &survey.Select{
		Message: "Select the content:",
		Options: readmeFiles,
	}
	err = survey.AskOne(prompt, &readmeToShow)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	readmeFullPath := filepath.Join(utils.UserHomeDir(), utils.MainDirName, utils.CheatsheetsDirName, cheatsheetToUse, readmeToShow)
	utils.PrettyPrintMarkdownFile(readmeFullPath)
}
