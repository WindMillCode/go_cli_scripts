package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/windmillcode/go_cli_scripts/v5/utils"
)

func main() {

	// utils.GetInputFromStdin(
	// 	utils.GetInputFromStdinStruct{
	// 		Prompt:  []string{"Cancel porgram on ctrl+c"},
	// 		Default: "",
	// 	},
	// )
	// // fmt.Println(myVal)

	// cliInfo := utils.ShowMenuModel{
	// 	Prompt:  "Cancel on ctrl +C on choose option ",
	// 	Choices: []string{"A", "B", "C"},
	// }
	// utils.ShowMenu(cliInfo, nil)

	// cliInfo = utils.ShowMenuModel{
	// 	Prompt:  "Progra, is continuing ",
	// 	Choices: []string{"A", "B", "C"},
	// }
	// utils.ShowMenu(cliInfo, nil)
	currentDir, _ := os.Getwd()
	folderPath := utils.JoinAndConvertPathToOSFormat(currentDir, "utils")

	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var concatenatedContent []byte

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(folderPath, file.Name())

			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}

			concatenatedContent = append(concatenatedContent, content...)
		}
	}

	outputFilePath := "./output.txt" // Change this to the desired output file path
	err = ioutil.WriteFile(outputFilePath, concatenatedContent, 0644)
	if err != nil {
		fmt.Println("Error writing concatenated content to file:", err)
		return
	}

	fmt.Println("Concatenated content saved to", outputFilePath)
}
