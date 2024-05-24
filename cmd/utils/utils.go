package utils

import (
	"github.com/mrtnhwttktc/kinto-cli/internal/localizer"
	"github.com/spf13/cobra"
)

// LocalizeHelpFlag localizes the help flag for a command
// allows to use one translation for all help flags
func LocalizeHelpFlag(cmd *cobra.Command, l *localizer.Localizer) {
	cmd.Flags().BoolP("help", "h", false, l.Translate("help for %s", cmd.Name()))
}

// LocalizeHelpFunc localizes the default help function
func LocalizeHelpFunc(c *cobra.Command, l *localizer.Localizer) {
	helpCommand := &cobra.Command{
		Use:   "help [command]",
		Short: l.Translate("Help about any command"),
		Long: l.Translate(`Help provides help for any command in the application.
Simply type %s help [path to command] for full details.`, c.Name()),
		Run: func(c *cobra.Command, args []string) {
			cmd, _, e := c.Root().Find(args)
			if cmd == nil || e != nil {
				c.Println(l.Translate("Unknown help topic %#q", args))
				cobra.CheckErr(c.Root().Usage())
			} else {
				cobra.CheckErr(cmd.Help())
			}
		},
	}
	LocalizeHelpFlag(helpCommand, l)
	c.SetHelpCommand(helpCommand)

}
