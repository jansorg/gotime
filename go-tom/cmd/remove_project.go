package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jansorg/tom/go-tom/cmd/util"
	"github.com/jansorg/tom/go-tom/context"
)

func newRemoveProjectCommand(context *context.GoTimeContext, parent *cobra.Command) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "project <project name or project ID> ...",
		Short: "removes new project and all its associated data (including subprojects)",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			removedFrameCount := 0
			removedProjects := 0

			for _, name := range args {
				projects := context.Query.ProjectsByShortNameOrID(name)

				for _, p := range projects {
					frames := context.Query.FramesByProject(p.ID, true)
					for _, f := range frames {
						if err := context.Store.RemoveFrame(f.ID); err != nil {
							util.Fatal(err)
						}
						removedFrameCount++
					}

					if err := context.Store.RemoveProject(p.ID); err != nil {
						util.Fatal(err)
					}
					removedProjects++
				}
			}

			fmt.Printf("Removed projects: %d\nRemoved frames: %d\n", removedProjects, removedFrameCount)
		},
	}
	parent.AddCommand(cmd)

	return cmd
}
