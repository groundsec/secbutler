package main

import (
	"github.com/groundsec/secbutler/cmd"
	"github.com/groundsec/secbutler/pkg/utils"
)

var version = "0.0.1"

func main() {
	utils.Banner(version)
	cmd.Execute()
}
