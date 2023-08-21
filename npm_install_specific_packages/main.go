package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/windmillcode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	workspaceRoot,err:= os.Getwd()
	if err !=nil {
		fmt.Println("there was an error while trying to receive the current dir")
	}
	projectsCLIString := utils.TakeVariableArgs(
		utils.TakeVariableArgsStruct{
			Prompt: "Provide the paths of all the projects where you want the actions to take place",
			Default:workspaceRoot,
		},
	)

	cliInfo := utils.ShowMenuModel{
		Prompt: "choose the package manager",
		Choices:[]string{"npm","yarn"},
		Default:"npm",
	}
	packageManager := utils.ShowMenu(cliInfo,nil)
	cliInfo = utils.ShowMenuModel{
		Other: true,
		Prompt:  "Choose the node.js app",
		Choices: []string{filepath.Join("./apps/frontend/AngularApp"), filepath.Join(".\\apps\\cloud\\FirebaseApp")},
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
		Prompt:  "reinstall?",
		Choices: []string{"true", "false"},
	}
	reinstall := utils.ShowMenu(cliInfo, nil)

	var wg sync.WaitGroup
	regex0 := regexp.MustCompile(" ")
	projectsList  := regex0.Split(projectsCLIString, -1)

	packagesList  := regex0.Split(packagesCLIString, -1)
	for _,project := range projectsList{
		app := filepath.Join(project,appLocation)
		wg.Add(1)
		go func(){
			defer wg.Done()
			if reinstall == "true" {
				if packageManager == "npm" {
					utils.RunCommandInSpecificDirectory(packageManager,  append([]string{"uninstall"} ,packagesList...),app)
				} else{
					utils.RunCommandInSpecificDirectory(packageManager, append( []string{"remove"} ,packagesList...),app)
				}
			}

			if packageManager == "npm" {
				utils.RunCommandInSpecificDirectory(packageManager, append( []string{"install",depType,"--verbose"} ,packagesList...),app)
			} else{
				utils.RunCommandInSpecificDirectory(packageManager, append( []string{"add", depType} ,packagesList...),app)
			}
		}()

	}
	wg.Wait()


}
