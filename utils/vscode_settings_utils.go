package utils

import (
	"encoding/json"
	"os"
)



type GitCloneSubdirsStruct struct {
	RepoURL    string `json:"repoURL"`
	StagingDir string `json:"stagingDir"`
	Subdirs    string `json:"subdirs"`
	DestDir    string `json:"destDir"`
}

type GitPushingWorkToGitRemoteStruct struct {
	RelativePaths []string `json:"relativePaths"`
	AbsolutePaths []string `json:"absolutePaths"`
}

type MiscOptimizeImagesStruct struct {
	Location        string `json:"location"`
	BackupLocation  string `json:"backupLocation"`
	OptimizePercent string `json:"optimizePercent"`
}

type AngularFrontendStruct struct{
	Configurations   []string `json:"configurations"`
}

type FlutterMobileBuild struct {
	ToolArgs                     []string `json:"toolArgs"`
	Args                         []string `json:"args"`
	VmAdditionalArgs             []string `json:"vmAdditionalArgs"`
	PlayStoreServiceAccountKey   string   `json:"playStoreServiceAccountKey"`
	PackageName                  string   `json:"packageName"`
	PublishTarget                string   `json:"publishTarget"`
	TrackName                    string   `json:"trackName"`
}


type WindmillcodeExtensionPack struct {
	TasksToRunOnFolderOpen       []string `json:"tasksToRunOnFolderOpen"`
	FlaskBackendDevHelperScript  string   `json:"flaskBackendDevHelperScript"`
	FlaskBackendTestHelperScript string   `json:"flaskBackendTestHelperScript"`
	ProxyURLs                    string   `json:"proxyURLs"`
	SQLDockerContainerName       string   `json:"sqlDockerContainerName"`
	DatabaseName                 string   `json:"databaseName"`
	DatabaseOptions              []string `json:"databaseOptions"`
	Environments                 []string `json:"environments"`
	SentryDSN                    string   `json:"sentryDSN"`
	OpenAIAPIKey0                string   `json:"openAIAPIKey0"`
	OpenAIAPIBase0               string   `json:"openAIAPIBase0"`
	LangCodes0                   string   `json:"langCodes0"`
	PythonVersion0               string   `json:"pythonVersion0"`
	NodeJSVersion0               string   `json:"nodeJSVersion0"`
	JavaVersion0                 string   `json:"javaVersion0"`
	GoVersion0                   string   `json:"goVersion0"`
	RubyVersion0                 string   `json:"rubyVersion0"`
	DartVersion0                 string   `json:"dartVersion0"`
	CSharpVersion0               string   `json:"cSharpVersion0"`
	SwiftVersion0                string   `json:"swiftVersion0"`
	PHPVersion0                  string   `json:"phpVersion0"`
	RustVersion0                 string   `json:"rustVersion0"`
	KotlinVersion0               string   `json:"kotlinVersion0"`
	ScalaVersion0                string   `json:"scalaVersion0"`
	PerlVersion0                 string   `json:"perlVersion0"`
	LuaVersion0                  string   `json:"luaVersion0"`
	HaskellVersion0              string   `json:"haskellVersion0"`
	ClojureVersion0              string   `json:"clojureVersion0"`
	ErlangVersion0               string   `json:"erlangVersion0"`
	JuliaVersion0                string   `json:"juliaVersion0"`
	ObjectiveCVersion0           string   `json:"objectiveCVersion0"`
	FSharpVersion0               string   `json:"fSharpVersion0"`
	VisualBasicVersion0          string   `json:"visualBasicVersion0"`
	FlutterMobileBuild           FlutterMobileBuild `json:"flutterMobileBuild"`
	GitCloneSubdirs              GitCloneSubdirsStruct    `json:"gitCloneSubdirs"`
	GitPushingWorkToGitRemote    GitPushingWorkToGitRemoteStruct `json:"gitPushingWorkingToGitRemote"`
	MiscOptimizeImages           MiscOptimizeImagesStruct `json:"miscOptimizeImages"`
	AngularFrontend              AngularFrontendStruct    `json:"angularFrontend"`
}


type VSCodeSettings struct {
	ExtensionPack WindmillcodeExtensionPack `json:"windmillcode-extension-pack-0"`
}

func GetSettingsJSON(workSpaceFolder string) (VSCodeSettings, error) {
	settingsJSONFilePath := JoinAndConvertPathToOSFormat(workSpaceFolder, "/.vscode/settings.json")
	var settings VSCodeSettings
	content, err := os.ReadFile(settingsJSONFilePath)
	if err != nil {
		LogErrorWithTraceBack("Error reading file:", err)
		return settings, err
	}
	standardJSON, err := RemoveComments(content)
	err = json.Unmarshal([]byte(standardJSON), &settings)
	if err != nil {
		LogErrorWithTraceBack("Error unmarshalling JSON:", err)
		return settings, err
	}
	return settings, nil
}

