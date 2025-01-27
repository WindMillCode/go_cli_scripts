# Overivew

* useful go utility functions


# Changelog

## v1.0.0
* utils.GetItemsInFolder now has just one argument for convenience

## v1.0.1

* added FilterArray which makes it more convenient to filter an array given an array and predicate(condition function)
* added methods that would cast the interface to the appropriate type

## v1.0.2
* made a slight change

## v1.0.3
* added index to element

## v1.1.0
* added FilterMapByKeys all  the respective casting fns

## v1.1.1
* added func OverwriteMap which is like js Object.assign

## v1.1.2
* added func OverwriteMap which is like js Object.assign

## v2.0.0
* replaced FilterMapByKeys with FilterMap where its predicate fn accepts key,val


## v2.0.1
* added FilterJSONByPredicate

## v2.0.2
* added WriteCustomFormattedJSONToFile to help format files

## v2.0.5
* WriteCustomFormattedJSONToFile supports bytes and interface strcutures, added UnicodeUnquote which will remove all unicode from a string when writing from bytes to a file,

## v2.0.6
* added AddContentToEachLineInFile AddContentToFile, which takes a predicate function and updates the file based on  the return of the predicate fn

## v2.1.1
* added MergeDirectories fn which would merge all files and folders from target dir into source dir w/o overrting anything

## v2.1.2
* added RunCommandWithOptions fn which supports optional target dir,optional get output, and panic on error

## v2.1.3
* added TraverseDirectory with a predicate fn

## v2.1.4
* added TruncateStringByRegex fn which allows the end user to provide a regex and has a predicate fn for every match in the pattern matcher which if returns true removes the substr from the array


## v2.1.5
* additional updates

## v3.0.0
* added cli and changed name to go_cli_scripts

## v3.1.0
* updated RunCommandWithOptions to print out standard err along with the reason why the command failed

## v3.1.1
* indicated all RunCommand fns are deprecated and RunCommandWithOptions should be used instead

## v3.1.2
* added ProcessFoldersMatchingPattern just like ProcessFilesMatchingPattern



## v3.1.5
added CreateStringObject an object that will give all sorts of cases
  camelCase
  kebab
  snakeCase
  classify
* fixied CreateStringObject

## v3.1.6
added IsRunningInDocker to see whether the given go script is running in a docker container or not

## v3.2.0
* [PATCH] updated GetInputFromStdin and ShowMenu to immediately cancel when the user hits Ctrl+C

## v3.2.1 [1/5/24]
* [PATCH] build mechanism seems to unexpectedly add unwanted code this patch should fix that

## v3.3.0 [1/6/24]
* [PATCH] fixed issue with RunCommandWithOptions where output gets returned from the program and printed to console
* added Uppercase method to CreateStringObjectType

## v3.3.1 [1/11/24]
* [PATCH] IsRunningInDocker will able to detect for macbooks as well

## v3.3.3 [1/17/24]
* [UPDATE] added 	ProxyURLs string `json:"proxyURLs"`
for the VSCodeSettings.ExtensionPack struct

## v4.0.2 [1/17/24]
* [BREAKING CHANGE] TakeVariableArgs returns TakeVariableArgsResultStruct
```go
type TakeVariableArgsResultStruct struct{
	inputString  string
	inputArray   []string
}
```

## v4.1.0 [1/17/24]
* added github_utils and

* file_utils
FindExecutable
GetSourceFilePath
ExtractArchive
DownloadFile

* string_utils
ContainsAny

methods

## v4.2.0 [1/18/24]
* [MAJOR UPDATE]
finally got RunCommandWithOptions to work as intended
stdout prints to terminal in addition to returning output from the inner running program to the main program

## v4.2.1 [1/18/24]
* [UPDATE] added PrintOutput to CommandOptions to optionally prevent output from getting to the command line

## v4.2.4 [1/18/24]
* [UPDATE] added PrintOutputOnly to CommandOptions by default stdOut. returns the final value of the program to the outer command running it
and prints output to the command line, because of edge case this was added to only print output to the command line for convience sake
* [UPDATE] added TasksToRunOnFolderOpen for the vscode extension

## v4.3.4 [1/24/24]

