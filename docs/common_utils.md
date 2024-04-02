# Docs for common_utils.go

## Usage

The `common_utils.go` file in the `utils` package provides a set of general utility functions for Go applications, including type introspection, parameter retrieval, path and branch information, and console screen clearing.

## GetType

### Description
`GetType` returns the data type of a given value as a string.

### Usage
```go
typeName := GetType(value)
```

### Reference
| Parameter | Type | Description |
|-----------|------|-------------|
| value | `interface{}` | The value whose type you want to determine. |
| return value | `string` | The type of the provided value as a string. |

## GetParamValue

### Description
`GetParamValue` returns the provided parameter value or outputs a message if it is nil.

### Usage
```go
paramValue := GetParamValue(parameterName, parameterValue)
```

### Reference
| Parameter | Type | Description |
|-----------|------|-------------|
| parameterName | `string` | The name of the parameter. Used in the output message if the value is nil. |
| parameterValue | `interface{}` | The value of the parameter to check. |
| return value | `interface{}` | The original `parameterValue` if not nil; otherwise, nil. |

## GetCurrentPath

### Description
`GetCurrentPath` retrieves the directory path of the executable.

### Usage
```go
currentPath := GetCurrentPath()
```

### Reference
| Return Value | Type | Description |
|--------------|------|-------------|
| currentPath | `string` | The directory path of the current executable. |

## GetCurrentBranch

### Description
`GetCurrentBranch` obtains the current Git branch name of the working directory.

### Usage
```go
branchName, err := GetCurrentBranch()
```

### Reference
| Return Value | Type | Description |
|--------------|------|-------------|
| branchName | `string` | The current Git branch name. |
| err | `error` | Error, if any occurred during the command execution. |

## ClearScreen

### Description
`ClearScreen` clears the terminal screen.

### Usage
```go
ClearScreen()
```

### Reference
There are no parameters or return values for this function. It simply clears the terminal screen where the Go program is running.
