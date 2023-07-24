package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/windmillcode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	utils.CDToTestNGApp()
	testNGApp, err := os.Getwd()
	if err != nil {
		return
	}
	pageFolder := filepath.Join(testNGApp, "src", "main", "java", "pages")
	testFolder := filepath.Join(testNGApp, "src", "test", "java", "e2e")
	pageName := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"The name of the page on the website type in snake case "},
			ErrMsg: "Your must provide a page name",
		},
	)
	utils.CDToLocation(pageFolder)
	myPrefix := strcase.ToCamel(pageName)
	myDir := strings.ToLower(myPrefix)
	err = os.Mkdir(myDir, 0755)
	if err != nil {
		fmt.Printf("Error: ", err.Error())
	}
	myAct := filepath.Join(pageFolder, myDir, myPrefix+"ActController.java")
	myPage := filepath.Join(pageFolder, myDir, myPrefix+"Page.java")
	myVerify := filepath.Join(pageFolder, myDir, myPrefix+"VerifyController.java")
	utils.CopyFile(filepath.Join(".", "template", "TemplateActController.java"), myAct)
	utils.CopyFile(filepath.Join(".", "template", "TemplatePage.java"), myPage)
	utils.CopyFile(filepath.Join(".", "template", "TemplateVerifyController.java"), myVerify)

	utils.CDToLocation(testFolder)
	myTest := filepath.Join(testFolder, myPrefix+"Test.java")
	utils.CopyFile(filepath.Join(".", "TemplateTest.java"), myTest)
	for _, v := range []string{myAct, myPage, myVerify, myTest} {

		fileString, err := utils.ReadFile(v)
		if err != nil {
			return
		}
		fileString = strings.ReplaceAll(fileString, "Template", myPrefix)
		fileString = strings.ReplaceAll(fileString, "template", myDir)
		utils.OverwriteFile(v, fileString)
	}
}
