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
			Prompt: "Provide the paths of all the application where you want the actions to take place",
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
		Choices: []string{
			filepath.Join("./apps/frontend/AngularApp"),
			filepath.Join(".\\apps\\cloud\\FirebaseApp"),
		},
	}
	appLocation := utils.ShowMenu(cliInfo, nil)

	cliInfo = utils.ShowMenuModel{
		Prompt:  "reinstall?",
		Choices: []string{"true", "false"},
	}
	reinstall := utils.ShowMenu(cliInfo, nil)


	var wg sync.WaitGroup
	regex0 := regexp.MustCompile(" ")
	projectsList  := regex0.Split(projectsCLIString, -1)
	for _,project := range projectsList{
		app := filepath.Join(project,appLocation)

		wg.Add(1)
		go func(){
			defer wg.Done()
			if reinstall == "true" {
				utils.RunCommandInSpecificDirectory("rm", []string{"package-lock.json"},app)
				utils.RunCommandInSpecificDirectory("rm", []string{"yarn.lock"},app)
				if packageManager == "yarn"  {
					utils.RunCommandInSpecificDirectory(packageManager, []string{"cache", "clean"},app)
				}
				utils.RunCommandInSpecificDirectory("rm", []string{"-r", "node_modules"},app)
			}
			if packageManager == "npm" {
				utils.RunCommandInSpecificDirectory(packageManager, []string{"install","-s"},app)
			} else{
				utils.RunCommandInSpecificDirectory(packageManager, []string{"install"},app)
			}
		}()
	}
	wg.Wait()


}
