package update

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mrtnhwttktc/kinto-cli/internal/localizer"
	"github.com/spf13/cobra"
)

// UpdateCmd represents the update command

func NewUpdateCmd() *cobra.Command {
	l := localizer.GetLocalizer()
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: l.Translate("Update the CLI to the latest version."),
		Long:  l.Translate("Update the CLI to the latest version. This command will check for the latest version of the CLI and update it if necessary."),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Check if the command was run with sudo
			if os.Geteuid() != 0 {
				fmt.Println(l.Translate("Please run this command with sudo."))
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			// get where the binary is stored
			binaryPath, err := os.Executable()
			if err != nil {
				fmt.Println("Failed to get binary path, please provide it using the --path flag.")
				os.Exit(1)
			}
			fmt.Println("Binary path:", binaryPath)

			// Check if the binary path is a symlink
			realPath, err := filepath.EvalSymlinks(binaryPath)
			if err != nil {
				fmt.Println("Failed to evaluate symlink")
				os.Exit(1)
			}
			fmt.Println("Real path:", realPath)
			if realPath != binaryPath {
				fmt.Println("Binary path is a symlink, please provide the actual binary path using the --path flag.")
				os.Exit(1)
			}
		},
	}
	setFlags(updateCmd, l)
	return updateCmd
}

func setFlags(cmd *cobra.Command, l *localizer.Localizer) {
	cmd.Flags().String("path", "", l.Translate("Path to the binary to update. Defaults to the current binary path."))
}