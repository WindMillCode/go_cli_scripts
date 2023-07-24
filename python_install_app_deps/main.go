package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/windmillcode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	// cliInfo := utils.ShowMenuModel{
	// 	Prompt: "Choose an option:",
	// 	Choices:[]string{".\\apps\\backend\\FlaskApp"},
	// }
	// appLocation := utils.ShowMenu(cliInfo,nil)
	// appLocation = filepath.Join(appLocation)
	appLocation := filepath.Join(".\\apps\\backend\\FlaskApp")

	pythonVersion := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"provide a python version for pyenv to use"},
			Default: "3.11.4",
		},
	)
	utils.RunCommand("pyenv", []string{"shell", pythonVersion})
	cliInfo := utils.ShowMenuModel{
		Prompt:  "reinstall?",
		Choices: []string{"true", "false"},
	}
	reinstall := utils.ShowMenu(cliInfo, nil)
	utils.CDToLocation(appLocation)

	switch targetOs := runtime.GOOS; targetOs {
	case "windows":

		sitePackages := filepath.Join(".", "site-packages", "windows")
		if reinstall == "true" {
			if err := os.RemoveAll(sitePackages); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}
		utils.RunCommand("pip", []string{"install", "-r", "requirements.txt", "--target", sitePackages})
	case "linux", "darwin":
		sitePackages := filepath.Join(".", "site-packages", "linux")
		if reinstall == "true" {
			if err := os.RemoveAll(sitePackages); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}
		utils.RunCommand("pip", []string{"install", "-r", "requirements.txt", "--target", sitePackages})

	default:
		fmt.Println("Unknown Operating System:", targetOs)
	}

}
