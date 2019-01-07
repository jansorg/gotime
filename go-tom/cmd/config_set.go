package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jansorg/gotime/go-tom/config"
	"github.com/jansorg/gotime/go-tom/context"
)

func newConfigSetCommand(ctx *context.GoTimeContext, parent *cobra.Command) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "set",
		Short: "sets a configuration value",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			createEmptyConfig(viper.GetString(config.KeyDataDir))

			viper.Set(args[0], args[1])
			if err := viper.WriteConfig(); err != nil {
				fatal("error updating configuration file: ", err)
			}

			fmt.Println("Successfully updated the configuration value of " + args[0])
		},
	}

	parent.AddCommand(cmd)
	return cmd
}
