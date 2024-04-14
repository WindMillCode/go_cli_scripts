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

## [MiscOptimizeImagesStruct]
### Usage
Specifies settings for an image optimization task in the VS Code configuration.

### Reference
| Field            | Type   | Description                                      |
|------------------|--------|--------------------------------------------------|
| Location         | string | Directory containing images to optimize.         |
| BackupLocation   | string | Location for backups of the original images.     |
| OptimizePercent  | string | Percentage of optimization to apply.             |

## [AngularFrontendStruct]
### Usage
Holds configuration settings for an Angular frontend project in the VS Code environment.

### Reference
| Field            | Type     | Description                                 |
|------------------|----------|---------------------------------------------|
| Configurations   | []string | Angular build configurations.               |

## [FlutterMobileBuild]
### Usage
Configures settings for Flutter mobile build processes in the VS Code environment.

### Reference
| Field              | Type       | Description                                                  |
|--------------------|------------|--------------------------------------------------------------|
| ToolArgs           | []string   | Additional tool arguments used during the build process.     |
| Args               | []string   | Arguments to pass through to the build process.              |
| VmAdditionalArgs   | []string   | Additional VM arguments for the Flutter engine.              |


## [WindmillcodeExtensionPack]
### Usage
Encapsulates a set of configurations for a specific VS Code extension pack.

### Reference
| Field                            | Type                        | Description                                                    |
|----------------------------------|-----------------------------|----------------------------------------------------------------|
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
| PythonVersion0                   | string                      | Python version.                                                |
| GitCloneSubdirs                  | GitCloneSubdirsStruct       | Settings for cloning Git subdirectories.                       |
| MiscOptimizeImages               | MiscOptimizeImagesStruct    | Settings for optimizing images.                                |
| AngularFrontend                  | AngularFrontendStruct       | Angular frontend configurations.                               |
| FlutterMobileBuild               | FlutterMobileBuild          | Configurations specific to Flutter mobile builds.              |

## [VSCodeSettings]
### Usage
Represents the VS Code settings, specifically for the `windmillcode-extension-pack-0`.

### Reference
| Field          | Type                        | Description                                        |
|----------------|-----------------------------|----------------------------------------------------|
| ExtensionPack  | WindmillcodeExtensionPack   | Configuration for the `windmillcode-extension-pack-0`. |

## [GetSettingsJSON]
### Usage
Reads and returns the VS Code settings from the `.vscode/settings.json` file in the specified workspace folder, converting it to a `VSCodeSettings` struct.

### Reference
| Parameter        | Type   | Description                                      |
|------------------|--------|--------------------------------------------------|
| workSpaceFolder  | string | The workspace folder path where the settings file is located. |

| Returns          | Type             | Description                                     |
|------------------|------------------|-------------------------------------------------|
| settings         | VSCodeSettings   | The VS Code settings parsed from the settings file. |
| error            | error            | An error, if there's an issue reading or parsing the file. |

These structures and the function help manage and interact with a variety of configurations specific to the VS Code environment, enhancing automation and efficiency within the editor.
