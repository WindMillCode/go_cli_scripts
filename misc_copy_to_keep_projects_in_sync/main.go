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

	projectSrcCLIString := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"Place the path to the src project to copy the files/dirs from"},
			Default: workspaceRoot,
		},
	)

	projectsDestCLIString := utils.TakeVariableArgs(
		utils.TakeVariableArgsStruct{
			Prompt: "Provide the paths of all the projects to copy the files/dirs",
			ErrMsg: "Projects must be provided",
		},
	)

	target := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"provide the path to the files/dirs (IMPT: the path to the target in the source should be same as dest to avoid unexpected consequences)"},
			Default: "",
		},
	)

	sourcePath := filepath.Join(projectSrcCLIString, target)
	regex0 := regexp.MustCompile(" ")
	projectsList := regex0.Split(projectsDestCLIString, -1)

	var wg sync.WaitGroup
	for _, project := range projectsList {
		destPath := filepath.Join(project, target)

		wg.Add(1)
		go func() {
			defer wg.Done()
			fileInfo, err := os.Stat(sourcePath)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			if fileInfo.IsDir() {
				utils.CopyDir(sourcePath, destPath)
			} else {
				utils.CopyFile(sourcePath, destPath)
			}

		}()
	}
	wg.Wait()

}
