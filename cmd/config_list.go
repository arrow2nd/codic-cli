package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (c *Cmd) newConfigListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Display the current setting parameters",
		Args:  cobra.NoArgs,
		Run:   c.execConfigListCmd,
	}
}

func (c *Cmd) execConfigListCmd(cmd *cobra.Command, args []string) {
	showValue("Case Type", "case")
	showValue("Prefix Style", "prefix")
}

func showValue(title, valueName string) {
	fmt.Printf("%s %s\n", color.HiBlackString("%s:", title), viper.GetString(valueName))
}
