## [GitCloneSubdirsStruct]

### Usage
Defines the structure for the `gitCloneSubdirs` settings in the VS Code configuration.

### Reference
| Field       | Type   | Description                                       |
|-------------|--------|---------------------------------------------------|
| RepoURL     | string | URL of the Git repository.                        |
| StagingDir  | string | Directory where the repo is initially cloned.     |
| Subdirs     | string | Specific subdirectories to clone from the repo.   |
| DestDir     | string | Destination directory for the cloned subdirectories. |

---

## [GitPushingWorkToGitRemoteStruct]

### Usage
Defines the structure for pushing work to a remote Git repository.

### Reference
| Field          | Type     | Description                                          |
|----------------|----------|------------------------------------------------------|
| RelativePaths  | []string | Relative file paths to include in the Git push.      |
| AbsolutePaths  | []string | Absolute file paths to include in the Git push.      |

---

## [MiscOptimizeImagesStruct]

### Usage
Specifies settings for an image optimization task in the VS Code configuration.

### Reference
| Field            | Type   | Description                                      |
|------------------|--------|--------------------------------------------------|
| Location         | string | Directory containing images to optimize.         |
| BackupLocation   | string | Location for backups of the original images.     |
| OptimizePercent  | string | Percentage of optimization to apply.             |

---

## [AngularFrontendStruct]

### Usage
Holds configuration settings for an Angular frontend project in the VS Code environment.

### Reference
| Field            | Type     | Description                                 |
|------------------|----------|---------------------------------------------|
| Configurations   | []string | Angular build configurations.               |

---

## [ShopifyRunStruct]

### Usage
Specifies configuration for running Shopify-related tasks.

### Reference
| Field        | Type   | Description                  |
|--------------|--------|------------------------------|
| ProjectName  | string | Name of the Shopify project. |

---

## [FlutterMobileBuildStruct]

### Usage
Configures settings for Flutter mobile build processes in the VS Code environment.

### Reference
| Field                      | Type       | Description                                                  |
|----------------------------|------------|--------------------------------------------------------------|
| ToolArgs                   | []string   | Additional tool arguments used during the build process.     |
| Args                       | []string   | Arguments to pass through to the build process.              |
| VmAdditionalArgs           | []string   | Additional VM arguments for the Flutter engine.              |
| PlayStoreServiceAccountKey | string     | Service account key for Play Store uploads.                  |
| PackageName                | string     | Package name of the Flutter application.                     |
| PublishTarget              | string     | Target environment for publishing (e.g., Play Store).         |
| TrackName                  | string     | Release track for publishing (e.g., Alpha, Beta).            |

---

## [FirebaseCloudRunEmulatorsStruct]

### Usage
Holds configuration settings for Firebase Cloud Run emulators.

### Reference
| Field                        | Type       | Description                                    |
|------------------------------|------------|------------------------------------------------|
| GlobalDomain                 | string     | Global domain for the Firebase emulators.     |
| UIDomain0                    | string     | Domain for Firebase UI emulator.              |
| AuthDomain0                  | string     | Domain for Firebase Auth emulator.            |
| StorageDomain0               | string     | Domain for Firebase Storage emulator.         |
| DatabaseDomain0              | string     | Domain for Firebase Database emulator.        |
| HostingDomain0               | string     | Domain for Firebase Hosting emulator.         |
| FunctionsDomain0             | string     | Domain for Firebase Functions emulator.       |
| PubSubDomain0                | string     | Domain for Firebase Pub/Sub emulator.         |
| FirestoreDomain0             | string     | Domain for Firebase Firestore emulator.       |
| KillPortOutputFile           | string     | File to log killed ports.                     |
| KillPortOutputFileAcceptDefault | bool   | Whether to accept the default kill port file. |
| AdditonalPortsToKill         | []int      | Additional ports to terminate during cleanup. |

---

## [MiscReinitializeProjectStruct]

### Usage
Specifies settings for reinitializing a project within the VS Code configuration.