* [NEW FEATURE] Enhanced `WatchDirectory` function for dynamic directory monitoring with advanced filtering options.

    ```go
    type WatchDirectoryParams struct {
        DirectoryToWatch string
        DebounceInMs     int
        Predicate        func(event fsnotify.Event)
        StartOnWatch     bool
        IncludePatterns  []string
        ExcludePatterns  []string
    }
    ```

    - The `WatchDirectory` function now includes parameters for include and exclude patterns using glob strings, allowing for more precise monitoring of directory changes.
    - Added `StartOnWatch` boolean parameter to immediately invoke the predicate function on the existing files in the directory at the start.
    - The include and exclude patterns are compiled using `CompileGlobs` and checked against each file event using `MatchAnyGlob`.

* [UPDATE] `CommandOptions` struct enhancements for process control.

    ```go
    type CommandOptions struct {
        // ... existing fields ...
        CmdObj      *exec.Cmd
        NonBlocking bool
    }

    func (c CommandOptions) EndProcess() error {
        return c.CmdObj.Process.Kill()
    }
    ```

    - Enhanced `CommandOptions` to include `CmdObj` for direct process handling and `NonBlocking` boolean for asynchronous command execution.
    - Added `EndProcess` method to facilitate immediate termination of the process associated with `CmdObj`.


## v4.4.1 [1/24/24]
* [UPDATE] - added ShowMenuMultiple to allow the user to select multiple options along with
	SelectionLimit        int
	SelectedValues        []string
	SelectedDelimiter     string
properties
* [UPDATE] - added RemoveContentFromFile function
* [FIX] - fixed EndProcess command to leverage the pointer to Self on the options.CommandOptions struct via the Self property
* you the developer have to provide the self property in order for EndProcess to find the command and  kill it effectively

## v4.4.2 [1/28/24]
[PATCH] - fixed issue with EndProcess where it would kill the process but hang in the function it was called
* a glitch allows this fix to be found in 4.4.1


## v4.4.3 [1/30/24]
* [UPDATE] - added RemoveElementsNotInSource ArrayContainsAny array fns

## v4.4.4 [2/29/24]
* [UPDATE] - updated CDToLocation to accept a createIfNotExist parameter (2 argument)

## v4.5.0 [3/9/24]
* [UPDATE] - Added a new task "major update project" in .vscode/tasks.json.
* [FEATURE] - Introduced a new utility file docker-utils.go with a function to check if running inside Docker.
* [PATCH] - Updated vscode_settings_utils.go to use yaml.Unmarshal instead of json.Unmarshal.
* [REMOVE] - Removed redundant IsRunningInDocker function from vscode_settings_utils.go.

## v4.5.1 [3/9/24]
* [UPDATE] - Added a new dependency github.com/tailscale/hujson in go.mod and updated go.sum accordingly.
* [FEATURE] - Introduced a new function StandardizeJSON in utils/json_utils.go using hujson to standardize JSON input.
* [FIX] - Modified GetSettingsJSON in utils/vscode_settings_utils.go to use StandardizeJSON for processing input before unmarshalling.

## v4.5.4 [3/10/24]
* [REMOVE] - Removed the dependency github.com/tailscale/hujson from go.mod.
* [UPDATE] - Updated go.sum to reflect changes in dependencies.
* [UPDATE] - Modified the StandardizeJSON function in utils/json_utils.go to RemoveComments, which now removes line and block comments from JSON strings.
* [UPDATE] - Updated GetSettingsJSON in utils/vscode_settings_utils.go to use RemoveComments instead of StandardizeJSON for processing input before unmarshalling.
* [FIX] - Added error handling in RemoveComments to check for JSON validity after removing comments.

## v4.5.5 [3/28/24]
[REMOVE] `output.txt` and `scratchpad.go` files were deleted. The `output.txt` file contained various utility functions and `scratchpad.go` included main package content with file operations.

[UPDATE] `copy_utils.go` and `file_utils.go` in the `utils` directory were updated. In `copy_utils.go`, comments were removed, and error handling was improved for directory copying logic. In `file_utils.go`, a new function `FilterFilesByExtension` was added to filter files based on their extensions.

## v5.0.0 [4-2-2024 9:22:00 AM EST]

[UPDATE] Made a big change in the README file. It's now an overview of what the go_cli_scripts package does instead of a changelog.
[BREAKING CHANGE] In cd_utils.go, tweaked the path in CDToWorkspaceRoot to go up three levels instead of four. Just a small path correction.
[FIX] In file_utils.go, corrected a message that was supposed to say "Error writing to file" instead of "Error reading from file". Also added two new handy functions: RemovePathPrefix (gets rid of the start of a path) and EnsureDirAndCreateFile (makes sure a folder exists before creating a file in it).


