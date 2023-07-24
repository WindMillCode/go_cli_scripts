package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/windmillcode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	cliInfo := utils.ShowMenuModel{
		Prompt:  "run with cache?",
		Choices: []string{"true", "false"},
	}
	runWithCache := utils.ShowMenu(cliInfo, nil)
	utils.CDToAngularApp()
	if runWithCache == "false" {
		if err := os.RemoveAll(filepath.Join(".", ".angular")); err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	utils.RunCommand("npx", []string{"ng", "serve", "--ssl=true", fmt.Sprintf("--ssl-key=%s", os.Getenv("WML_CERT_KEY0")), fmt.Sprintf("--ssl-cert=%s", os.Getenv("WML_CERT0"))})
}
