package main

import (
	"path/filepath"

	"github.com/windmillcode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	cliInfo := utils.ShowMenuModel{
		Prompt:  "Choose the node.js app",
		Choices: []string{filepath.Join("./apps/frontend/AngularApp"), filepath.Join(".\\apps\\cloud\\FirebaseApp")},
	}
	appLocation := utils.ShowMenu(cliInfo, nil)

	cliInfo = utils.ShowMenuModel{
		Prompt:  "reinstall?",
		Choices: []string{"true", "false"},
	}
	reinstall := utils.ShowMenu(cliInfo, nil)
	utils.CDToLocation(appLocation)
	if reinstall == "true" {
		utils.RunCommand("rm", []string{"yarn.lock"})
		utils.RunCommand("rm", []string{"-r", "node_modules"})
		utils.RunCommand("yarn", []string{"cache", "clean"})
	}

	utils.RunCommand("yarn", []string{"install"})
}
