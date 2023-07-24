package main

import (
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/windmillcode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
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
	targetName := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"enter the name of the entity PLEASE USE DASHES OR UNDERLINE FOR SPACES"},
		},
	)
	pattern := `\s+`
	matched := regexp.MustCompile(pattern).MatchString(targetName)
	if matched == true {
		fmt.Printf("The document name cannot contain any speaces PLEASE USE DASHES OR UNDERLINE FOR SPACES !!!!!!!!!     :)")
		return
	}
	targetPath := filepath.Join(docLocation, targetName)
	templatePath := filepath.Join(docLocation, "template")
	utils.CopyDir(templatePath, targetPath)
	utils.RunCommand("code", []string{targetPath})
}
