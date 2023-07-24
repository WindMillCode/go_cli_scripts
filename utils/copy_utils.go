package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyDir(src, dest string) error {
	// Get the file information of the source directory
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	// Create the destination directory with the same permissions as the source
	if err := os.MkdirAll(dest, srcInfo.Mode()); err != nil {
		return err
	}

	// Get a list of all entries (files and subdirectories) in the source directory
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			// If the entry is a subdirectory, recursively copy it
			if err := CopyDir(srcPath, destPath); err != nil {
				return err
			}
		} else {
			// If the entry is a file, copy it to the destination
			if err := CopyFile(srcPath, destPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func CopyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

type CopySelectFilesToDestinationStruct struct {
	SourceFiles  []string
	GlobPattern string //regex
	DestinationDir string
}
func CopySelectFilesToDestination(c CopySelectFilesToDestinationStruct) error {



	// Move files with the specified pattern to the destination directory
	if c.GlobPattern !=""{
		var err error
		c.SourceFiles,err = filepath.Glob(c.GlobPattern)
		if err != nil {
				return err
		}
	}


	for _, file := range c.SourceFiles {
			if err := CopyFile(file, filepath.Join(c.DestinationDir, filepath.Base(file))); err != nil {
					return err
			}
	}

	return nil
}


func main() {
	srcDir := "/path/to/source/directory"
	destDir := "/path/to/destination/directory"

	err := CopyDir(srcDir, destDir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Directory copied successfully.")
}
