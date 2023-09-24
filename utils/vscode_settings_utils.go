package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type MiscOptimizeImagesStruct struct {
	Location       string `json:"location"`
	BackupLocation string `json:"backupLocation"`
	OptimizePercent string `json:"optimizePercent"`
}

type GitCloneSubdirsStruct  struct {
	RepoURL    string `json:"repoURL"`
	StagingDir string `json:"stagingDir"`
	Subdirs    string `json:"subdirs"`
	DestDir    string `json:"destDir"`
}

type WindmillcodeExtensionPack struct {
	SQLDockerContainerName string                   `json:"sqlDockerContainerName"`
	DatabaseName           string                   `json:"databaseName"`
	DatabaseOptions        []string                 `json:"databaseOptions"`
	OpenAIAPIKey0          string                   `json:"openAIAPIKey0"`
	LangCodes0	           string                   `json:"langCodes0"`
	PythonVersion0         string                   `json:"pythonVersion0"`
	GitCloneSubdirs        GitCloneSubdirsStruct    `json:"gitCloneSubdirs"`
	MiscOptimizeImages     MiscOptimizeImagesStruct `json:"miscOptimizeImages"`
}

type VSCodeSettings struct {
	ExtensionPack WindmillcodeExtensionPack `json:"windmillcode-extension-pack-0"`
}


func GetSettingsJSON (workSpaceFolder string) (VSCodeSettings,error){
	settingsJSONFilePath := filepath.Join(workSpaceFolder,"/.vscode/settings.json")
	var settings VSCodeSettings
	content, err := ioutil.ReadFile(settingsJSONFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err.Error())
		return settings,err
	}
	err = json.Unmarshal(content, &settings)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err.Error())
		return settings,err
	}
	return settings,nil
}
