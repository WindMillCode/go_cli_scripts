package main

import (
	"github.com/windmillcode/go_scripts/v2/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	utils.CDToFlutterApp()

	utils.RunCommand("flutter", []string{"channel", "stable"})
	utils.RunCommandAndGetOutput("flutter", []string{"upgrade"})
	// utils.RunCommandAndGetOutput("flutter", []string{"upgrade","--force"})
}
