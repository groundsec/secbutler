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

	"github.com/groundsec/secbutler/pkg/logger"
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
	payloadsDir := filepath.Join(homeDir, MainDirName, PayloadsDirName)

	// Creating main dir
	if _, err := os.Stat(secbutlerDir); os.IsNotExist(err) {
		logger.Info(fmt.Sprintf("Creating %s/%s directory...", homeDir, MainDirName))
		if err := os.Mkdir(secbutlerDir, 0700); err != nil {
			logger.Info(fmt.Sprintf("Failed to create %s/%s directory", homeDir, MainDirName))
			os.Exit(1)
		}
	}

	// Creating payloads dir
	if _, err := os.Stat(payloadsDir); os.IsNotExist(err) {
		logger.Info(fmt.Sprintf("Creating %s/%s/%s directory...", homeDir, MainDirName, PayloadsDirName))
		if err := os.Mkdir(payloadsDir, 0700); err != nil {
			logger.Info(fmt.Sprintf("Failed to create %s/%s/%s directory", homeDir, MainDirName, PayloadsDirName))
			os.Exit(1)
		}
	}

	// Creating tools dir
	if _, err := os.Stat(payloadsDir); os.IsNotExist(err) {
		logger.Info(fmt.Sprintf("Creating %s/%s/%s directory...", homeDir, MainDirName, ToolsDirName))
		if err := os.Mkdir(payloadsDir, 0700); err != nil {
			logger.Info(fmt.Sprintf("Failed to create %s/%s/%s directory", homeDir, MainDirName, ToolsDirName))
			os.Exit(1)
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
