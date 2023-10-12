package main

import (
	"fmt"
	"path/filepath"

	"github.com/windmillcode/go_scripts/v2/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	cliInfo := utils.ShowMenuModel{
		Prompt: "Choose an option:",
		Choices: []string{
			"docs\\tasks_docs",
			"docs\\application_documentation",
			"issues",
		},
	}
	docLocation := utils.ShowMenu(cliInfo, nil)
	docLocation = filepath.Join(docLocation)
	entityNames, err := utils.GetItemsInFolder(docLocation)
	if err != nil {

		fmt.Println("Error retrieving file names please check the spelling of the provided/selected folder")
	}
	cliInfo = utils.ShowMenuModel{
		Prompt:  "Select the entity to open",
		Choices: entityNames,
		Other:   true,
	}
	targetName := utils.ShowMenu(cliInfo, nil)
	targetPath := filepath.Join(docLocation, targetName)
	utils.RunCommand("code", []string{targetPath})
}
