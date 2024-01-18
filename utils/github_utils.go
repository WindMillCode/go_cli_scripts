package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
)

type GitHubRelease struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func GetLatestRelease(repoURL string) (*GitHubRelease, error) {
	// GitHub API URL for the latest release
	apiURL := repoURL + "/releases/latest"

	// Make the HTTP request
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Unmarshal the JSON data into the GitHubRelease struct
	var release GitHubRelease
	if err := json.Unmarshal(body, &release); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	// Return the entire release struct
	return &release, nil
}

func GetDownloadURLForCurrentOS(release *GitHubRelease) (string, error) {
	if release == nil || len(release.Assets) == 0 {
		return "", fmt.Errorf("no assets found in the release")
	}

	// Determine the current operating system and architecture
	osName := runtime.GOOS
	arch := runtime.GOARCH

	// Define alternate names for architectures
	archMappings := map[string][]string{
		"amd64": {"amd64", "x86_64", "x64"},
		"386":   {"386", "x86"},
		"arm64": {"arm64", "aarch64"},
		"arm":   {"arm"},
	}

	// Get the possible alternate names for the architecture
	archAlternates, ok := archMappings[arch]
	if !ok {
		return "", fmt.Errorf("no alternate names found for architecture: %s", arch)
	}

	// Loop through the assets to find a match for the current OS and one of the architecture alternates
	for _, asset := range release.Assets {
		assetNameLower := strings.ToLower(asset.Name)
		if strings.Contains(assetNameLower, osName) && ContainsAny(assetNameLower, archAlternates) {
			return asset.BrowserDownloadURL, nil
		}
	}

	return "", fmt.Errorf("no zip asset found for the OS: %s and architecture: %s", osName, arch)
}
