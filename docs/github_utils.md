## [GitHubRelease]
### Usage
`GitHubRelease` is a struct that represents a GitHub release, capturing the tag name and a list of assets available for that release.

### Fields

| Field      | Type         | Description                                      |
|------------|--------------|--------------------------------------------------|
| TagName    | string       | The tag name of the release.                     |
| Assets     | []struct     | A slice of structs representing the assets of the release. Each asset includes the file name and the browser download URL. |

## [GetLatestRelease]
### Usage
`GetLatestRelease` fetches the latest release information from a GitHub repository's release API endpoint.

```go
release, err := GetLatestRelease("https://api.github.com/repos/owner/repo")
if err != nil {
	log.Fatal(err)
}
fmt.Println("Latest release tag:", release.TagName)
```

### Reference

| Parameter   | Type        | Description                                       |
|-------------|-------------|---------------------------------------------------|
| repoURL     | string      | The API URL of the repository to fetch the latest release from. |

| Returns     | Description                                  |
|-------------|----------------------------------------------|
| *GitHubRelease | The latest release information.            |
| error       | An error if the operation fails, otherwise nil. |

## [GetDownloadURLForCurrentOS]
### Usage
`GetDownloadURLForCurrentOS` determines the download URL for the current operating system and architecture from the assets of a release.

```go
url, err := GetDownloadURLForCurrentOS(release)
if err != nil {
	log.Fatal(err)
}
fmt.Println("Download URL for current OS:", url)
```

### Reference

| Parameter   | Type            | Description                                       |
|-------------|-----------------|---------------------------------------------------|
| release     | *GitHubRelease  | The release from which to find the download URL.  |

| Returns     | Description                                  |
|-------------|----------------------------------------------|
| string      | The download URL for the asset.               |
| error       | An error if the operation fails, otherwise nil. |

These functions allow interaction with GitHub's API to fetch release details and determine the appropriate download asset for the current system.
