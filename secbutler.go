package main

import (
	"github.com/groundsec/secbutler/cmd"
	"github.com/groundsec/secbutler/pkg/utils"
)

var version = "0.1.0"

func main() {
	utils.Banner(version)
	cmd.Execute()
}
