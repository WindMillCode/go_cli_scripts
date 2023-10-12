package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/windmillcode/go_scripts/v2/utils"
)

func main() {

	scriptLocation, err := os.Getwd()
	if err != nil {
		return
	}
	templateLocation := filepath.Join(scriptLocation, "template")
	templateEndpointFile := filepath.Join(templateLocation, "template_endpoint.py")
	templateHandlerFile := filepath.Join(templateLocation, "template_handler.py")
	utils.CDToWorkspaceRoot()
	utils.CDToFlaskApp()
	targetApp, err := os.Getwd()
	if err != nil {
		return
	}

	endpointsFolder := filepath.Join(targetApp, "endpoints")
	handlersFolder := filepath.Join(targetApp, "handlers")

	targetName := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"Please provide the name of the controller"},
			ErrMsg: "You must provide the name of the controller",
		},
	)
	urlPrefix := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"provide the url prefix for the controller"},
			Default: targetName,
		},
	)
	snakeCaseUrlPrefix := strcase.ToSnake(urlPrefix)
	snakeCaseTargetName := strcase.ToSnake(targetName)
	snakeCaseEndpointTargetName := strcase.ToSnake(targetName + "_endpoint")
	snakeCaseHandlersTargetName := strcase.ToSnake(targetName + "_handler")
	endpointsFile := filepath.Join(endpointsFolder, fmt.Sprintf("%s.py", snakeCaseEndpointTargetName))
	handlersFile := filepath.Join(handlersFolder, fmt.Sprintf("%s.py", snakeCaseHandlersTargetName))
	utils.CopyFile(templateEndpointFile, endpointsFile)
	utils.CopyFile(templateHandlerFile, handlersFile)

	for _, path := range []string{endpointsFile, handlersFile} {
		fileString, err := utils.ReadFile(path)
		if err != nil {
			return
		}
		fileString = strings.ReplaceAll(fileString, "wml_template_url_prefix", snakeCaseUrlPrefix)
		fileString = strings.ReplaceAll(fileString, "wml_template", snakeCaseTargetName)

		utils.OverwriteFile(path, fileString)
	}

	updateAppFile(targetApp)

}

func updateAppFile(targetApp string) {
	// appFile := filepath.Join(targetApp, "app.py")
}
