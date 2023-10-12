package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/windmillcode/go_scripts/v2/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	workspaceFolder, err := os.Getwd()
	if err != nil {
		fmt.Println("there was an error while trying to receive the current dir")
	}
	settings, err := utils.GetSettingsJSON(workspaceFolder)
	if err != nil {
		return
	}
	cliInfo := utils.ShowMenuModel{
		Other:   true,
		Prompt:  "Choose an option:",
		Choices: []string{".\\apps\\backend\\FlaskApp"},
	}
	appLocation := utils.ShowMenu(cliInfo, nil)
	appLocation = filepath.Join(appLocation)

	pythonVersion := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"provide a python version for pyenv to use"},
			Default: settings.ExtensionPack.PythonVersion0,
		},
	)
	utils.RunCommand("pyenv", []string{"shell", pythonVersion})

	packageList := utils.TakeVariableArgs(
		utils.TakeVariableArgsStruct{
			Prompt: "Provide the names of the packages you would like to install",
			ErrMsg: "You must provide packages for installation",
		},
	)

	cliInfo = utils.ShowMenuModel{
		Prompt:  "reinstall?",
		Choices: []string{"true", "false"},
	}
	reinstall := utils.ShowMenu(cliInfo, nil)
	utils.CDToLocation(appLocation)
	var sitePackages string
	targetOs := runtime.GOOS
	requirementsFile := targetOs + "-requirements.txt"
	switch targetOs {
	case "windows":

		sitePackages = filepath.Join(".", "site-packages", "windows")

	case "linux", "darwin":
		sitePackages = filepath.Join(".", "site-packages", "linux")

	default:
		fmt.Println("Unknown Operating System:", targetOs)
	}
	if reinstall == "true" {
		utils.RunCommand("pip", []string{"uninstall", packageList})
	}
	utils.RunCommand("pip", []string{"install", packageList, "--target", sitePackages})
	utils.RunCommand("pip", []string{"freeze", "--all", "--path", sitePackages, ">", requirementsFile})

}
