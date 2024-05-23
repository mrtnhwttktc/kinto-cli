package set

import (
	"fmt"

	"github.com/mrtnhwttktc/kinto-cli/cmd/set/language"
	"github.com/spf13/cobra"
)

func NewSetCmd() *cobra.Command {
	setCmd := &cobra.Command{
		Use:   "set",
		Short: "A brief description of your command",
		Long:  "Description of your command.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("set called")
		},
	}

	setCmd.AddCommand(language.NewLanguageCmd())

	return setCmd
}
