package runners

import (
	"fmt"
	"strings"

	"github.com/groundsec/gogetfp"
	"github.com/groundsec/secbutler/pkg/logger"
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

	fmt.Printf(" 📝  Untested\t\t%t\n", untestedFlag)
	fmt.Printf(" 🎲  Random\t\t%t\n", randomFlag)
	fmt.Printf(" 🔒  HTTPS\t\t%t\n", httpsFlag)
	fmt.Printf(" 🌟  Elite\t\t%t\n", eliteFlag)
	fmt.Printf(" 🎭  Anonymous\t\t%t\n", anonFlag)
	fmt.Printf(" 🔍  Google\t\t%t\n", googleFlag)
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
