package main

import (
	"fmt"
	"os"
	// "path/filepath"
	"strings"

	"github.com/windmillcode/go_cli_scripts/v4/utils"
)

func main() {

	folderPaths := []string{
		utils.ConvertPathToOSFormat("C:\\Users\\Restop-1294\\My_Apps\\Windmillcode_app_tutorials\\tutorials\\finding-all-in-one-db-client\\vscode-sqltools"),
	}

	files := []string{}
	for _, folderPath := range folderPaths {
		dirParams := utils.TraverseDirectoryParams{
			RootDir: folderPath, // Specify your directory here
			Predicate: func(path string, info os.FileInfo) {
				// Action to perform on each .dart file that is not a _test.dart or g.dart file
				fmt.Println("Found  file:", path)
				files = append(files, path)
			},
			Filter: func(path string, info os.FileInfo) bool {
				ignorePatterns := []string{

					// "webview",
					"coverage",
					"docs",
					"test",
					".git",
					"dist",
					"node_modules",
					"npm-debug.log",
					"yarn-debug.log",
					"yarn-error.log",
					"pnpm-debug.log",
					".env",
					".env.production",
					".DS_Store",
				}

				suffixes := []string{
					".code-workspace",
					".scss",
					".npmignore",
					"LICENSE",
					".eslintignore",
					".d.ts",
					".vscodeignore",
					".json",
					".gitignore",
					".gif",
					".js",
					".js.map",
					".mdx",
					".test.ts",
					".sql",
					".css",
					".md",
					"extensions.json",
					".svg",
					".lock",
					".png",
					".jpg",
					".ico",
					".sh",
					".yaml",
					".yml",

				}

				// Check if the path ends with .go or .svg (to be ignored)
				for _, suffix := range suffixes {
					if strings.HasSuffix(path, suffix) {
						return false
					}
				}

				// Check against other ignore patterns
				for _, pattern := range ignorePatterns {
					if strings.Contains(path, pattern) {
						return false
					}
				}

				return true
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


			// fmt.Println(myFilePath)
			// fmt.Println(folderPaths[0])
			fileName := strings.Split(myFilePath, folderPaths[0])[1]
			concatenatedContent = append(concatenatedContent, []byte(fmt.Sprintf("\n# FileName: %s \n\n", fileName))...)
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