## v5.0.3  [4/11/2024 7:14:00 PM EST]

- [UPDATE] We've added a new `environments` array in the `WindmillcodeExtensionPack` in your settings. This is in the `vscode_settings_utils.go` file. It's like a new list where you can keep track of different settings for various environments you work in. Also added `sentryDSN`

- [FIX] Squashed a tiny bug in `file_utils.go`! Now, when you're filtering files, it won't ignore the ones with `_test.dart` at the end. This means all your Dart files get the attention they deserve, even the test ones!

## v5.1.0 [4/14/2024 8:47:00 AM EST]

[UPDATE] Added FlutterMobileBuild struct with new settings to WindmillcodeExtensionPack. Now it includes args, toolArgs, and vmAdditionalArgs for configuring Flutter mobile builds. Check out vscode_settings_utils.go to see the changes.

## v5.1.1 4/14/2024 5:47:00 PM EST

[UPDATE] Yo! We just added a sweet new function FindRelativeToTarget in array_utils.go. It helps you find a string in an array that's right before or after the one you're looking for. Pretty handy, huh?

[UPDATE] Heads up! The RunCommand function in stdio_utils.go is no longer deprecated. Use it if you need a shorter version of run command wit hoptions

[UPDATE] Check this out! We’ve beefed up our string tools in string_utils.go. Now you can turn your strings into kebab-case with the new KebabCase method in CreateStringObjectType. Keep your strings looking sharp!

## v5.1.2 4/15/2024 10:00:23 AM EST

[UPDATE] Yo! We just beefed up our FlutterMobileBuild struct in the vscode_settings_utils.go file. Now it includes fields for PlayStoreServiceAccountKey, PackageName, PublishTarget, and TrackName. If you're working on publishing to the Play Store, these new fields will help configure things right from your code!


## v5.1.3 4/24/2024 12:45:23 AM EST

[UPDATE] Extended programming language support in WindmillcodeExtensionPack within utils/vscode_settings_utils.go. Added new fields for NodeJS, Java, Go, Ruby, Dart, CSharp, Swift, PHP, Rust, Kotlin, Scala, Perl, Lua, Haskell, Clojure, Erlang, Julia, ObjectiveC, FSharp, VisualBasic versions. This update will require developers to adjust any related configurations that depend on these language version settings.

## v5.2.1 [5/29/2024 10:17:45 AM EST]

[UPDATE] Removed commented-out code from `examples/fast_immutable_collections.go`, `examples/flutter_translate.go`, `examples/go_cli_scripts_create_output.go`, `examples/json_serilizable_create_output.go`, and `examples/webview_flutter.go`. The unnecessary comments are gone, making the files cleaner and easier to read.

[UPDATE] Cleaned up `examples/output.md` by removing all irrelevant, offensive, and unnecessary content. The file is now streamlined and professional.

[FIX] Fixed error handling in `utils/vscode_settings_utils.go` by replacing `fmt.Println` with `LogErrorWithTraceBack` in the `GetSettingsJSON` function. This ensures better logging and error tracking.

[UPDATE] Added `LogErrorWithTraceBack` function in `utils/common_utils.go` to log errors with traceback information. This new function helps in debugging by providing detailed error logs.

[PATCH] Changed the `Filter` function in `examples/webview_flutter.go` to remove redundant spaces and clean up the code.

[UPDATE] Added `log` and `runtime/debug` imports in `utils/common_utils.go` for enhanced error logging.

[BUG] Fixed the file reading and JSON unmarshalling errors in `GetSettingsJSON` by adding detailed logging in `utils/vscode_settings_utils.go`. Now it logs errors with traceback for easier debugging.

## v5.2.2 [5/29/2024 11:42:00 AM EST]

[PATCH] utils/file_utils.go: Added ReadLines function to read lines from a file into a map.

## v5.2.3
[UPDATE] -added so users can  specifiy their own paths
```go
type GitPushingWorkToGitRemoteStruct struct {
	RelativePaths []string `json:"relativePaths"`
	AbsolutePaths []string `json:"absolutePaths"`
}
```

## v5.3.1
[6/5/2024 10:34:22 AM EST]

[UPDATE] Added utility functions in `stdio_utils.go` and `string_utils.go` to enhance command execution and argument handling.

**File:** `utils/stdio_utils.go`
- **Function Added:** `RunElevatedCommand`
  - Description: Run commands with elevated privileges (admin/sudo) based on the OS.
  - Developer Impact: Now you can easily execute commands that need admin privileges without writing additional code.

