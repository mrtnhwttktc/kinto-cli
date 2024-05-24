package set

import (
	"fmt"

	"github.com/mrtnhwttktc/kinto-cli/cmd/cli/set/language"
	"github.com/mrtnhwttktc/kinto-cli/cmd/utils"
	"github.com/mrtnhwttktc/kinto-cli/internal/localizer"
	"github.com/spf13/cobra"
)

func NewSetCmd() *cobra.Command {
	l := localizer.GetLocalizer()
	setCmd := &cobra.Command{
		Use:   "set",
		Short: l.Translate("Set global configurations."),
		Long:  l.Translate("Set global configurations for ktcli. Updates the local configuration file with the selected option."),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("set called")
		},
	}
	bindFlags(setCmd, l)

	setCmd.AddCommand(language.NewLanguageCmd())

	return setCmd
}

func bindFlags(cmd *cobra.Command, l *localizer.Localizer) {
	utils.LocalizeHelpFlag(cmd, l)
}
