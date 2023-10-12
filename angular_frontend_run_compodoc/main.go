package main

import (
	"github.com/windmillcode/go_scripts/v2/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	utils.CDToAngularApp()

	utils.RunCommand("yarn", []string{"compodoc:build-and-serve"})
}