**File:** `utils/string_utils.go`
- **Function Added:** `JoinArgs`
  - Description: Joins command arguments into a single string suitable for command-line execution.
  - Developer Impact: Simplifies the process of formatting arguments for command-line execution, particularly useful for the new `RunElevatedCommand` function.


[UPDATE] Adjusted error handling in `RunCommandWithOptions` to include detailed stderr output in case of failure.

**File:** `utils/stdio_utils.go`
- **Function Modified:** `RunCommandWithOptions`
  - Description: Now includes detailed standard error output in error messages.
  - Developer Impact: Easier debugging and error tracing when commands fail.


[FIX] Refined JSON handling functions to remove comments and filter based on predicates.

**File:** `utils/json_utils.go`
- **Functions Modified:**
  - `RemoveComments`
  - `FilterJSONByPredicate`
  - Developer Impact: Improved JSON processing capabilities, ensuring cleaner data and more flexible filtering options.

## v5.4.0
[7/9/2024 3:15:00 PM EST]

* Updated dependencies

## v5.4.2
7/25/2024 04:15:30 PM EST

[UPDATE] Added CDToShopifyApp function
- **File**: utils/cd_utils.go
- **Function**: CDToShopifyApp
- **Details**: Added a new function to change directory to the ShopifyApp location. Developers working with ShopifyApp can now easily navigate using this function.

[UPDATE]
Added `FindElement` function to `utils/array_utils.go`
  - This function helps find an element in an array based on a predicate
  - Returns index, element, and an error if the element is not found


## v5.5.0
7/28/2024 06:15:00 PM EST


Updated `CHANGELOG.md`

[ADDED]
Added:
- CDToShopifyApp function to `utils/cd_utils.go`

[UPDATED]
Updated:
- `FindElement` function to `utils/array_utils.go`
  - This function helps find an element in an array based on a predicate
  - Returns index, element, and an error if the element is not found


Updated `utils/array_utils.go`

[ADDED]
Added:
- `FindElement` function to find an element in an array
- `RemoveDuplicates` function to remove duplicate elements from an array
- New conditions to `ConvertToStringArray` to handle integer types

Updated `utils/file_utils.go`

[ADDED]
Added:
- `ReadFileFromPackage` function to read a file from package directory
- `GetFilePathFromPackage` function to get the absolute path of a file in the package

Updated `utils/stdio_utils.go`

[ADDED]
Added:
- `KillPorts` function to kill processes running on specified ports
- Expanded `CommandOptions` struct to include `EnvVars` and `IsElevated`


[ADDED]
Added:
- `IntToStr` function to convert an integer to a string

Updated `utils/vscode_settings_utils.go`

[ADDED]
- `WindmillcodeExtensionPack` struct to include new propertiesAdded:

- `settings.WindmillcodeExtensionPack.FirebaseCloudRunEmulatorsStruct` for Firebase Cloud Run emulator settings
- `settings.WindmillcodeExtensionPack.ProcessIfDefaultIsPresentStruct` to handle default process settings
- `settings.WindmillcodeExtensionPack.WMLPorts` struct to define various ports
- New methods to `WMLPorts` to get Firebase ports


## v5.6.0
[7/29/2024 4:15:00 PM EST]

[UPDATE] Added `nonInteractive` field to `ShowMenuModel` and `ShowMenuMultipleModel` structs in `utils/show_menu.go`.

[PATCH] Modified `ShowMenu` and `ShowMenuMultipleOptions` functions in `utils/show_menu.go` to use default choices in non-interactive mode.

[UPDATE] Added `SetGlobalVarsOptions` struct and `SetGlobalVars` function in `utils/stdio_utils.go` for setting global variables.

[PATCH] Updated `TakeVariableArgs`, `GetInputFromStdin`, and other input functions in `utils/stdio_utils.go` to handle non-interactive mode based on global variables.

[UPDATE] Added `ShopifyRunStruct` to `vscode_settings_utils.go` and included it in `WindmillcodeExtensionPack` struct.

[PATCH] Updated `FirebaseCloudRunEmulatorsStruct` in `vscode_settings_utils.go` to include `KillPortOutputFileAcceptDefault` field.

These changes improve handling of non-interactive modes across various functions and structs. Developers should now set global variables using `SetGlobalVars` and can use defaults in non-interactive mode by leveraging these updates.


