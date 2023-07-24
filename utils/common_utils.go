package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
)



// getType returns the type of a given value as a string
func GetType(value interface{}) string {
	return reflect.TypeOf(value).String()
}


func GetParamValue(parameterName string, parameterValue interface{}) interface{} {
	if parameterValue != nil {
		return parameterValue
	} else {
		fmt.Printf("Parameter '%s' value not found.\n", parameterName)
		return nil
	}
}

func GetCurrentPath() string {
	executablePath, err := os.Executable()
	if err != nil {
		// Handle the error if necessary
		return ""
	}
	return filepath.Dir(executablePath)
}

func GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Remove leading/trailing whitespaces and newline characters
	branch := strings.TrimSpace(string(output))

	return branch, nil
}


// Function to clear the console screen
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
