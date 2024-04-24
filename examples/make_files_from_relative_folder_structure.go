package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/windmillcode/go_cli_scripts/v5/utils"
)

func main() {

	folderPaths := []string{
		utils.ConvertPathToOSFormat("C:\\Users\\Restop-1294\\My_Apps\\Windmillcode_app_tutorials\\tutorials\\refactor_flutter_translate\\flutter_translate\\lib"),
	}

	files := []string{}
	for _, folderPath := range folderPaths {
		dirParams := utils.TraverseDirectoryParams{
			RootDir: folderPath, // Specify your directory here
			Predicate: func(path string, info os.FileInfo) {
				// Action to perform on each .dart file that is not a _test.dart or g.dart file
				// fmt.Println("Found dart file:", path)
				files = append(files, path)
			},
			Filter: func(path string, info os.FileInfo) bool {
				return strings.HasSuffix(path, ".dart") && !strings.HasSuffix(path, "_test.dart") && !strings.Contains(path, "g.dart")
			},
		}

		err := utils.TraverseDirectory(dirParams)
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}
	}

	// currentDir,_ := os.Getwd()
	// rootFolder := utils.JoinAndConvertPathToOSFormat(currentDir,".","output")
	rootFolder := utils.ConvertPathToOSFormat("C:\\Users\\Restop-1294\\My_Apps\\Windmillcode_app_tutorials\\tutorials\\refactor_flutter_translate\\flutter_translate\\test")
	for _, myFilePath := range files {
		fileInfo, _ := os.Stat(myFilePath)
		if !fileInfo.IsDir() {

			suffixedPath := utils.RemovePathPrefix(myFilePath, folderPaths)

			newFilePath := utils.JoinAndConvertPathToOSFormat(rootFolder, suffixedPath)

			newFilePathDir := filepath.Dir(newFilePath)
			newFilePathBase := filepath.Base(newFilePath)
			newFilePathBase = strings.ReplaceAll(newFilePathBase, ".dart", "_test.dart")
			newFilePath = utils.JoinAndConvertPathToOSFormat(newFilePathDir, newFilePathBase)
			newFile, err := utils.EnsureDirAndCreateFile(newFilePath)
			if err != nil {
				fmt.Println(err)
				// return nil, err
			}
			fmt.Println(newFile.Name())
		}
	}

}
