package main

import (
	"fmt"
	"path/filepath"

	"github.com/WindMillCode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	initScript := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"docker init script to run relative to workspace root "},
			Default: filepath.Join("ignore\\Local\\docker_init_container.go"),
		},
	)
	initScriptArgs := utils.TakeVariableArgs(
		utils.TakeVariableArgsStruct{},
	)
	fmt.Println(fmt.Sprintf("%s %s", initScript, initScriptArgs))

	utils.RunCommand("windmillcode_go", []string{"run", initScript, initScriptArgs})
}
