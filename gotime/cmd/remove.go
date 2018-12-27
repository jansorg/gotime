package cmd

import (
	"github.com/spf13/cobra"

	"../context"
)

func newRemoveCommand(ctx *context.GoTimeContext, parent *cobra.Command) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "remove",
		Short: "create projects, tags or frames",
	}

	newRemoveProjectCommand(ctx, cmd)

	parent.AddCommand(cmd)
	return cmd
}
