package main

import (

	"github.com/windmillcode/go_cli_scripts/v3/utils"
)

func main() {

	utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"Cancel porgram on ctrl+c"},
			Default: "",
		},
	)
	// fmt.Println(myVal)

	cliInfo := utils.ShowMenuModel{
		Prompt:  "Cancel on ctrl +C on choose option ",
		Choices: []string{"A", "B", "C"},
	}
	utils.ShowMenu(cliInfo, nil)


	cliInfo = utils.ShowMenuModel{
		Prompt:  "Progra, is continuing ",
		Choices: []string{"A", "B", "C"},
	}
	utils.ShowMenu(cliInfo, nil)
}
