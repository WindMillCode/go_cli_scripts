package utils

import (
	"fmt"
	"os"
	"github.com/ghodss/yaml"
)



type GitCloneSubdirsStruct struct {
	RepoURL    string `json:"repoURL"`
	StagingDir string `json:"stagingDir"`
	Subdirs    string `json:"subdirs"`
	DestDir    string `json:"destDir"`
}

type MiscOptimizeImagesStruct struct {
	Location        string `json:"location"`
	BackupLocation  string `json:"backupLocation"`
	OptimizePercent string `json:"optimizePercent"`
}

type AngularFrontendStruct struct{
	Configurations   []string `json:"configurations"`
}

type WindmillcodeExtensionPack struct {
	TasksToRunOnFolderOpen            []string                 `json:"tasksToRunOnFolderOpen"`
	FlaskBackendDevHelperScript  			string                   `json:"flaskBackendDevHelperScript"`
	FlaskBackendTestHelperScript 			string                   `json:"flaskBackendTestHelperScript"`
	ProxyURLs                    			string                   `json:"proxyURLs"`
	SQLDockerContainerName       			string                   `json:"sqlDockerContainerName"`
	DatabaseName                 			string                   `json:"databaseName"`
	DatabaseOptions              			[]string                 `json:"databaseOptions"`
	OpenAIAPIKey0                			string                   `json:"openAIAPIKey0"`
	OpenAIAPIBase0               			string                   `json:"openAIAPIBase0"`
	LangCodes0                   			string                   `json:"langCodes0"`
	PythonVersion0               			string                   `json:"pythonVersion0"`
	GitCloneSubdirs              			GitCloneSubdirsStruct    `json:"gitCloneSubdirs"`
	MiscOptimizeImages           			MiscOptimizeImagesStruct `json:"miscOptimizeImages"`
	AngularFrontend              			AngularFrontendStruct    `json:"angularFrontend"`
}

type VSCodeSettings struct {
	ExtensionPack WindmillcodeExtensionPack `json:"windmillcode-extension-pack-0"`
}

func GetSettingsJSON(workSpaceFolder string) (VSCodeSettings, error) {
	settingsJSONFilePath := JoinAndConvertPathToOSFormat(workSpaceFolder, "/.vscode/settings.json")
	var settings VSCodeSettings
	content, err := os.ReadFile(settingsJSONFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err.Error())
		return settings, err
	}
	err = yaml.Unmarshal(content, &settings)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err.Error())
		return settings, err
	}
	return settings, nil
}

