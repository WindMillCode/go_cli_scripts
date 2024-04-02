# Docs for cd_utils.go

## Usage

The `cd_utils.go` file in the `utils` package provides functions to change the current working directory in a Go program, with options to create the target directory if it does not exist.

## CDToLocation

### Description
`CDToLocation` changes the current working directory to the specified location. It can optionally create the directory if it doesn't exist.

### Usage
```go
CDToLocation(location, opts...)
```

### Reference
| Parameter | Type | Description |
|-----------|------|-------------|
| location | `string` | The path to the directory to change to. |
| opts | `...interface{}` | Optional arguments, where the first argument can be a boolean to specify whether to create the directory if it does not exist. |

### Additional Functions
- **CDToWorkspaceRoot**: Navigates to the root directory of the workspace. (meant to be used in a Windmillcode project its highly recommneded to make your own this is very unstable right now in the future this script may walk up the file tree till it finds the .windmillcode folder )
- **CDToAngularApp**: Changes the directory to the Angular application within a predefined project structure.
- **CDToFirebaseApp**: Navigates to the Firebase application directory.
- **CDToFlaskApp**: Changes the current working directory to the Flask backend application.
- **CDToTestNGApp**: Sets the current directory to the TestNG application's location.
- **CDToFlutterApp**: Moves the working directory to the Flutter mobile application's location.

Each of these functions leverages `CDToLocation` to change the working directory, providing a convenient way to navigate through a project's directory structure programmatically.
