package main

import (
	"github.com/WindMillCode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	utils.CDToAngularApp()

	utils.RunCommand("yarn", []string{"compodoc:build-and-serve"})
}
