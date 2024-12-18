package main

import (
	"fmt"
	"github.com/windmillcode/go_cli_scripts/v6/utils"
	"os"
	"strings"
)

func main() {

	folderPaths := []string{
		utils.ConvertPathToOSFormat("..\\utils"),
	}

	files := []string{}
	for _, folderPath := range folderPaths {
		dirParams := utils.TraverseDirectoryParams{
			RootDir: folderPath, // Specify your directory here
			Predicate: func(path string, info os.FileInfo) {
				// Action to perform on each .dart file that is not a _test.dart or g.dart file
				fmt.Println("Found dart file:", path)
				files = append(files, path)
			},
			Filter: func(path string, info os.FileInfo) bool {
				return strings.HasSuffix(path, ".go")
			},
		}

		err := utils.TraverseDirectory(dirParams)
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}
	}

	var concatenatedContent []byte

	for _, myFilePath := range files {
		fileInfo, _ := os.Stat(myFilePath)
		if !fileInfo.IsDir() {

			content, err := os.ReadFile(myFilePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}

			fileName := utils.RemovePathPrefix(myFilePath, folderPaths)
			concatenatedContent = append(concatenatedContent, []byte(fmt.Sprintf("# FileName: %s \n\n", fileName))...)
			concatenatedContent = append(concatenatedContent, content...)
		}
	}

	outputFilePath := "./output.md" // Change this to the desired output file path
	err := os.WriteFile(outputFilePath, concatenatedContent, 0644)
	if err != nil {
		fmt.Println("Error writing concatenated content to file:", err)
		return
	}

	fmt.Println("Concatenated content saved to", outputFilePath)
}
