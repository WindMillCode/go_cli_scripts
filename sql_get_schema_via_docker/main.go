package main

import (
	"fmt"
	"os"

	"github.com/windmillcode/go_scripts/utils"
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
			Prompt:  []string{"the name of the docker container"},
			ErrMsg:  "a docker container must be provided",
			Default: settings.ExtensionPack.SQLDockerContainerName,
		},
	)

	cliInfo := utils.ShowMenuModel{
		Prompt:  "select from one of the following databases",
		Choices: settings.ExtensionPack.DatabaseOptions,
	}
	databaseSoftwareName := utils.ShowMenu(cliInfo, nil)
	if databaseSoftwareName == "mysql" {
		mysqlUsername := utils.GetInputFromStdin(
			utils.GetInputFromStdinStruct{
				Prompt:  []string{"enter the mysql username"},
				Default: "myadmin",
			},
		)
		mysqlPass := utils.GetInputFromStdin(
			utils.GetInputFromStdinStruct{
				Prompt:  []string{"enter the mysql password"},
				Default: "my-secret-pw",
			},
		)
		databaseName := utils.GetInputFromStdin(
			utils.GetInputFromStdinStruct{
				Prompt:  []string{"provide the database name"},
				Default: settings.ExtensionPack.DatabaseName,
			},
		)

		utils.RunCommand("docker", []string{
			"exec",
			"--workdir",
			"/root",
			dockerContainerName,
			"mysqldump",
			"-u",
			mysqlUsername,
			"--password=" + mysqlPass,
			"--single-transaction",
			"--no-data",
			"--no-create-db",
			databaseName,
			">",
			"backup.sql",
		})
	} else {
		fmt.Printf(fmt.Sprintf("%s is not supported as of right now"), databaseSoftwareName)
	}

}
