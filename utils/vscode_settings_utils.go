package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)



type WindmillcodeExtensionPack struct {
	SQLDockerContainerName string   `json:"sqlDockerContainerName"`
	DatabaseName           string   `json:"databaseName"`
	DatabaseOptions        []string `json:"databaseOptions"`
	OpenAIAPIKey0          string   `json:"openAIAPIKey0"`
	LangCodes0	           string   `json:"langCodes0"`
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
