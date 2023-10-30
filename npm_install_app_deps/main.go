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
			Prompt:  "Provide the paths of all the application where you want the actions to take place",
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

	cliInfo = utils.ShowMenuModel{
		Prompt:  "reinstall?",
		Choices: []string{"true", "false"},
	}
	reinstall := utils.ShowMenu(cliInfo, nil)

	cliInfo = utils.ShowMenuModel{
		Prompt:  "force?",
		Choices: []string{"true", "false"},
	}
	force := utils.ShowMenu(cliInfo, nil)

	var wg sync.WaitGroup
	regex0 := regexp.MustCompile(" ")
	projectsList := regex0.Split(projectsCLIString, -1)
	for _, project := range projectsList {
		app := filepath.Join(project, appLocation)

		wg.Add(1)
		go func() {
			defer wg.Done()
			if reinstall == "true" {
				os.Remove(utils.JoinAndConvertPathToOSFormat(app,"package-lock.json"))
				os.Remove(utils.JoinAndConvertPathToOSFormat(app,"yarn.lock"))
				if packageManager == "yarn" {
					utils.RunCommandInSpecificDirectory(packageManager, []string{"cache", "clean"}, app)
				}
				os.RemoveAll(utils.JoinAndConvertPathToOSFormat(app,"node_modules"))
			}
			if packageManager == "npm" {
				command := []string{"install", "-s"}
				if force == "true" {
					command = append(command, "--force")
				}
				command = append(command, "--verbose")
				utils.RunCommandInSpecificDirectory(packageManager, command, app)
			} else {
				command := []string{"install"}
				if force == "true" {
					command = append(command, "--force")
				}
				command = append(command, "--verbose")
				utils.RunCommandInSpecificDirectory(packageManager, command, app)
			}
		}()
	}
	wg.Wait()

}
