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
	utils.CDToWorkspaceRoot()
	utils.CDToFlutterApp()
	flutterApp, err := os.Getwd()
	if err != nil {
		return
	}
	targetName := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"The name you would like to give to the widget"},
			ErrMsg: "You must provide a value",
		},
	)
	snakeCaseName := strcase.ToSnake(targetName)
	snakeCaseWidgetName := strcase.ToSnake(targetName + "Widget")
	providerLocation := filepath.Join(flutterApp, "lib", "pages", snakeCaseName)
	newTemplatePath := filepath.Join(providerLocation, fmt.Sprintf("%s.dart", snakeCaseWidgetName))
	newRiverPodProviderPath := filepath.Join(providerLocation, fmt.Sprintf("%s_riverpod_provider.dart", snakeCaseWidgetName))
	utils.CopyDir(templateLocation, providerLocation)
	os.Rename(
		filepath.Join(providerLocation, "template_widget.dart"),
		newTemplatePath,
	)
	os.Rename(
		filepath.Join(providerLocation, "template_riverpod_provider.dart"),
		newRiverPodProviderPath,
	)

	for _, path := range []string{newTemplatePath, newRiverPodProviderPath} {
		fileString, err := utils.ReadFile(path)
		if err != nil {
			return
		}
		fileString = strings.ReplaceAll(fileString, "WMLTemplate", strcase.ToCamel(targetName))
		fileString = strings.ReplaceAll(fileString, "Wml", "WML")
		fileString = strings.ReplaceAll(fileString, "template", snakeCaseWidgetName)
		utils.OverwriteFile(path, fileString)
	}

}
