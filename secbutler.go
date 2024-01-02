package main

import (
	"github.com/groundsec/secbutler/cmd"
	"github.com/groundsec/secbutler/pkg/utils"
)

func main() {
	utils.Banner(utils.Version)
	cmd.Execute()
}
