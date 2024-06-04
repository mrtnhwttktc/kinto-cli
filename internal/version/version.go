package version

import (
	"log/slog"
	"net/http"
)

// Version is the current version of the CLI.
// Updated by the build process
var Version = "dev"
var binURL = ""

// Check if a new version is available.
func IsNewVersionAvailable() bool {
	resp, err := http.Head(binURL)
	if err != nil {
		slog.Warn("Error checking for new version. Head request returned an error.", slog.String("error", err.Error()))
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		slog.Warn("Error checking for new version. Received a non 200 response code.", slog.String("status", resp.Status))
		return false
	}
	metadataVersin := resp.Header.Get("x-aws-version")
	if metadataVersin == "" {
		slog.Warn("Error checking for new version. No version header found.")
		return false
	}
	if metadataVersin != Version && Version != "dev" {
		return true
	}
	return false
}
