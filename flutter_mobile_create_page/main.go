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
	pageName := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"The name you would like to give to the page"},
			ErrMsg: "You must provide a value",
		},
	)
	entityName := pageName + "Page"
	snakeCasePageName := strcase.ToSnake(pageName)
	providerLocation := filepath.Join(flutterApp, "lib", "pages", snakeCasePageName)
	newTemplatePath := filepath.Join(providerLocation, fmt.Sprintf("%s.dart", snakeCasePageName))
	newRiverPodProviderPath := filepath.Join(providerLocation, fmt.Sprintf("%s_riverpod_provider.dart", snakeCasePageName))
	utils.CopyDir(templateLocation, providerLocation)
	os.Rename(
		filepath.Join(providerLocation, "template_page.dart"),
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
		fileString = strings.ReplaceAll(fileString, "WMLTemplate", strcase.ToCamel(entityName))
		fileString = strings.ReplaceAll(fileString, "Wml", "WML")
		fileString = strings.ReplaceAll(fileString, "template", snakeCasePageName)
		utils.OverwriteFile(path, fileString)
	}

}
