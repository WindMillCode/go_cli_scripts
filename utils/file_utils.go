package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)


func ReadFile(filePath string) (string, error) {
	// Read the entire content of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading from file %s", err)
		return "", err
	}

	// Convert content to a string and return it
	return string(content), nil
}

func OverwriteFile(filePath string, content string) error {
	err := os.WriteFile(filePath, []byte(content), 0644)
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

func GetItemsInFolder(folderPath string) ([]string, error) {
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

func HasSuffixInArray(str string, suffixes []string, removeSuffix bool) string {
	for _, suffix := range suffixes {
		if strings.HasSuffix(str, suffix) {
			if removeSuffix == true {

				return strings.TrimSuffix(str, suffix)
			} else {
				return str
			}
		}
	}
	return ""
}

func HasPrefixInArray(str string, prefixes []string, removeSuffix bool) string {
	for _, prefix := range prefixes {
		if strings.HasPrefix(str, prefix) {
			if removeSuffix == true {

				return strings.TrimPrefix(str, prefix)
			} else {
				return str
			}
		}
	}
	return ""
}

func RemoveDrivePath(folderPath string) string {

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

func ProcessFilesMatchingPattern(directory, pattern string, predicateFn func(string)) error {
	// Compile the regular expression pattern
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	// Walk the directory and apply the predicate function to matching files
	err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && regex.MatchString(info.Name()) {
			// Apply the predicate function to the full path of the matching file
			// fmt.Println(path)
			predicateFn(path)
		}

		return nil
	})

	return err
}

func ProcessFoldersMatchingPattern(directory, pattern string, predicateFn func(string)) error {
	// Compile the regular expression pattern
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	// Walk the directory and apply the predicate function to matching files
	err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && regex.MatchString(info.Name()) {
			// Apply the predicate function to the full path of the matching file
			// fmt.Println(path)
			predicateFn(path)
		}

		return nil
	})

	return err
}

func AddContentToFile(filePath, valueToAdd string,positon string) error {
	// Read the original file content
	originalContent, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Prepend the prefix to the content
	var newContent []byte
	if positon == "prefix"{
		newContent = []byte(valueToAdd + string(originalContent))
	} else {
		newContent = []byte(  string(originalContent)+valueToAdd)
	}

	// Write the modified content back to the file
	err = os.WriteFile(filePath, newContent, os.ModePerm)
	if err != nil {
		return err
	}

	fmt.Printf("Modified file: %s\n", filePath)
	return nil
}

func AddContentToEachLineInFile(filePath string, predicate func(string) string) error {
	// Open the file for reading and writing
	file, err := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a temporary file to store modified content
	tempFile, err := os.CreateTemp("", "tempfile")
	if err != nil {
		return err
	}
	defer tempFile.Close()

	// Create a scanner to read from the original file
	scanner := bufio.NewScanner(file)

	// Create a writer for the temporary file
	writer := bufio.NewWriter(tempFile)

	for scanner.Scan() {
		line := scanner.Text()
		newLine := predicate(line)

		// Write the modified line to the temporary file
		_, err := writer.WriteString(newLine + "\n")
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Flush any remaining data to the temporary file
	if err := writer.Flush(); err != nil {
		return err
	}

	// Close both files before replacing the original with the temporary file
	file.Close()
	tempFile.Close()

	// Replace the original file with the temporary file
	err = os.Rename(tempFile.Name(), filePath)
	if err != nil {
		return err
	}

	return nil
}

func MergeDirectories(sourceDir, targetDir string, overwrite bool) error {
	return filepath.Walk(sourceDir, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(sourceDir, srcPath)
		if err != nil {
			return err
		}

		destPath := filepath.Join(targetDir, relPath)

		if info.IsDir() {
			if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
				return err
			}
		} else {
			_, err := os.Stat(destPath)
			if err == nil && !overwrite {
				return nil
			}

			if err := CopyFile(srcPath, destPath); err != nil {
				return err
			}
		}

		return nil
	})
}

type TraverseDirectoryParams struct {
	RootDir   string
	Predicate func(string, os.FileInfo)
	Filter    func(string, os.FileInfo) bool
}

func TraverseDirectory(config TraverseDirectoryParams) error {
	return filepath.Walk(config.RootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // Return any error that occurs during traversal
		}

		// Apply the filter function if provided
		if config.Filter != nil && !config.Filter(path, info) {
			return nil // Skip entries that don't match the filter
		}

		config.Predicate(path, info)
		return nil
	})
}

func DownloadFile(url, localPath string) error {
	outFile, err := os.Create(localPath)
	if err != nil {
			return fmt.Errorf("error creating file: %v", err)
	}
	defer outFile.Close()

	response, err := http.Get(url)
	if err != nil {
			return fmt.Errorf("error making GET request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
			return fmt.Errorf("bad status: %s", response.Status)
	}

	_, err = io.Copy(outFile, response.Body)
	if err != nil {
			return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

func ExtractArchive(archiveURL string,removeArchiveFile bool) string {
	// Get the filename from the URL
	segments := strings.Split(archiveURL, "/")
	filename := segments[len(segments)-1]

	// Get the current working directory
	sourceDir, err := GetSourceFilePath()
	if err != nil {
			return fmt.Sprintf("error getting source file directory: %v", err)
	}

	// Construct the full path for the archive
	archivePath := JoinAndConvertPathToOSFormat(sourceDir, filename)

	// Check if the file exists locally, if not, download it
	if _, err := os.Stat(archivePath); os.IsNotExist(err) {
			fmt.Printf("File not found locally. Downloading from %s\n", archiveURL)
			if err := DownloadFile(archiveURL, archivePath); err != nil {
					return fmt.Sprintf("error downloading file: %v", err)
			}
	}

	// Extract the archive using 7z
	sevenZCommandOptions :=CommandOptions{
		Command: "7z",
		Args: []string{"x",archivePath,"-aoa"},
		TargetDir: filepath.Dir(archivePath),
	}
	RunCommandWithOptions(sevenZCommandOptions)

	fmt.Printf("Archive extracted successfully: %s\n", archivePath)
	if removeArchiveFile == true{
		if err := os.Remove(archivePath); err != nil {
			return fmt.Sprintf("error removing archive file: %v", err)
	}

	fmt.Println("Archive file deleted successfully")
	}

	return filepath.Dir(archivePath)
}

func GetSourceFilePath() (string, error) {
	executable, err := os.Executable()
	if err != nil {
			return "", fmt.Errorf("unable to get this programs executable path")
	}
	return filepath.Dir(executable), nil
}

func FindExecutable(executablePrefix, searchDir string) string {
	var executablePath string
	var found bool

	// Define the executable name pattern based on the OS
	executablePattern := executablePrefix
	if runtime.GOOS == "windows" {
			executablePattern += ".exe"
	}

	// Define the filter function to limit the search to executable files
	filterFunc := func(path string, info os.FileInfo) bool {
			return !info.IsDir() && filepath.Base(path) == executablePattern
	}

	// Define the predicate function to capture the path of the first matching file
	predicateFunc := func(path string, info os.FileInfo) {
			executablePath = path
			found = true
	}

	// Traverse the directory
	err := TraverseDirectory(
		TraverseDirectoryParams{
			RootDir:   searchDir,
			Predicate: predicateFunc,
			Filter:    filterFunc,
		},
	)

	if err != nil {
			return  fmt.Sprintf("error traversing directory: %v", err)
	}

	if !found {
			return  fmt.Sprintf("NOTFOUND")
	}

	return executablePath
}
