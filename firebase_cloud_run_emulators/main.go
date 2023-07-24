package main

import (
	"github.com/WindMillCode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	utils.CDToFirebaseApp()

	utils.RunCommand("yarn", []string{"cleanup"})
	utils.RunCommand("npx", []string{"firebase", "emulators:start", "--import='devData'", "--export-on-exit"})
}
