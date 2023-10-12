package main

import (
	"github.com/windmillcode/go_scripts/v2/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	utils.CDToTestNGApp()

	utils.RunCommand("", []string{})
}
