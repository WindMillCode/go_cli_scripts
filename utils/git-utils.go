package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)


func GitSparseClone(repoURL string, localDir string, subdirectories ...string) error {
	// Create the local directory if it doesn't exist
	if err := os.MkdirAll(localDir, os.ModePerm); err != nil {
		return err
	}

	// Change to the local directory
	if err := os.Chdir(localDir); err != nil {
		return err
	}

	// Initialize a Git repository
	cmd := exec.Command("git", "init")
	if err := cmd.Run(); err != nil {
		return err
	}

	// Add a remote and fetch
	cmd = exec.Command("git", "remote", "add", "-f", "origin", repoURL)
	if err := cmd.Run(); err != nil {
		return err
	}

	// Configure sparse checkout
	cmd = exec.Command("git", "config", "core.sparseCheckout", "true")
	if err := cmd.Run(); err != nil {
		return err
	}

	// Write subdirectories to .git/info/sparse-checkout
	sparseFile := filepath.Join(".git", "info", "sparse-checkout")
	file, err := os.Create(sparseFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, subdir := range subdirectories {
		fmt.Fprintln(file, subdir)
	}

	// Pull from the remote repository
	cmd = exec.Command("git", "pull", "origin", "master")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func ExtractBranchNames(input string) []string {
	fmt.Printf(input)
	var branchNames []string

    lines := strings.Split(input, "[new branch]")

    for _, line := range lines {
        parts := strings.Fields(line)
        if len(parts) >= 4 && parts[3] == "->" {
            branchName := strings.TrimSpace(parts[2])
            branchNames = append(branchNames, branchName)
        }
    }

    return branchNames
}
