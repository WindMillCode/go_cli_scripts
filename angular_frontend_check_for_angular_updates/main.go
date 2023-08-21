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
	angularAppLocation := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"provide the relative path to the angular app (note : for every project  the relative path should be the same other wi)"},
			Default: "apps/frontend/AngularApp",
		},
	)

	var wg sync.WaitGroup
	regex0 := regexp.MustCompile(" ")
	projectsList  := regex0.Split(projectsCLIString, -1)
	for _,project := range projectsList{
		app := filepath.Join(project,angularAppLocation)
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			utils.RunCommandInSpecificDirectory("npx", []string{"ng", "update"},app)
		}()
	}
	wg.Wait()
}
