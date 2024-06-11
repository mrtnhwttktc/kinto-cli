package version

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

// Version is the current version of the CLI.
// Updated by the build process
var Version = "dev"
var BinURL = ""

// Check if a new version is available.
func IsNewVersionAvailable() bool {
	resp, err := http.Head(BinURL)
	if err != nil {
		slog.Warn("Error checking for new version. Head request returned an error.", slog.String("error", err.Error()))
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		slog.Warn("Error checking for new version. Received a non 200 response code.", slog.String("status", resp.Status))
		return false
	}
	metadataVersion := resp.Header.Get("x-amz-meta-version")
	if metadataVersion == "" {
		slog.Warn("Error checking for new version. No version header found.")
		return false
	}
	if metadataVersion != Version && Version != "dev" {
		return true
	}
	return false
}

func DownloadBin() (string, error) {
	tmpFile, err := os.CreateTemp("", "ktcli")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer tmpFile.Close()

	resp, err := http.Get(BinURL)
	if err != nil {
		return "", fmt.Errorf("failed to download latest version: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download latest version: %s", resp.Status)
	}

	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to write to temporary file: %w", err)
	}

	return tmpFile.Name(), nil
}
