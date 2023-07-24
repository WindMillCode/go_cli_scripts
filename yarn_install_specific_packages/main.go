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

	packageList := utils.TakeVariableArgs(
		utils.TakeVariableArgsStruct{
			Prompt: "Provide the names of the packages you would like to install",
			ErrMsg: "You must provide packages for installation",
		},
	)

	cliInfo = utils.ShowMenuModel{
		Prompt:  "chose whether its a dev dependency (-D) or dependency (-s)",
		Choices: []string{"-D", "-s"},
	}
	depType := utils.ShowMenu(cliInfo, nil)

	cliInfo = utils.ShowMenuModel{
		Prompt:  "reinstall?",
		Choices: []string{"true", "false"},
	}
	reinstall := utils.ShowMenu(cliInfo, nil)

	utils.CDToLocation(appLocation)
	if reinstall == "true" {
		utils.RunCommand("yarn", []string{"remove", packageList})
	}

	utils.RunCommand("yarn", []string{"add", depType, packageList})
}
