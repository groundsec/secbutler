package runners

import (
	"fmt"
	"strings"
)

func revshellFormatter(shellTpl string, host string, port int) string {
	reverseShell := strings.ReplaceAll(shellTpl, "{{HOST}}", host)
	reverseShell = strings.ReplaceAll(reverseShell, "{{PORT}}", fmt.Sprint(port))
	return reverseShell
}

func GetReverseShell() {

}
