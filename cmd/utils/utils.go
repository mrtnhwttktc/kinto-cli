package utils

import (
	"strings"

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

// LocalizeUsageTemplate localizes the default usage template that is displayed when using the --help flag or help command
func LocalizeUsageTemplate(c *cobra.Command, l *localizer.Localizer) {

	usage := l.Translate("Usage:")
	aliases := l.Translate("Aliases:")
	examples := l.Translate("Examples:")
	availableCommands := l.Translate("Available Commands:")
	additionalCommands := l.Translate("Additional Commands:")
	flags := l.Translate("Flags:")
	globalFlags := l.Translate("Global Flags:")
	additionalHelpTopics := l.Translate("Additional help topics:")
	helpUsage1 := l.Translate("Use")
	helpUsage2 := l.Translate("for more information about a command.")
	tpl := c.UsageTemplate()
	tpl = strings.NewReplacer(
		"Usage:", usage,
		"Aliases:", aliases,
		"Examples:", examples,
		"Available Commands:", availableCommands,
		"Additional Commands:", additionalCommands,
		"Flags:", flags,
		"Global Flags:", globalFlags,
		"Additional help topics:", additionalHelpTopics,
		"Use ", helpUsage1,
		" for more information about a command.", helpUsage2,
	).Replace(tpl)
	c.SetUsageTemplate(tpl)
}