### Reference
| Field                                | Type       | Description                                               |
|-------------------------------------|------------|-----------------------------------------------------------|
| ProjectName                         | string     | Name of the project to reinitialize.                      |
| OrganizationName                    | string     | Name of the organization owning the project.              |
| AmountToAppendToPortNumberString    | int        | Increment value for port numbers.                         |
| VCSPrivateKey                       | string     | Private key for version control system.                   |
| WebSEODescription                   | string     | SEO description for the project.                          |
| WebSEOKeywords                      | string     | SEO keywords for the project.                             |
| ProxyURLs                           | []string   | List of proxy URLs for the project.                       |
| FlutterAndroidGoogleAdsID           | string     | Google Ads ID for Flutter Android apps.                   |
| FlutterAndroidGoogleApplicationID   | string     | Google Application ID for Flutter Android apps.           |
| FlutterIosGoogleAdsID               | string     | Google Ads ID for Flutter iOS apps.                       |
| FlutterIosGoogleApplicationID       | string     | Google Application ID for Flutter iOS apps.               |
| FlutterIosFacebookAppID             | string     | Facebook App ID for Flutter iOS apps.                     |
| FlutterIosFacebookClientToken       | string     | Facebook Client Token for Flutter iOS apps.               |
| FlutterIosFacebookCustomURLScheme   | string     | Facebook Custom URL Scheme for Flutter iOS apps.          |
| FlutterIosGoogleOAuthURLSchemes     | []string   | Google OAuth URL schemes for Flutter iOS apps.            |
| ChromeDriverPath                    | string     | Path to the Chrome WebDriver binary.                      |
| FirefoxDriverPath                   | string     | Path to the Firefox WebDriver binary.                     |
| OperaDriverPath                     | string     | Path to the Opera WebDriver binary.                       |
| EdgeDriverPath                      | string     | Path to the Edge WebDriver binary.                        |

---

## [WMLPorts]

### Usage
Defines port configurations for various development environments.

### Reference
| Field                      | Type   | Description                              |
|----------------------------|--------|------------------------------------------|
| AngularRun0                | int    | Port for running the Angular application.|
| AngularTest0               | int    | Port for testing the Angular application.|
| FirebaseEmulatorUI0        | int    | Port for Firebase Emulator UI.           |
| FirebaseEmulatorAuth0      | int    | Port for Firebase Auth emulator.         |
| FirebaseEmulatorStorage0   | int    | Port for Firebase Storage emulator.      |
| FirebaseEmulatorDatabase0  | int    | Port for Firebase Database emulator.     |
| FirebaseEmulatorHosting0   | int    | Port for Firebase Hosting emulator.      |
| FirebaseEmulatorFunctions0 | int    | Port for Firebase Functions emulator.    |
| FirebaseEmulatorFirestore0 | int    | Port for Firebase Firestore emulator.    |
| DiodeProxies0              | int    | Port for Diode Proxies.                  |

---

## [WindmillcodeExtensionPack]

### Usage
Encapsulates a set of configurations for the `windmillcode-extension-pack-0`.

### Reference
| Field                            | Type                        | Description                                                    |
|----------------------------------|-----------------------------|----------------------------------------------------------------|
| MiscReinitializeProject          | MiscReinitializeProjectStruct | Settings for project reinitialization.                        |
| TasksToRunOnFolderOpen           | []string                    | Tasks to execute when a folder is opened in VS Code.           |
| FlaskBackendDevHelperScript      | string                      | Script to assist in Flask backend development.                 |
| FlaskBackendTestHelperScript     | string                      | Script to assist in testing Flask backend.                     |
| ProxyURLs                        | string                      | Proxy URLs configuration.                                      |
| SQLDockerContainerName           | string                      | Docker container name for SQL.                                 |
| DatabaseName                     | string                      | Name of the database.                                          |
| DatabaseOptions                  | []string                    | Database options.                                              |
| OpenAIAPIKey0                    | string                      | OpenAI API key.                                                |
| OpenAIAPIBase0                   | string                      | OpenAI API base URL.                                           |
| LangCodes0                       | string                      | Language codes.                                                |
| Ports                            | WMLPorts                    | Port configurations for different environments.                |

---

## [VSCodeSettings]

### Usage
Represents the VS Code settings, specifically for the `windmillcode-extension-pack-0`.

### Reference
| Field          | Type                        | Description                                        |
|----------------|-----------------------------|----------------------------------------------------|
| ExtensionPack  | WindmillcodeExtensionPack   | Configuration for the `windmillcode-extension-pack-0`. |

---

## [GetSettingsJSON]

### Usage
Reads and returns the VS Code settings from the `.vscode/settings.json` file in the specified workspace folder, converting it to a `VSCodeSettings` struct.

### Reference
| Parameter        | Type   | Description                                      |
|------------------|--------|--------------------------------------------------|
| workSpaceFolder  | string | The workspace folder path where the settings file is located. |

| Returns          | Type             | Description                                     |
|------------------|------------------|------------------------------------------------|
| settings         | VSCodeSettings   | The VS Code settings parsed from the settings file. |
| error            | error            | An error, if there's an issue reading or parsing the file. |
