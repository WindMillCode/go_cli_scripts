package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/windmillcode/go_scripts/v2/utils"
)

type Task struct {
	Label   string `json:"label"`
	Type    string `json:"type"`
	Windows struct {
		Command string `json:"command"`
	} `json:"windows"`
	Linux struct {
		Command string `json:"command"`
	} `json:"linux"`
	Osx struct {
		Command string `json:"command"`
	} `json:"osx"`
	RunOptions struct {
		RunOn         string `json:"runOn"`
		InstanceLimit int    `json:"instanceLimit"`
	} `json:"runOptions"`
}

type Input struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Default     string `json:"default"`
	Type        string `json:"type"`
}
type TasksJSON struct {
	Version string  `json:"version"`
	Tasks   []Task  `json:"tasks"`
	Inputs  []Input `json:"inputs"`
}

func main() {
	workSpaceFolder := os.Args[1]
	extensionFolder := os.Args[2]
	tasksJsonRelativeFilePath := os.Args[3]
	goExecutable := os.Args[4]
	cliInfo := utils.ShowMenuModel{
		Prompt:  "This will delete your vscode/tasks.json in your workspace folder. If you don't have a .vscode/tasks.json or you have not used this command before, it is safe to choose TRUE. Otherwise, consult with your manager before continuing",
		Choices: []string{"TRUE", "FALSE"},
	}
	proceed := utils.ShowMenu(cliInfo, nil)

	tasksJsonFilePath := filepath.Join(extensionFolder, tasksJsonRelativeFilePath)

	content, err, shouldReturn := createTasksJson(tasksJsonFilePath,false)
	if shouldReturn {
		return
	}
	// fileContent := string(content)
	// fmt.Println("File Contents:")
	// fmt.Println(fileContent)
	var tasksJSON TasksJSON
	err = json.Unmarshal(content, &tasksJSON)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	goScriptsSourceDirPath := filepath.Join(extensionFolder, "task_files/go_scripts")
	goScriptsDestDirPath := filepath.Join(workSpaceFolder, "ignore/Windmillcode/go_scripts")

	if proceed == "TRUE" {

		for index, task := range tasksJSON.Tasks {

			pattern0 := ":"
			regex0 := regexp.MustCompile(pattern0)
			programLocation0 := regex0.Split(task.Label, -1)
			pattern1 := " "
			regex1 := regexp.MustCompile(pattern1)
			programLocation1 := regex1.Split(strings.Join(programLocation0, ""), -1)
			programLocation2 := strings.Join(programLocation1, "_")
			programLocation3 := "ignore//${input:current_user_0}//go_scripts//" + programLocation2
			linuxTaskExecutable := ".//main"
			linuxCommand0 := "cd " + programLocation3 + " ; " + linuxTaskExecutable
			windowsCommand0 := "cd " + strings.Replace(programLocation3, "//", "\\", -1) + " ; " + strings.Replace(linuxTaskExecutable, "//", "\\", -1)

			tasksJSON.Tasks[index].Windows.Command = windowsCommand0
			tasksJSON.Tasks[index].Osx.Command = linuxCommand0
			tasksJSON.Tasks[index].Linux.Command = linuxCommand0

		}

		tasksJSONData, err := json.MarshalIndent(tasksJSON, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		workspaceTasksJSONFilePath := filepath.Join(workSpaceFolder, "/.vscode/tasks.json")
		workspaceTasksJSONFile, err := os.OpenFile(workspaceTasksJSONFilePath, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer workspaceTasksJSONFile.Close()
		_, err = workspaceTasksJSONFile.Write(tasksJSONData)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		utils.CopyDir(goScriptsSourceDirPath, goScriptsDestDirPath)
	}

	var rebuild string
	if proceed == "TRUE" {
		rebuild = "TRUE"
	} else {
		cliInfo = utils.ShowMenuModel{
			Prompt:  "Do you want to rebuild the go programs into binary exectuables ",
			Choices: []string{"TRUE", "FALSE"},
		}
		rebuild = utils.ShowMenu(cliInfo, nil)
	}

	if rebuild == "TRUE" {
		var wg sync.WaitGroup
		fmt.Print(len(tasksJSON.Tasks))
		for _, task := range tasksJSON.Tasks {
			wg.Add(1)

			pattern0 := ":"
			regex0 := regexp.MustCompile(pattern0)
			programLocation0 := regex0.Split(task.Label, -1)
			pattern1 := " "
			regex1 := regexp.MustCompile(pattern1)
			programLocation1 := regex1.Split(strings.Join(programLocation0, ""), -1)
			programLocation2 := strings.Join(programLocation1, "_")
			absProgramLocation := filepath.Join(goScriptsDestDirPath, programLocation2)
			go func() {
				defer wg.Done()
				buildGoCLIProgram(absProgramLocation, goExecutable)
			}()
		}
		wg.Wait()
	}

}

func createTasksJson(tasksJsonFilePath string, triedCreateOnError bool) ([]byte, error, bool) {
	content, err := os.ReadFile(tasksJsonFilePath)
	if err != nil {
		if triedCreateOnError {
			return nil, err, true
		}

		// If the file doesn't exist, create it.
		_, createErr := os.Create(tasksJsonFilePath)
		if createErr != nil {
			return nil, createErr, true
		}

		// Recursively attempt to read the file after creating it.
		return createTasksJson(tasksJsonFilePath, true)
	}

	return content, nil, false
}


func buildGoCLIProgram(programLocation string, goExecutable string) {

	fmt.Printf("%s \n", programLocation)
	utils.RunCommandInSpecificDirectory(goExecutable, []string{"build", "main.go"}, programLocation)
	fmt.Printf("Finished building %s \n", programLocation)

}
