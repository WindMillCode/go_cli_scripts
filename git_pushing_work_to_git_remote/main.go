package main

import (
	"fmt"
	"path/filepath"

	"github.com/windmillcode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	cliInfo := utils.ShowMenuModel{
		Prompt: "choose a location to push to git remote",
		Choices: []string{
			filepath.Join("."),
			filepath.Join(".", "apps", "frontend", "AngularApp"),
			filepath.Join(".", "apps", "backend", "RailsApp"),
			filepath.Join(".", "apps", "backend", "FlaskApp"),
		},
	}
	repoLocation := utils.ShowMenu(cliInfo, nil)
	cliInfo = utils.ShowMenuModel{
		Prompt:  "choose the commit type",
		Choices: []string{"UPDATE", "FIX", "PATCH", "BUG", "MERGE", "COMPLEX MERGE", "CHECKPOINT"},
	}
	commitType := utils.ShowMenu(cliInfo, nil)
	commitType = fmt.Sprintf("[%s]", commitType)
	commitMsg := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"Enter your commit msg: (leave blank to default to 'additional work')"},
			Default: "additional work",
		},
	)
	utils.CDToLocation(repoLocation)

	commitFullMsg := fmt.Sprintf("\"%s %s\"", commitType, commitMsg)
	utils.RunCommand("git", []string{"add", "."})
	utils.RunCommand("git", []string{"commit", "-m", commitFullMsg})
	utils.RunCommand("git", []string{"branch", "--unset-upstream"})
	utils.RunCommand("git", []string{"push", "origin", "HEAD"})
}
