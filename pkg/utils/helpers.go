package utils

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/groundsec/secbutler/pkg/logger"
	"github.com/groundsec/secbutler/pkg/types"
)

func BooleanColorCode(boolValue bool) string {
	if boolValue {
		return ANSICodes["Green"]
	} else {
		return ANSICodes["Red"]
	}
}

func UserHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logger.Fatal(err)
	}
	return homeDir
}

func CheckAndCreateSecbutlerDir() {
	homeDir := UserHomeDir()
	secbutlerDir := filepath.Join(homeDir, MainDirName)

	// Define SecButler dirs
	payloadsDir := filepath.Join(homeDir, MainDirName, PayloadsDirName)
	toolsDir := filepath.Join(homeDir, MainDirName, ToolsDirName)
	cheatsheetsDir := filepath.Join(homeDir, MainDirName, CheatsheetsDirName)

	secbutlerDirs := []types.SecButlerDir{
		{Name: PayloadsDirName, Path: payloadsDir},
		{Name: ToolsDirName, Path: toolsDir},
		{Name: CheatsheetsDirName, Path: cheatsheetsDir},
	}

	// Creating main dir
	if _, err := os.Stat(secbutlerDir); os.IsNotExist(err) {
		logger.Info(fmt.Sprintf("Creating %s/%s directory...", homeDir, MainDirName))
		if err := os.Mkdir(secbutlerDir, 0700); err != nil {
			logger.Info(fmt.Sprintf("Failed to create %s/%s directory", homeDir, MainDirName))
			os.Exit(1)
		}
	}

	for _, d := range secbutlerDirs {
		if _, err := os.Stat(d.Path); os.IsNotExist(err) {
			logger.Info(fmt.Sprintf("Creating %s/%s/%s directory...", homeDir, MainDirName, d.Name))
			if err := os.Mkdir(d.Path, 0700); err != nil {
				logger.Info(fmt.Sprintf("Failed to create %s/%s/%s directory", homeDir, MainDirName, d.Name))
				os.Exit(1)
			}
		}
	}
}

func IsCommandAvailable(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func DownloadFile(url string, filepath string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func GetCurrentIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

func RemoveEmptyStrings(stringArr []string) []string {
	nonEmptyLines := []string{}
	for _, line := range stringArr {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}
	return nonEmptyLines
}

func GetCurrentShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "unknown" // Default value if SHELL variable is not set
	}
	return shell
}

func GetShellRc() string {
	shell := GetCurrentShell()
	switch shell {
	case "/bin/bash":
		return "~/.bashrc"
	case "/bin/zsh":
		return "~/.zshrc"
	default:
		return "~/.bashrc"
	}
}

func PrettyPrintMarkdownFile(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		logger.Fatalf("failed to read file: %v", err)
	}

	// Render Markdown to ANSI using glamour
	renderer, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"), // or "light"
	)
	if err != nil {
		logger.Fatalf("failed to create renderer: %v", err)
	}

	out, err := renderer.Render(string(content))
	if err != nil {
		logger.Fatalf("failed to render markdown: %v", err)
	}

	// Output the pretty-printed Markdown
	os.Stdout.WriteString(out)
}
