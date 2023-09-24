package main

import (
	"github.com/windmillcode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	utils.CDToFlutterApp()

	utils.RunCommand("dart", []string{"fix", "--apply"})
	utils.RunCommand("flutter", []string{"build", "appbundle"})
}
