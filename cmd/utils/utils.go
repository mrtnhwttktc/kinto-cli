package utils

import (
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/mrtnhwttktc/kinto-cli/internal/localizer"
	"github.com/spf13/cobra"
)

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
	c.SetHelpCommand(helpCommand)
}

// LocalizeUsageTemplate localizes the default usage template that is displayed when using the --help flag or help command
func LocalizeUsageTemplate(c *cobra.Command, l *localizer.Localizer) {
	tpl := c.UsageTemplate()
	tpl = strings.NewReplacer(
		"Usage:", l.Translate("Usage:"),
		"Aliases:", l.Translate("Aliases:"),
		"Examples:", l.Translate("Examples:"),
		"Available Commands:", l.Translate("Available Commands:"),
		"Additional Commands:", l.Translate("Additional Commands:"),
		"Flags:", l.Translate("Flags:"),
		"Global Flags:", l.Translate("Global Flags:"),
		"Additional help topics:", l.Translate("Additional help topics:"),
		"Use ", l.Translate("Use"),
		" for more information about a command.", l.Translate("for more information about a command."),
	).Replace(tpl)
	c.SetUsageTemplate(tpl)
}

func LocalizeVersionTemplate(c *cobra.Command, l *localizer.Localizer) {
	tpl := c.VersionTemplate()
	tpl = strings.NewReplacer(
		"version", l.Translate("version"),
	).Replace(tpl)
	c.SetVersionTemplate(tpl)
}

func CustomPromptuiStyle() *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\u23F5 {{ . | blue | bold }}",
		Inactive: "  {{ . | white | faint }}",
		Selected: "{{ \"\U00002714\" | green }} {{ . | green }}",
		Details:  "",
	}
}
