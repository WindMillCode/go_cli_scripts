package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(filePath string) (string, error) {
	// Read the entire content of the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading from file",err )
			return "", err
	}

	// Convert content to a string and return it
	return string(content), nil
}

func OverwriteFile(filePath string, content string) error {
	err := ioutil.WriteFile(filePath, []byte(content), 0644);
	if  err != nil{
		fmt.Printf("Error reading from file",err )
	}
	return err
}

func FolderExists(path string) bool {
	_, err := os.Stat(path)
	if ; !os.IsNotExist(err) {
		// path/to/whatever exists
		return true
	} else{
		return false
	}
}

