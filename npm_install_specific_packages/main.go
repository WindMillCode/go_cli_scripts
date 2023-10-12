package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/windmillcode/go_scripts/v2/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	workspaceRoot, err := os.Getwd()
	if err != nil {
		fmt.Println("there was an error while trying to receive the current dir")
	}
	projectsCLIString := utils.TakeVariableArgs(
		utils.TakeVariableArgsStruct{
			Prompt:  "Provide the paths of all the projects where you want the actions to take place",
			Default: workspaceRoot,
		},
	)

	cliInfo := utils.ShowMenuModel{
		Prompt:  "choose the package manager",
		Choices: []string{"npm", "yarn"},
		Default: "npm",
	}
	packageManager := utils.ShowMenu(cliInfo, nil)
	cliInfo = utils.ShowMenuModel{
		Other:  true,
		Prompt: "Choose the node.js app",
		Choices: []string{
			filepath.Join("./apps/frontend/AngularApp"),
			filepath.Join(".\\apps\\cloud\\FirebaseApp"),
			filepath.Join("."),
		},
	}
	appLocation := utils.ShowMenu(cliInfo, nil)

	packagesCLIString := utils.TakeVariableArgs(
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
		Prompt:  "uninstall?",
		Choices: []string{"true", "false"},
	}
	uninstall := utils.ShowMenu(cliInfo, nil)
	cliInfo = utils.ShowMenuModel{
		Prompt:  "install?",
		Choices: []string{"true", "false"},
	}
	install := utils.ShowMenu(cliInfo, nil)

	cliInfo = utils.ShowMenuModel{
		Prompt:  "force",
		Choices: []string{"true", "false"},
	}
	force := utils.ShowMenu(cliInfo, nil)

	var wg sync.WaitGroup
	regex0 := regexp.MustCompile(" ")
	projectsList := regex0.Split(projectsCLIString, -1)

	packagesList := regex0.Split(packagesCLIString, -1)
	for _, project := range projectsList {
		app := filepath.Join(project, appLocation)
		wg.Add(1)
		go func() {
			defer wg.Done()
			uninstallPackages(uninstall, packageManager, force, packagesList, app)

			installPackages(install, packageManager, depType, force, packagesList, app)

		}()

	}
	wg.Wait()

}

func installPackages(
	install string, packageManager string, depType string, force string, packagesList []string, app string) {
	if install == "true" {
		if packageManager == "npm" {
			commands := []string{"install", depType, "--verbose"}
			if force == "true" {
				commands = append(commands, "--force")
			}
			utils.RunCommandInSpecificDirectory(packageManager, append(commands, packagesList...), app)
		} else {
			commands := []string{"add", depType, "--verbose"}
			if force == "true" {
				commands = append(commands, "--force")
			}
			utils.RunCommandInSpecificDirectory(packageManager, append(commands, packagesList...), app)
		}
	}
}

func uninstallPackages(uninstall string, packageManager string, force string, packagesList []string, app string) {
	if uninstall == "true" {
		if packageManager == "npm" {
			commands := []string{"uninstall"}
			if force == "true" {
				commands = append(commands, "--force")
			}
			utils.RunCommandInSpecificDirectory(packageManager, append(commands, packagesList...), app)
		} else {
			commands := []string{"remove"}
			if force == "true" {
				commands = append(commands, "--force")
			}
			utils.RunCommandInSpecificDirectory(packageManager, append([]string{"remove", "-f"}, packagesList...), app)
		}
	}
}
