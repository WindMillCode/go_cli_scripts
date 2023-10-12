package main

import (
	"os"

	"github.com/windmillcode/go_scripts/v2/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	workSpaceFolder, err := os.Getwd()
	if err != nil {
		return
	}
	testArgs := utils.GetTestNGArgs(
		utils.GetTestNGArgsStruct{
			WorkspaceFolder: workSpaceFolder,
		},
	)

	utils.CDToTestNGApp()
	envVarContent, err := utils.ReadFile(testArgs.EnvVarsFile)
	if err != nil {
		return
	}
	err = utils.OverwriteFile(".env", envVarContent)
	if err != nil {
		return
	}
}
