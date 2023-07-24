package main

import (
	"github.com/windmillcode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	utils.CDToAngularApp()
	utils.RunCommand("npx", []string{"ng", "update"})
}
