package utils

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)


func ReadFile(filePath string) (string, error) {
	// Read the entire content of the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading from file", err)
		return "", err
	}

	// Convert content to a string and return it
	return string(content), nil
}

func OverwriteFile(filePath string, content string) error {
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error reading from file", err)
	}
	return err
}

func FolderExists(path string) bool {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		// path/to/whatever exists
		return true
	} else {
		return false
	}
}

func GetItemsInFolder(folderPath string, itemsToRetrieve []string) ([]string, error) {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	var filenames []string
	for _, file := range files {
		// if !file.IsDir() {
		filenames = append(filenames, file.Name())
		// }
	}

	return filenames, nil
}

func GetItemsInFolderRecursive(folderPath string, recursive bool) ([]string, error) {
	var filenames []string

	err := filepath.WalkDir(folderPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !recursive && path != folderPath {
			return filepath.SkipDir // Skip subdirectories when not in recursive mode
		}

		filenames = append(filenames, path)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return filenames, nil
}

func HasSuffixInArray(str string, suffixes []string,removeSuffix bool) string {
	for _, suffix := range suffixes {
		if strings.HasSuffix(str, suffix) {
			if removeSuffix == true {

				return strings.TrimSuffix(str, suffix)
			} else{
				return str
			}
		}
	}
	return ""
}

func HasPrefixInArray(str string, prefixes []string,removeSuffix bool) string {
	for _, prefix := range prefixes {
		if strings.HasPrefix(str, prefix) {
			if removeSuffix == true {

				return strings.TrimPrefix(str, prefix)
			} else{
				return str
			}
		}
	}
	return ""
}


func RemoveDrivePath(folderPath string) (string) {

	folderPath = filepath.ToSlash(folderPath)
	parts := strings.Split(folderPath, "/")
	if len(parts) >= 2 && strings.HasSuffix(parts[0], ":") {

			parts = parts[1:]
	}
	resultPath := filepath.Join(parts...)
	resultPath = filepath.FromSlash(resultPath)
	return resultPath
}

func IsFileOrFolder(path string) (string, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
			return "", err
	}

	if fileInfo.IsDir() {
			return "dir", nil
	}

	return "file", nil
}

func ConvertPathToOSFormat(inputPath string) string {
	return filepath.FromSlash(inputPath)
}

func JoinAndConvertPathToOSFormat(inputPathParts ...string) string {
	inputPath := filepath.Join(inputPathParts...)
	return ConvertPathToOSFormat(inputPath)
}
