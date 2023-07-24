package main

import (
	"github.com/WindMillCode/vscode-extension-libraries/windmillcode-extension-pack-0/task_files/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	sourceBranch := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"the branch to merge changes from:"},
			Default: "dev",
		},
	)

	utils.RunCommand("git", []string{"checkout", sourceBranch})
	utils.RunCommand("git", []string{"pull", "origin", sourceBranch})
	utils.RunCommand("git", []string{"checkout", "-"})
	utils.RunCommand("git", []string{"merge", sourceBranch})
}
