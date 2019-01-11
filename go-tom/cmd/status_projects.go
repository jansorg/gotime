package cmd

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/jansorg/tom/go-tom/context"
	"github.com/jansorg/tom/go-tom/report"
)

type projectStatusList struct {
	reports []*report.ProjectSummary
}

func (o projectStatusList) size() int {
	return len(o.reports)
}

func (o projectStatusList) get(index int, prop string, format string) (interface{}, error) {
	summary := o.reports[index]

	switch prop {
	case "id":
		return summary.Project.ID, nil
	case "parentID":
		return summary.Project.ParentID, nil
	case "fullName":
		return summary.Project.FullName, nil
	case "name":
		return summary.Project.Name, nil
	case "trackedDay":
		return summary.TrackedDay, nil
	case "trackedWeek":
		return summary.TrackedWeek, nil
	case "trackedMonth":
		return summary.TrackedMonth, nil
	case "trackedYear":
		return summary.TrackedYear, nil
	case "totalTrackedDay":
		return summary.TotalTrackedDay, nil
	case "totalTrackedWeek":
		duration := summary.TotalTrackedWeek
		return duration, nil
	case "totalTrackedMonth":
		duration := summary.TotalTrackedMonth
		return duration, nil
	case "totalTrackedYear":
		return summary.TotalTrackedYear, nil
	default:
		return "", fmt.Errorf("unknown property %s", prop)
	}
}

func newProjectsStatusCommand(ctx *context.GoTimeContext, parent *cobra.Command) *cobra.Command {
	showEmpty := false

	var cmd = &cobra.Command{
		Use:   "projects",
		Short: "Prints project status",
		Run: func(cmd *cobra.Command, args []string) {
			projectReports := report.CreateProjectReports(time.Now(), showEmpty, ctx)

			var reportList []*report.ProjectSummary
			for _, v := range projectReports {
				reportList = append(reportList, v)
			}
			sort.Slice(reportList, func(i, j int) bool {
				return strings.Compare(reportList[i].Project.FullName, reportList[j].Project.FullName) < 0
			})

			if err := printList(cmd, projectStatusList{reports: reportList}, ctx); err != nil {
				fatal(err)
			}
		},
	}

	cmd.Flags().BoolVarP(&showEmpty, "show-empty", "e", showEmpty, "Includes projects without tracked time in the output")

	addListOutputFlags(cmd, "fullName,trackedDay,trackedWeek,trackedMonth", []string{"id", "fullName", "name", "parentID", "trackedDay", "trackedWeek", "trackedMonth", "trackedYear", "totalTrackedDay", "totalTrackedWeek", "totalTrackedMonth", "totalTrackedYear"})
	parent.AddCommand(cmd)
	return cmd
}