## [GitSparseClone]
### Usage
`GitSparseClone` performs a sparse clone of a Git repository, cloning only specified subdirectories into a local directory.

```go
err := GitSparseClone("https://github.com/user/repo.git", "/local/dir", "subdir1", "subdir2")
if err != nil {
    log.Fatal(err)
}
```

### Reference

| Parameter       | Type         | Description                                       |
|-----------------|--------------|---------------------------------------------------|
| repoURL         | string       | The URL of the Git repository to clone.           |
| localDir        | string       | The local directory where the repository will be cloned. |
| subdirectories  | ...string    | The subdirectories to include in the sparse clone.|

| Returns         | Description                                  |
|-----------------|----------------------------------------------|
| error           | An error if the operation fails, otherwise nil. |

## [ExtractBranchNames]
### Usage
`ExtractBranchNames` extracts branch names from a string, typically used to parse the output of a Git fetch or pull command.

```go
branches := ExtractBranchNames("Fetching origin\n[new branch]   master -> origin/master")
fmt.Println(branches) // Outputs: ["master"]
```

### Reference

| Parameter       | Type         | Description                                       |
|-----------------|--------------|---------------------------------------------------|
| input           | string       | The string input containing Git branch information.|

| Returns         | Description                                  |
|-----------------|----------------------------------------------|
| []string        | An array of extracted branch names.          |

These functions provide utilities for working with Git repositories, such as performing a sparse clone and extracting branch names from command output.