## v5.6.1 7/29/2024 03:15:45 PM EST

[UPDATE] Made some JSON fields optional in Firebas
eConfig in main.go

## v5.6.2 7/31/2024 11:27:30 AM EST

[UPDATE]
- console log message optimizations

## v5.6.3 8/31/2024 03:58:12 PM EST

[UPDATE] Added a new function `ParseJSONFromString[T any]` in `utils/json_utils.go`. This function makes it easier to parse JSON strings into Go structs, and it includes error handling to catch JSON parsing issues. If you're working with JSON in Go, this is a handy new tool to use.

## v5.6.4 [11/29/2024 11:23:45 AM EST]

[UPDATE] Added a new function `ParseJSONFromString[T any]` in `utils/json_utils.go`. This function makes it easier to parse JSON strings into Go structs. It also handles errors when the JSON format is bad. If you're dealing with JSON in Go, this will save you time.

[UPDATE] Added a new function `CDToWxtApp` in `utils/cd_utils.go`. It changes the directory to the location of the `WxtApp`. Use this when working with apps/extensions under the `apps/extensions/WxtApp` path.

[UPDATE] Updated `WindmillcodeExtensionPack` in `utils/vscode_settings_utils.go`. Added two new fields:
- `WxtBuildSafari`: Contains `BundleIdentifier` for Safari builds.
- `AngularDeployToFirebase`: Includes settings for deploying Angular apps to Firebase and Sentry.

`AngularDeployToFirebase` has options like:
- **RunLint**: Run lint checks before deployment.
- **RunSSGScript**: Run the SSG script for static site generation.
- **RemoveBuildDirectories**: Clears old build folders.
- **DeployToSentry**: Pushes source maps to Sentry (requires `SentryOrg`, `SentryProject`, and `SentryAuthToken`).
- **DeployToFirebase**: Deploys the app to Firebase.

If you're using these features, update your settings JSON to include these fields.


## v5.6.5 [11/29/2024 12:23:45 AM EST]

[UPDATE] `AngularDeployToFirebase` has an additional option :
	FirebaseProjectId             string    `json:"firebaseProjectId"`
  to specify the firebase projec

## v5.7.0 [11/29/2024 16:23:45 PM EST]

[UPDATE]
Renamed `RemoveComments` function to `CleanJSON` in `utils/json_utils.go`. It now:
1. Removes comments (both `//` and `/* */`).
2. Removes trailing commas from arrays and objects.
3. Throws better error messages when the JSON is invalid.


## v6.0.0 [12/1/2024 12:50:00 PM EST]

[BREAKING CHANGE] Changed `cd_utils.md` to rename **CDToTestNGApp** to **CDToSeleniumApp** for clarity. Updated the description to match the function's purpose.

[FIX] Renamed `CDToTestNGApp` to `CDToSeleniumApp` in `cd_utils.go`. Updated the directory path to `"apps/testing/SeleniumApp"`.

[FIX] Improved the `TraverseDirectory` function in `file_utils.go` to skip missing paths gracefully. Added a check for `os.IsNotExist` errors and adjusted the behavior to print a message and continue processing.

[UPDATE] Added the `ReplaceAllSubstrings` function in `string_utils.go`. Developers can use this to replace all occurrences of a target substring in a string. It includes validation to ensure the target substring is not empty.

[PATCH] Enhanced the `Capitalize` function in `string_utils.go` to handle cases where the input string is empty. Returns only the suffix in such cases. Adjusted logic to properly capitalize the first letter of non-empty strings.

[UPDATE] Added `omitempty` tags to multiple struct fields in `vscode_settings_utils.go` for improved JSON serialization. This change reduces unnecessary fields in serialized outputs when values are empty.

[BUG] Fixed directory paths in `testng_utils.go` to ensure compatibility with OS-specific formats. Updated the `EnvVarsFile` and `TestNGFolder` default paths using `JoinAndConvertPathToOSFormat`.

[PATCH] Added a new struct `MiscReinitializeProjectStruct` in `vscode_settings_utils.go` to support reinitialization logic for projects. Includes fields like `ProjectName`, `OrganizationName`, and various platform-specific configuration options.

## v6.0.1 [12/7/2024 12:03:22 PM EST]

[UPDATE] Updated the `go get` command in the `README.md` file from `v5` to `v6` for the `go_cli_scripts` package. If you're integrating the library, make sure to use the updated version `v6`.

