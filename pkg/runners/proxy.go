package runners

import (
	"fmt"
	"strings"

	"github.com/groundsec/gogetfp"
	"github.com/groundsec/secbutler/pkg/logger"
	"github.com/groundsec/secbutler/pkg/utils"
)

func GetProxy(untestedFlag bool, randomFlag bool, httpsFlag bool, eliteFlag bool, anonFlag bool, googleFlag bool, countriesFlag string) {
	var countryID []string
	if len(countriesFlag) > 0 {
		countryID = strings.Split(countriesFlag, ",")
	} else {
		countryID = []string{}
	}
	config := gogetfp.FreeProxyConfig{
		Random:    randomFlag,
		HTTPS:     httpsFlag,
		Elite:     eliteFlag,
		Anonym:    anonFlag,
		Google:    googleFlag,
		CountryID: countryID,
	}

	fmt.Println("Proxy config:")
	fmt.Printf(" 📝  Untested\t\t%s%t%s\n", utils.BooleanColorCode(untestedFlag), untestedFlag, utils.ANSICodes["Reset"])
	fmt.Printf(" 🎲  Random\t\t%s%t%s\n", utils.BooleanColorCode(randomFlag), randomFlag, utils.ANSICodes["Reset"])
	fmt.Printf(" 🔒  HTTPS\t\t%s%t%s\n", utils.BooleanColorCode(httpsFlag), httpsFlag, utils.ANSICodes["Reset"])
	fmt.Printf(" 🌟  Elite\t\t%s%t%s\n", utils.BooleanColorCode(eliteFlag), eliteFlag, utils.ANSICodes["Reset"])
	fmt.Printf(" 🎭  Anonymous\t\t%s%t%s\n", utils.BooleanColorCode(anonFlag), anonFlag, utils.ANSICodes["Reset"])
	fmt.Printf(" 🔍  Google\t\t%s%t%s\n", utils.BooleanColorCode(googleFlag), googleFlag, utils.ANSICodes["Reset"])
	fmt.Printf(" 🌍  Countries\t\t%s\n", countryID)
	fmt.Println()

	fp := gogetfp.New(config)
	var proxyFn func() (string, error)

	if untestedFlag {
		proxyFn = fp.GetProxy
	} else {
		proxyFn = fp.GetWorkingProxy
	}

	proxy, err := proxyFn()
	if err != nil {
		logger.Error("Unable to find proxy")
	}
	logger.Info("Proxy correctly found")
	fmt.Println(proxy)
}
