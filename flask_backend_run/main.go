package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/windmillcode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	workspaceFolder, err := os.Getwd()
	if err != nil {
		return
	}
	utils.CDToFlaskApp()
	flaskAppFolder, err := os.Getwd()
	if err != nil {
		return
	}
	envVarsFile := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"where are the env vars located"},
			Default: filepath.Join(workspaceFolder, ".\\ignore\\Local\\flask_backend_shared.go"),
		},
	)
	pythonVersion := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"provide a python version for pyenv to use"},
			Default: "3.11.4",
		},
	)
	if pythonVersion != "" {
		utils.RunCommand("pyenv", []string{"shell", pythonVersion})
	}
	for {
		utils.CDToLocation(workspaceFolder)
		envVars := utils.RunCommandAndGetOutput("windmillcode_go", []string{"run", envVarsFile, filepath.Dir(envVarsFile)})
		envVarsArray := strings.Split(envVars, ",")
		for _, x := range envVarsArray {
			keyPair := []string{}
			for _, y := range strings.Split(x, "=") {
				keyPair = append(keyPair, strings.TrimSpace(y))
			}
			os.Setenv(keyPair[0], keyPair[1])
		}
		utils.CDToLocation(flaskAppFolder)
		utils.RunCommand("python", []string{"app.py"})

	}

}
