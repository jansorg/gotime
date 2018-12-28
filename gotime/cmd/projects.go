package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jansorg/gotime/gotime/context"
)

func newProjectsCommand(context *context.GoTimeContext, parent *cobra.Command) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "projects",
		Short: "Print a listing of all projects",
		Run: func(cmd *cobra.Command, args []string) {
			projects := context.Store.Projects()

			if context.JsonOutput {
				if bytes, err := json.MarshalIndent(projects, "", "  "); err != nil {
					fatal(err)
				} else {
					fmt.Println(string(bytes))
				}
			} else {
				for _, p := range projects {
					fmt.Println(p.ID)
				}
			}
		},
	}

	parent.AddCommand(cmd)
	return cmd
}
