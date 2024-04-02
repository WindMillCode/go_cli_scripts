# Docs for copy_utils.go

## Usage

The `copy_utils.go` file in the `utils` package provides functions for copying files and directories, allowing for deep copying of directory structures and selective file copying based on patterns.

## CopyDir

### Description
`CopyDir` recursively copies a directory along with its subdirectories and files from the source to the destination.

### Usage
```go
err := CopyDir(src, dest)
```

### Reference
| Parameter | Type | Description |
|-----------|------|-------------|
| src | `string` | The source directory path. |
| dest | `string` | The destination directory path. |
| return value | `error` | Error object if an error occurs, otherwise nil. |

## CopyFile

### Description
`CopyFile` copies a single file from the source path to the destination path.

### Usage
```go
err := CopyFile(src, dest)
```

### Reference
| Parameter | Type | Description |
|-----------|------|-------------|
| src | `string` | The source file path. |
| dest | `string` | The destination file path. |
| return value | `error` | Error object if an error occurs, otherwise nil. |

## CopySelectFilesToDestination

### Description
`CopySelectFilesToDestination` copies selected files from the source directory to the destination directory based on a glob pattern.

### Usage
```go
err := CopySelectFilesToDestination(c CopySelectFilesToDestinationStruct)
```

### Reference
| Parameter | Type | Description |
|-----------|------|-------------|
| c | `CopySelectFilesToDestinationStruct` | Struct containing source files, glob pattern, and destination directory. |
| return value | `error` | Error object if an error occurs, otherwise nil. |

### CopySelectFilesToDestinationStruct Fields
| Field | Type | Description |
|-------|------|-------------|
| SourceFiles | `[]string` | Array of source file paths. |
| GlobPattern | `string` | Glob pattern to match files. |
| DestinationDir | `string` | The destination directory path. |

This utility function streamlines the process of copying specific files, enhancing file management operations in Go applications.
