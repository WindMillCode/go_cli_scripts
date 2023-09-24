package main

import (
	"github.com/windmillcode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	utils.CDToTestNGApp()

	utils.RunCommand("", []string{})
}
