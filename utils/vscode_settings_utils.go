package utils

import (
	"encoding/json"
	"os"
)

type GitCloneSubdirsStruct struct {
	RepoURL    string `json:"repoURL,omitempty"`
	StagingDir string `json:"stagingDir,omitempty"`
	Subdirs    string `json:"subdirs,omitempty"`
	DestDir    string `json:"destDir,omitempty"`
}

type GitPushingWorkToGitRemoteStruct struct {
	RelativePaths []string `json:"relativePaths,omitempty"`
	AbsolutePaths []string `json:"absolutePaths,omitempty"`
}

type MiscOptimizeImagesStruct struct {
	Location        string `json:"location,omitempty"`
	BackupLocation  string `json:"backupLocation,omitempty"`
	OptimizePercent string `json:"optimizePercent,omitempty"`
}

type AngularFrontendStruct struct {
	Configurations []string `json:"configurations,omitempty"`
}

type ShopifyRunStruct struct {
	ProjectName string `json:"projectName,omitempty"`
}

type FlutterMobileBuildStruct struct {
	ToolArgs                   []string `json:"toolArgs,omitempty"`
	Args                       []string `json:"args,omitempty"`
	VmAdditionalArgs           []string `json:"vmAdditionalArgs,omitempty"`
	PlayStoreServiceAccountKey string   `json:"playStoreServiceAccountKey,omitempty"`
	PackageName                string   `json:"packageName,omitempty"`
	PublishTarget              string   `json:"publishTarget,omitempty"`
	TrackName                  string   `json:"trackName,omitempty"`
}

type FirebaseCloudRunEmulatorsStruct struct {
	GlobalDomain                    string `json:"globalDomain,omitempty"`
	UIDomain0                       string `json:"UIDomain0,omitempty"`
	AuthDomain0                     string `json:"AuthDomain0,omitempty"`
	StorageDomain0                  string `json:"StorageDomain0,omitempty"`
	DatabaseDomain0                 string `json:"DatabaseDomain0,omitempty"`
	HostingDomain0                  string `json:"HostingDomain0,omitempty"`
	FunctionsDomain0                string `json:"FunctionsDomain0,omitempty"`
	PubSubDomain0                   string `json:"PubSubDomain0,omitempty"`
	FirestoreDomain0                string `json:"FirestoreDomain0,omitempty"`
	KillPortOutputFile              string `json:"killPortOutputFile,omitempty"`
	KillPortOutputFileAcceptDefault bool   `json:"killPortOutputFileAcceptDefault,omitempty"`
	AdditonalPortsToKill            []int  `json:"additonalPortsToKill,omitempty"`
}

type ProcessIfDefaultIsPresentStruct struct {
	Global bool `json:"global,omitempty"`
}

type WMLPorts struct {
	AngularRun0                int `json:"Angular_Run_0,omitempty"`
	AngularTest0               int `json:"Angular_Test_0,omitempty"`
	AngularCoverageTest0       int `json:"Angular_Coverage_Test_0,omitempty"`
	AngularAnalyzer0           int `json:"Angular_Analyzer_0,omitempty"`
	AngularSSG0                int `json:"Angular_SSG_0,omitempty"`
	AngularSSG1                int `json:"Angular_SSG_1,omitempty"`
	AngularSSG2                int `json:"Angular_SSG_2,omitempty"`
	FlaskRun0                  int `json:"Flask_Run_0,omitempty"`
	FlaskTest0                 int `json:"Flask_Test_0,omitempty"`
	Postgres0                  int `json:"Postgres_0,omitempty"`
	MySQL0                     int `json:"My_SQL_0,omitempty"`
  ReactNativeExpoRun0        int `json:"React_Native_Expo_Run_0,omitempty"`
  LaravelRun0                int `json:"Laravel_Run_0,omitempty"`
	FirebaseEmulatorUI0        int `json:"Firebase_Emulator_UI_0,omitempty"`
	FirebaseEmulatorAuth0      int `json:"Firebase_Emulator_Auth_0,omitempty"`
	FirebaseEmulatorStorage0   int `json:"Firebase_Emulator_Storage_0,omitempty"`
	FirebaseEmulatorDatabase0  int `json:"Firebase_Emulator_Database_0,omitempty"`
	FirebaseEmulatorHosting0   int `json:"Firebase_Emulator_Hosting_0,omitempty"`
	FirebaseEmulatorFunctions0 int `json:"Firebase_Emulator_Functions_0,omitempty"`
	FirebaseEmulatorPubSub0    int `json:"Firebase_Emulator_PubSub_0,omitempty"`
	FirebaseEmulatorFirestore0 int `json:"Firebase_Emulator_Firestore_0,omitempty"`
	DiodeProxies0              int `json:"Diode_Proxies_0,omitempty"`
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
		w.FirebaseEmulatorUI0,
	}
}

