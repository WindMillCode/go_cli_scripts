package main

import (
	"fmt"
	"path/filepath"

	"github.com/windmillcode/go_scripts/v2/utils"
)

func main() {

	utils.CDToWorkspaceRoot()

	initScript := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"docker init script to run relative to workspace root "},
			Default: filepath.Join("ignore\\Local\\docker_init_container.go"),
		},
	)
	initScriptArgs := utils.TakeVariableArgs(
		utils.TakeVariableArgsStruct{},
	)
	initScriptArgs = fmt.Sprintf("%s %s", filepath.Join("..", "..", ".."), initScriptArgs)
	initScriptLocation := filepath.Dir(initScript)
	utils.CDToLocation(initScriptLocation)
	initScript = filepath.Base(initScript)

	utils.RunCommand("windmillcode_go", []string{"run", initScript, initScriptArgs})
}
