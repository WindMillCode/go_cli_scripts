package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/windmillcode/go_scripts/utils"
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
	if proceed == "FALSE" {
		return
	}
	tasksJsonFilePath := filepath.Join(extensionFolder, tasksJsonRelativeFilePath)

	content, err := ioutil.ReadFile(tasksJsonFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
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
	for index, task := range tasksJSON.Tasks {
		pattern0 := ":"
		regex0 := regexp.MustCompile(pattern0)
		programLocation0 := regex0.Split(task.Label, -1)
		pattern1 := " "
		regex1 := regexp.MustCompile(pattern1)
		programLocation1 := regex1.Split(strings.Join(programLocation0, ""), -1)
		programLocation2 := strings.Join(programLocation1, "_")
		programLocation3 := "ignore//${input:current_user_0}//go_scripts//" + programLocation2
		linuxCommand0 := "cd " + programLocation3 + " ; " + goExecutable + " run . "
		windowsCommand0 := "cd " + strings.Replace(programLocation3, "//", "\\", -1) + " ; " + goExecutable + " run . "

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

	goScriptsSourceDirPath := filepath.Join(extensionFolder, "task_files/go_scripts")
	goScriptsDestDirPath := filepath.Join(workSpaceFolder, "ignore/Windmillcode/go_scripts")

	if err := os.RemoveAll(goScriptsDestDirPath); err != nil {
		fmt.Println("Error:", err)
		return
	}
	utils.CopyDir(goScriptsSourceDirPath, goScriptsDestDirPath)
}
