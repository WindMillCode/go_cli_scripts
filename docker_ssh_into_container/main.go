package main

import (
	"os"

	"github.com/windmillcode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
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
	cliInfo := utils.ShowMenuModel{
		Prompt:  "the command line shell",
		Default: "bash",
		Choices: []string{"sh", "bash", "dash", "zsh", "cmd", "fish", "ksh", "powershell"},
	}
	shell := utils.ShowMenu(cliInfo, nil)

	utils.RunCommand("docker", []string{"exec", "-it", dockerContainerName, shell})
}
