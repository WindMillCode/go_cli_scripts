package utils

import (
	"encoding/json"
	"os"
)
type VSCodeTasksShellOptions struct {
	Executable string   `json:"executable"`
	Args       []string `json:"args"`
}

type VSCodeTasksCommandOptions struct {
	Shell VSCodeTasksShellOptions `json:"shell"`
}

type VSCodeTasksMetadata struct {
	Name string `json:"name"`
}

type VSCodeTasksRunOptions struct {
	RunOn         string `json:"runOn,omitempty"`
	InstanceLimit int    `json:"instanceLimit"`
}

type VSCodeTasksTask struct {
	Label   string `json:"label"`
	Type    string `json:"type,omitempty"`
	Windows struct {
		Command string `json:"command"`
	} `json:"windows,omitempty"`
	Linux struct {
		Command string         `json:"command"`
		Options VSCodeTasksCommandOptions `json:"options"`
	} `json:"linux,omitempty"`
	Osx struct {
		Command string        `json:"command"`
		Args    []string      `json:"args"`
	} `json:"osx,omitempty"`
	RunOptions   VSCodeTasksRunOptions `json:"runOptions,omitempty"`
	Presentation struct {
		Panel string `json:"panel,omitempty"`
	} `json:"presentation,omitempty"`
	Metadata VSCodeTasksMetadata `json:"metadata,omitempty"`
}

type VSCodeTasksInput struct {
	ID          string   `json:"id,omitempty"`
	Description string   `json:"description,omitempty"`
	Default     string   `json:"default,omitempty"`
	Type        string   `json:"type,omitempty"`
	Metadata    VSCodeTasksMetadata `json:"metadata,omitempty"`
}

type VSCodeTasksTasksJSON struct {
	Version string  `json:"version,omitempty"`
	Tasks   []VSCodeTasksTask  `json:"tasks,omitempty"`
	Inputs  []VSCodeTasksInput `json:"inputs,omitempty"`
}

type VSCodeTasksDynamicTasksJSON struct {
	Version string            `json:"version,omitempty"`
	Tasks   []json.RawMessage `json:"tasks,omitempty"`
	Inputs  []json.RawMessage `json:"inputs,omitempty"`
}

func CreateTasksJson(tasksJsonFilePath string, triedCreateOnError bool) ([]byte, error, bool) {

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
		return CreateTasksJson(tasksJsonFilePath, true)
	}

	return content, nil, false
}
