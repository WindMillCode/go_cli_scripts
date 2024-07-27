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

type FlutterMobileBuildStruct struct {
	ToolArgs                     []string `json:"toolArgs"`
	Args                         []string `json:"args"`
	VmAdditionalArgs             []string `json:"vmAdditionalArgs"`
	PlayStoreServiceAccountKey   string   `json:"playStoreServiceAccountKey"`
	PackageName                  string   `json:"packageName"`
	PublishTarget                string   `json:"publishTarget"`
	TrackName                    string   `json:"trackName"`
}

type FirebaseCloudRunEmulatorsStruct struct {
	GlobalDomain    string `json:"globalDomain"`
  AuthDomain0      string `json:"AuthDomain0"`
  StorageDomain0   string `json:"StorageDomain0"`
  DatabaseDomain0  string `json:"DatabaseDomain0"`
  HostingDomain0   string `json:"HostingDomain0"`
  FunctionsDomain0 string `json:"FunctionsDomain0"`
  PubSubDomain0    string `json:"PubSubDomain0"`
  FirestoreDomain0 string `json:"FirestoreDomain0"`
}

type ProcessIfDefaultIsPresentStruct struct{
	Global  bool `json:"global"`
}

type WMLPorts struct {
  AngularRun0                int `json:"Angular_Run_0"`
  AngularTest0               int `json:"Angular_Test_0"`
  AngularCoverageTest0       int `json:"Angular_Coverage_Test_0"`
  AngularAnalyzer0           int `json:"Angular_Analyzer_0"`
  AngularSSG0                int `json:"Angular_SSG_0"`
  AngularSSG1                int `json:"Angular_SSG_1"`
  AngularSSG2                int `json:"Angular_SSG_2"`
  FlaskRun0                  int `json:"Flask_Run_0"`
  FlaskTest0                 int `json:"Flask_Test_0"`
  Postgres0                  int `json:"Postgres_0"`
  FirebaseEmulatorAuth0      int `json:"Firebase_Emulator_Auth_0"`
  FirebaseEmulatorStorage0   int `json:"Firebase_Emulator_Storage_0"`
  FirebaseEmulatorDatabase0  int `json:"Firebase_Emulator_Database_0"`
  FirebaseEmulatorHosting0   int `json:"Firebase_Emulator_Hosting_0"`
  FirebaseEmulatorFunctions0 int `json:"Firebase_Emulator_Functions_0"`
  FirebaseEmulatorPubSub0    int `json:"Firebase_Emulator_PubSub_0"`
  FirebaseEmulatorFirestore0 int `json:"Firebase_Emulator_Firestore_0"`
  DiodeProxies0              int `json:"Diode_Proxies_0"`
}

func (w *WMLPorts) GetFirebasePorts() []int {
  return []int{
    w.FirebaseEmulatorAuth0,
    w.FirebaseEmulatorStorage0,
    w.FirebaseEmulatorDatabase0,
    w.FirebaseEmulatorHosting0,
    w.FirebaseEmulatorFunctions0,
    w.FirebaseEmulatorPubSub0,
    w.FirebaseEmulatorFirestore0,
  }
}


type WindmillcodeExtensionPack struct {
	TasksToRunOnFolderOpen       []string                         `json:"tasksToRunOnFolderOpen"`
	FlaskBackendDevHelperScript  string                           `json:"flaskBackendDevHelperScript"`
	FlaskBackendTestHelperScript string                           `json:"flaskBackendTestHelperScript"`
	ProxyURLs                    string                           `json:"proxyURLs"`
	SQLDockerContainerName       string                           `json:"sqlDockerContainerName"`
	DatabaseName                 string                           `json:"databaseName"`
	DatabaseOptions              []string                         `json:"databaseOptions"`
	Environments                 []string                         `json:"environments"`
	SentryDSN                    string                           `json:"sentryDSN"`
	OpenAIAPIKey0                string                           `json:"openAIAPIKey0"`
	OpenAIAPIBase0               string                           `json:"openAIAPIBase0"`
	LangCodes0                   string                           `json:"langCodes0"`
	PythonVersion0               string                           `json:"pythonVersion0"`
	NodeJSVersion0               string                           `json:"nodeJSVersion0"`
	JavaVersion0                 string                           `json:"javaVersion0"`
	GoVersion0                   string                           `json:"goVersion0"`
	RubyVersion0                 string                           `json:"rubyVersion0"`
	DartVersion0                 string                           `json:"dartVersion0"`
	CSharpVersion0               string                           `json:"cSharpVersion0"`
	SwiftVersion0                string                           `json:"swiftVersion0"`
	PHPVersion0                  string                           `json:"phpVersion0"`
	RustVersion0                 string                           `json:"rustVersion0"`
	KotlinVersion0               string                           `json:"kotlinVersion0"`
	ScalaVersion0                string                           `json:"scalaVersion0"`
	PerlVersion0                 string                           `json:"perlVersion0"`
	LuaVersion0                  string                           `json:"luaVersion0"`
	HaskellVersion0              string                           `json:"haskellVersion0"`
	ClojureVersion0              string                           `json:"clojureVersion0"`
	ErlangVersion0               string                           `json:"erlangVersion0"`
	JuliaVersion0                string                           `json:"juliaVersion0"`
	ObjectiveCVersion0           string                           `json:"objectiveCVersion0"`
	FSharpVersion0               string                           `json:"fSharpVersion0"`
	VisualBasicVersion0          string                           `json:"visualBasicVersion0"`
	Ports                        WMLPorts                         `json:"ports"`
	ProcessIfDefaultIsPresent    ProcessIfDefaultIsPresentStruct  `json:"processIfDefaultIsPresent"`
	FirebaseCloudRunEmulators    FirebaseCloudRunEmulatorsStruct  `json:"firebaseCloudRunEmulators"`
	FlutterMobileBuild           FlutterMobileBuildStruct         `json:"flutterMobileBuild"`
	GitCloneSubdirs              GitCloneSubdirsStruct            `json:"gitCloneSubdirs"`
	GitPushingWorkToGitRemote    GitPushingWorkToGitRemoteStruct  `json:"gitPushingWorkingToGitRemote"`
	MiscOptimizeImages           MiscOptimizeImagesStruct         `json:"miscOptimizeImages"`
	AngularFrontend              AngularFrontendStruct            `json:"angularFrontend"`
}


type VSCodeSettings struct {
	ExtensionPack WindmillcodeExtensionPack `json:"windmillcode-extension-pack-0"`
}

func GetSettingsJSON(workSpaceFolder string) (VSCodeSettings, error) {
	settingsJSONFilePath := JoinAndConvertPathToOSFormat(workSpaceFolder, "/.vscode/settings.json")
	var settings VSCodeSettings
	if !settings.ExtensionPack.ProcessIfDefaultIsPresent.Global {
		settings.ExtensionPack.ProcessIfDefaultIsPresent.Global = true
	}
	content, err := os.ReadFile(settingsJSONFilePath)
	if err != nil {
		LogErrorWithTraceBack("Error reading file:", err)
		return settings, err
	}
	standardJSON, err := RemoveComments(content)
	if err != nil {
		return VSCodeSettings{}, err
	}
	err = json.Unmarshal([]byte(standardJSON), &settings)
	if err != nil {
		LogErrorWithTraceBack("Error unmarshalling JSON:", err)
		return settings, err
	}
	return settings, nil
}

