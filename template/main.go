package main

import (
	"github.com/WindMillCode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	utils.CDToTestNGApp()

	utils.RunCommand("", []string{})
}
