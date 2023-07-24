package main

import (
	"fmt"
	"strings"

	"github.com/WindMillCode/vscode-extension-libraries/windmillcode-extension-pack-0/task_files/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	cliInfo := utils.ShowMenuModel{
		Prompt:  "Is it a file or directory:",
		Choices: []string{"file", "directory"},
	}
	targetType := utils.ShowMenu(cliInfo, nil)

	targetName := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"name of the item"},
		},
	)

	var args0 []string
	if targetType == "file" {

		argsString := fmt.Sprintf("rm --cached --sparse %s", targetName)
		args0 = strings.Split(argsString, " ")
	} else {

		argsString := fmt.Sprintf("rm -r --cached --sparse %s", targetName)
		args0 = strings.Split(argsString, " ")
	}

	utils.RunCommand("git", args0)
}
