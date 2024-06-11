package update

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/mrtnhwttktc/kinto-cli/internal/localizer"
	v "github.com/mrtnhwttktc/kinto-cli/internal/version"

	"github.com/spf13/cobra"
)

func NewUpdateCmd() *cobra.Command {
	l := localizer.NewLocalizer()
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: l.Translate("Update the CLI to the latest version."),
		Long:  l.Translate("Update the CLI to the latest version. This command will check for the latest version of the CLI and update it if necessary."),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Check if the command was run with sudo
			if os.Geteuid() != 0 {
				out := cmd.OutOrStdout()
				fmt.Fprintln(out, l.Translate("Please run this command with sudo."))
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			out := cmd.OutOrStdout()
			// get where the binary is stored
			binaryPath, err := os.Executable()
			if err != nil {
				fmt.Fprintln(out, l.Translate("Failed to get binary path, please provide it using the --path flag."))
				os.Exit(1)
			}

			// Check if the binary path is a symlink
			err = checkBinaryPath(binaryPath)
			if err != nil {
				slog.Error("Error when checking for symlink.", slog.String("error", err.Error()))
				fmt.Fprintln(out, l.Translate("Failed to check binary path, please provide the actual binary path using the --path flag."))
				os.Exit(1)
			}

			dlBinPath, err := v.DownloadBin()
			if err != nil {
				slog.Error("Error when downloading the latest version of the CLI.", slog.String("error", err.Error()))
				fmt.Fprintln(out, l.Translate("Failed to download the latest version of the CLI."))
				os.Exit(1)
			}
			defer os.Remove(dlBinPath)

			err = addExecutePermission(dlBinPath)
			if err != nil {
				slog.Error("Error when adding execute permission to the new binary.", slog.String("error", err.Error()))
				fmt.Fprintln(out, l.Translate("Failed to add execute permission to the new binary."))
				os.Exit(1)
			}

			err = replaceBinary(binaryPath, dlBinPath)
			if err != nil {
				slog.Error("Error when replacing the current binary with the new one.", slog.String("error", err.Error()))
				fmt.Fprintln(out, l.Translate("Failed to replace the current binary with the downloaded binary."))
				os.Exit(1)
			}

			fmt.Fprintln(out, l.Translate("CLI updated successfully. Run `ktcli -v` to check the new version."))
		},
	}
	setFlags(updateCmd, l)
	return updateCmd
}

func addExecutePermission(binaryPath string) error {
	err := os.Chmod(binaryPath, 0755)
	if err != nil {
		return fmt.Errorf("failed to add execute permission to the new binary: %w", err)
	}
	return nil
}

func replaceBinary(binaryPath, newBinaryPath string) error {
	err := os.Rename(newBinaryPath, binaryPath)
	if err != nil {
		return fmt.Errorf("failed to replace the current binary with the new one: %w", err)
	}
	return nil
}

func checkBinaryPath(binaryPath string) error {
	// Check if the binary path is a symlink
	realPath, err := filepath.EvalSymlinks(binaryPath)
	if err != nil {
		return err
	}
	if realPath != binaryPath {
		return fmt.Errorf("binary path is a symlink")
	}
	return nil
}

func setFlags(cmd *cobra.Command, l *localizer.Localizer) {
	cmd.Flags().String("path", "", l.Translate("Path to the binary to update. Defaults to the current binary path."))
}