[UPDATE] Updated import statements in the `README.md` file from `v5` to `v6` for the `go_cli_scripts` package. Developers need to replace `v5` with `v6` in their imports.

[CHECKPOINT] Added a new function `CDToReactNativeExpoApp` in `utils/cd_utils.go` that lets you navigate to the React Native Expo app directory. If you're working on this project type, you can now use this utility function.

[PATCH] Introduced a new file `utils/vscode_tasks_utils.go` that handles creating and reading `tasks.json` for VSCode. If you're dealing with VSCode tasks, this file has utilities to help generate or read dynamic task configurations.

[BUG] Fixed an issue in `utils/vscode_tasks_utils.go` where the `CreateTasksJson` function might fail if the `tasks.json` file didn’t exist. The function now creates the file if it’s missing. Developers can now safely use this utility without manually creating the file.

## v6.0.2 [12/13/2024 01:47:00 PM EST]

[UPDATE] Added a new function `CDToLaravelApp` in `utils/cd_utils.go`. It helps developers quickly change directories to the Laravel app folder at `apps/backend/LaravelApp`. Super handy for Laravel developers.


[PATCH] Fixed `ProxyURLs` in `MiscReinitializeProjectStruct` in `utils/vscode_settings_utils.go`. It will now always include `ProxyURLs`, even if it’s empty. This ensures consistency for developers relying on this property during JSON handling.

## v6.0.3 [12/14/2024 10:45:23 AM EST]

[UPDATE]
Added two new properties `MySQL0`, `ReactNativeExpoRun0`, and `LaravelRun0` to the `WMLPorts` struct in `utils/vscode_settings_utils.go`. Developers can now specify ports for MySQL, React Native Expo, and Laravel in their JSON configurations.

[CHECKPOINT]
This update extends the WMLPorts struct to support more frameworks and services. Be sure to update any dependent JSON files to include these new properties if needed.

## v6.0.4 [1/13/2025 1:43:12 PM EST]

[UPDATE]
**File:** `utils/stdio_utils.go`
**What changed:** Added a new property `ExitRegex` in the `CommandOptions` struct.
**Why it matters:** Developers can now specify to decide when a process should be terminated

## v6.0.6  [1/23/2025 4:53:06 PM EST]

[UPDATE] Updated `go.mod` to require Go version 1.23. This means you'll need Go 1.23 or newer to use this module. ([Go](https://go.dev/doc/modules/gomod-ref?utm_source=chatgpt.com))

[DEPRECATED] Added a note that `ExtractArchive` function in `utils/file_utils.go` is now deprecated. You should start using `ExtractArchiveWithOptions` instead. ([HashiCorp Developer](https://developer.hashicorp.com/terraform/plugin/sdkv2/best-practices/deprecations?utm_source=chatgpt.com))

[UPDATE] Introduced a new function `ExtractArchiveWithOptions` in `utils/file_utils.go`. This function gives you more flexibility when extracting archives.

## v6.0.7 [1/25/2025 10:45:12 AM EST]

[UPDATE] Added new structs `NPMInstallAppDepsStruct`, `NPMInstallSpecifcPackagesStruct`, `PythonInstallAppDepsStruct`, and `PythonInstallSpecifcPackagesStruct` to the file `utils/vscode_settings_utils.go`. These structs handle app dependencies and package installations, giving developers more granular control over settings.

[UPDATE] Updated `WindmillcodeExtensionPack` in `utils/vscode_settings_utils.go` to include properties for managing `PythonInstallAppDeps`, `PythonInstallSpecifcPackages`, `NPMInstallSpecifcPackages`, and `NPMInstallAppDeps`. Developers now have fields to configure these in their settings.

[UPDATE] In `GetSettingsJSON` function, added initialization for `NodeJSAppLocations` and `PythonAppLocations` if they are `nil`. This ensures default paths are set for Node.js and Python apps when not provided, saving developers time in configuring paths manually.

## v6.0.8 [1/25/2025 6:42:15 PM EST]

[UPDATE] Added a new `PascalCase` property to the `CreateStringObjectType` struct in `utils/string_utils.go`. Developers can now generate PascalCase strings using the `CreateStringObject` function.

## v6.0.9 [1/27/2025 10:34:15 AM EST]

[UPDATE]
Added a new struct `MiscTranslateJsonStruct` to `vscode_settings_utils.go`. If you're dealing with translations, this struct is mapped in `WindmillcodeExtensionPack`. Developers can now leverage `MiscTranslateJson` for JSON translation features.

