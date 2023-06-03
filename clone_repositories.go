package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

func cloneRepositories(config *Config, downloadDir string) error {
	var wg sync.WaitGroup

	// Check if download directory exists, create it if it doesn't
	if _, err := os.Stat(downloadDir); os.IsNotExist(err) {
		err := os.MkdirAll(downloadDir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create download directory: %v", err)
		}
	}

	for _, repo := range config.Repositories {
		// Parse the repo URL into a URL type
		url, err := url.Parse(repo)
		if err != nil {
			return fmt.Errorf("failed to parse repo URL: %v", err)
		}

		// skip existing repo
		if isWithin(url, downloadDir) {
			continue
		}

		wg.Add(1)
		go func(repo string) {
			defer wg.Done()
			err := cloneRepo(url, downloadDir, config.Token)
			if err != nil {
				log.Printf("Failed to clone repository %s: %v\n", repo, err)
			}
		}(repo)
	}

	wg.Wait()

	return nil
}

func cloneRepo(u *url.URL, downloadDir, token string) error {
	fmt.Printf("Cloning repository: %s\n", u)

	// Construct the new URL with the token and modified repo path
	newRepo := fmt.Sprintf("%s://%s:x-oauth-basic@%s", u.Scheme, token, u.Host+u.Path)

	cmd := exec.Command("git", "clone", newRepo)
	cmd.Dir = downloadDir
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("clone error: %v, output: %s", err, string(output))
	}

	fmt.Printf("Repository cloned: %s\n", u)
	return nil
}

func isWithin(u *url.URL, downloadDir string) bool {
	repoPath := filepath.Join(downloadDir, filepath.Base(u.Path))
	return strings.HasPrefix(repoPath, downloadDir)
}