type MiscReinitializeProjectStruct struct {
	AmountToAppendToPortNumberString  int      `json:"amountToAppendToPortNumberString,omitempty"`
	ProjectName                       string   `json:"projectName,omitempty"`
	OrganizationName                  string   `json:"organizationName,omitempty"`
	VCSPrivateKey                     string   `json:"vcsPrivateKey,omitempty"`
	WebSEODescription                 string   `json:"webSEODescription,omitempty"`
	WebSEOKeywords                    string   `json:"webSEOKeywords,omitempty"`
	ProxyURLs                         []string `json:"proxyURLs"`
	FlutterAndroidGoogleAdsID        string `json:"flutterAndroidGoogleAdsID,omitempty"`
	FlutterAndroidGoogleApplicationID string `json:"flutterAndroidGoogleApplicationID,omitempty"`
	FlutterIosGoogleAdsID            string `json:"flutterIosGoogleAdsID,omitempty"`
	FlutterIosGoogleApplicationID    string `json:"flutterIosGoogleApplicationID,omitempty"`
	FlutterIosFacebookAppID    string `json:"flutterIosFacebookAppID,omitempty"`
	FlutterIosFacebookClientToken    string `json:"flutterIosFacebookClientToken,omitempty"`
	FlutterIosFacebookCustomURLScheme    string `json:"flutterIosFacebookCustomURLScheme,omitempty"`
	FlutterIosGoogleOAuthURLSchemes    []string `json:"flutterIosGoogleOAuthURLSchemes,omitempty"`
	ChromeDriverPath                 string `json:"chromeDriverPath,omitempty"`
	FirefoxDriverPath                string `json:"firefoxDriverPath,omitempty"`
	OperaDriverPath                  string `json:"operaDriverPath,omitempty"`
	EdgeDriverPath 	               string `json:"edgeDriverPath,omitempty"`
}
type WindmillcodeExtensionPack struct {
	MiscReinitializeProject      MiscReinitializeProjectStruct   `json:"miscReinitializeProject,omitempty"`
	TasksToRunOnFolderOpen       []string                        `json:"tasksToRunOnFolderOpen,omitempty"`
	FlaskBackendDevHelperScript  string                          `json:"flaskBackendDevHelperScript,omitempty"`
	FlaskBackendTestHelperScript string                          `json:"flaskBackendTestHelperScript,omitempty"`
	ProxyURLs                    string                          `json:"proxyURLs,omitempty"`
	SQLDockerContainerName       string                          `json:"sqlDockerContainerName,omitempty"`
	DatabaseName                 string                          `json:"databaseName,omitempty"`
	DatabaseOptions              []string                        `json:"databaseOptions,omitempty"`
	Environments                 []string                        `json:"environments,omitempty"`
	SentryDSN                    string                          `json:"sentryDSN,omitempty"`
	OpenAIAPIKey0                string                          `json:"openAIAPIKey0,omitempty"`
	OpenAIAPIBase0               string                          `json:"openAIAPIBase0,omitempty"`
	LangCodes0                   string                          `json:"langCodes0,omitempty"`
	PythonVersion0               string                          `json:"pythonVersion0,omitempty"`
	NodeJSVersion0               string                          `json:"nodeJSVersion0,omitempty"`
	JavaVersion0                 string                          `json:"javaVersion0,omitempty"`
	GoVersion0                   string                          `json:"goVersion0,omitempty"`
	RubyVersion0                 string                          `json:"rubyVersion0,omitempty"`
	DartVersion0                 string                          `json:"dartVersion0,omitempty"`
	CSharpVersion0               string                          `json:"cSharpVersion0,omitempty"`
	SwiftVersion0                string                          `json:"swiftVersion0,omitempty"`
	PHPVersion0                  string                          `json:"phpVersion0,omitempty"`
	RustVersion0                 string                          `json:"rustVersion0,omitempty"`
	KotlinVersion0               string                          `json:"kotlinVersion0,omitempty"`
	ScalaVersion0                string                          `json:"scalaVersion0,omitempty"`
	PerlVersion0                 string                          `json:"perlVersion0,omitempty"`
	LuaVersion0                  string                          `json:"luaVersion0,omitempty"`
	HaskellVersion0              string                          `json:"haskellVersion0,omitempty"`
	ClojureVersion0              string                          `json:"clojureVersion0,omitempty"`
	ErlangVersion0               string                          `json:"erlangVersion0,omitempty"`
	JuliaVersion0                string                          `json:"juliaVersion0,omitempty"`
	ObjectiveCVersion0           string                          `json:"objectiveCVersion0,omitempty"`
	FSharpVersion0               string                          `json:"fSharpVersion0,omitempty"`
	VisualBasicVersion0          string                          `json:"visualBasicVersion0,omitempty"`
	Ports                        WMLPorts                        `json:"ports,omitempty"`
	ProcessIfDefaultIsPresent    ProcessIfDefaultIsPresentStruct `json:"processIfDefaultIsPresent,omitempty"`
	FirebaseCloudRunEmulators    FirebaseCloudRunEmulatorsStruct `json:"firebaseCloudRunEmulators,omitempty"`
	FlutterMobileBuild           FlutterMobileBuildStruct        `json:"flutterMobileBuild,omitempty"`
	GitCloneSubdirs              GitCloneSubdirsStruct           `json:"gitCloneSubdirs,omitempty"`
	GitPushingWorkToGitRemote    GitPushingWorkToGitRemoteStruct `json:"gitPushingWorkingToGitRemote,omitempty"`
	MiscOptimizeImages           MiscOptimizeImagesStruct        `json:"miscOptimizeImages,omitempty"`
	AngularFrontend              AngularFrontendStruct           `json:"angularFrontend,omitempty"`
	ShopifyRun                   ShopifyRunStruct                `json:"shopifyRun,omitempty"`
	WxtBuildSafari               WxtBuildSafariStruct            `json:"wxtBuildSafari,omitempty"`
	AngularDeployToFirebase      AngularDeployToFirebaseStruct   `json:"angularDeployToFirebase,omitempty"`
}

type AngularDeployToFirebaseStruct struct {
	Environments           []string `json:"environments,omitempty"`
	RunLint                bool     `json:"runLint,omitempty"`
	RunSSGScript           bool     `json:"runSSGScript,omitempty"`
	RemoveBuildDirectories bool     `json:"removeBuildDirectories,omitempty"`
	DeployToSentry         bool     `json:"deployToSentry,omitempty"`
	SentryOrg              string   `json:"sentryOrg,omitempty"`
	SentryProject          string   `json:"sentryProject,omitempty"`
	SentryAuthToken        string   `json:"sentryAuthToken,omitempty"`
	DeployToFirebase       bool     `json:"deployToFirebase,omitempty"`
	FirebaseProjectId      string   `json:"firebaseProjectId,omitempty"`
}

type WxtBuildSafariStruct struct {
	BundleIdentifier string `json:"bundleIdentifier,omitempty"`
}

type VSCodeSettings struct {
	ExtensionPack WindmillcodeExtensionPack `json:"windmillcode-extension-pack-0,omitempty"`
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
	standardJSON, err := CleanJSON(content)
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
