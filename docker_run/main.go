package main

import (
	"os"

	"github.com/WindMillCode/vscode-extension-libraries/windmillcode-extension-pack-0/task_files/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	workspaceRoot, err := os.Getwd()
	if err != nil {
		return
	}
	settings, err := utils.GetSettingsJSON(workspaceRoot)
	if err != nil {
		return
	}
	dockerContainerName := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"the name of the container"},
			ErrMsg:  "you must provide a container to run",
			Default: settings.ExtensionPack.SQLDockerContainerName,
		},
	)

	utils.RunCommand("docker", []string{"start", dockerContainerName})
}
