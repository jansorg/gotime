package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jansorg/tom/go-tom/context"
	"github.com/jansorg/tom/go-tom/dataImport/watson"
)

func newImportWatsonCommand(ctx *context.GoTimeContext, parent *cobra.Command) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "watson",
		Short: "",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if result, err := watson.NewImporter().Import(args[0], ctx); err != nil {
				Fatal(err)
			} else {
				fmt.Println(result.String())
			}
		},
	}

	parent.AddCommand(cmd)
	return cmd
}
