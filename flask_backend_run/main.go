package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/windmillcode/go_scripts/v2/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	workspaceFolder, err := os.Getwd()
	if err != nil {
		fmt.Println("there was an error while trying to receive the current dir")
	}
	settings, err := utils.GetSettingsJSON(workspaceFolder)
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
			Default: filepath.Join(workspaceFolder, settings.ExtensionPack.FlaskBackendDevHelperScript),
		},
	)
	pythonVersion := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"provide a python version for pyenv to use"},
			Default: settings.ExtensionPack.PythonVersion0,
		},
	)
	if pythonVersion != "" {
		utils.RunCommand("pyenv", []string{"shell", pythonVersion})
	}
	for {
		utils.CDToLocation(workspaceFolder)
		envVars := utils.RunCommandAndGetOutput("windmillcode_go", []string{"run", envVarsFile, filepath.Dir(envVarsFile), workspaceFolder})
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
