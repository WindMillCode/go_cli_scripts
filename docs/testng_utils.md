## [GetTestNGArgsStruct and GetTestNGArgs]
### Usage
`GetTestNGArgs` prompts the user to input various paths and settings for TestNG testing configuration, populating a `GetTestNGArgsStruct` with these inputs.

### Fields in `GetTestNGArgsStruct`

| Field           | Type   | Description                                                       |
|-----------------|--------|-------------------------------------------------------------------|
| WorkspaceFolder | string | The root folder of the workspace.                                 |
| EnvVarsFile     | string | Path to the script setting environment variables for the app.     |
| TestNGFolder    | string | Location of the TestNG application.                               |
| SuiteFile       | string | XML suite file needed for TestNG, relative to the TestNG folder.  |
| ParamEnv        | string | Environment parameter for TestNG.                                 |

### Function: `GetTestNGArgs`
Collects user inputs for TestNG configuration through a series of prompts.

### Reference for `GetTestNGArgs`
| Parameter | Type                | Description                          |
|-----------|---------------------|--------------------------------------|
| c         | GetTestNGArgsStruct | The struct to populate with user inputs. |

| Returns   | Type                | Description                              |
|-----------|---------------------|------------------------------------------|
| c         | GetTestNGArgsStruct | The struct populated with user inputs.   |

### Interactive Prompts in `GetTestNGArgs`
1. **EnvVarsFile:** Prompt for the script where environment variables are set, defaulting to a path within the workspace.
2. **TestNGFolder:** Prompt for the TestNG application location, defaulting to a path within the workspace.
3. **SuiteFile:** Prompt for the XML suite file path relative to the TestNG folder, defaulting to a typical TestNG resource path.
4. **ParamEnv:** Uses `ShowMenu` to allow the user to select an environment to test, with options for "DEV", "PREVIEW", and "PROD", and an option to specify a different environment.

This function streamlines the configuration process for TestNG tests by gathering necessary file paths and environment settings from the user, facilitating a more dynamic and interactive setup process.
