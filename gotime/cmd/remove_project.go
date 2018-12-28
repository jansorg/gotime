package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jansorg/gotime/gotime/context"
)

func newRemoveProjectCommand(context *context.GoTimeContext, parent *cobra.Command) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "project <project name or project ID> ...",
		Short: "removes new project and all its associated data",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			remvoedFrames := 0
			remvoedProjects := 0

			for _, name := range args {
				projects := context.Query.ProjectsByShortNameOrID(name)

				for _, p := range projects {
					frames := context.Query.FramesByProject(p.ID)
					for _, f := range frames {
						if err := context.Store.RemoveFrame(f.Id); err != nil {
							fatal(err)
						}
						remvoedFrames++
					}

					if err := context.Store.RemoveProject(p.ID); err != nil {
						fatal(err)
					}
					remvoedProjects++
				}
			}

			fmt.Printf("Removed projects: %d\nRemoved frames: %d\n", remvoedProjects, remvoedFrames)
		},
	}
	parent.AddCommand(cmd)

	return cmd
}
