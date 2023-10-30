package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/windmillcode/go_scripts/v2/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	workspaceRoot, err := os.Getwd()
	settings, err := utils.GetSettingsJSON(workspaceRoot)
	if err != nil {
		return
	}
	angularFrontend := settings.ExtensionPack.AngularFrontend

	cliInfo := utils.ShowMenuModel{
		Prompt:  "run with cache?",
		Choices: []string{"true", "false"},
	}
	runWithCache := utils.ShowMenu(cliInfo, nil)

	cliInfo = utils.ShowMenuModel{
		Prompt: "the configuration to run",
		Choices:angularFrontend.Configurations,
		Default:"development",
		Other:true,
	}
	serveConfiguration := utils.ShowMenu(cliInfo,nil)
	utils.CDToAngularApp()
	if runWithCache == "false" {
		if err := os.RemoveAll(filepath.Join(".", ".angular")); err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	utils.RunCommand("npx", []string{"ng", "serve","-c",serveConfiguration, "--ssl=true", fmt.Sprintf("--ssl-key=%s", os.Getenv("WML_CERT_KEY0")), fmt.Sprintf("--ssl-cert=%s", os.Getenv("WML_CERT0"))})
}
