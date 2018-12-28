package cmd

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/jansorg/gotime/gotime/context"
	"github.com/jansorg/gotime/gotime/dateUtil"
	"github.com/jansorg/gotime/gotime/report"
)

func newReportCommand(context *context.GoTimeContext, parent *cobra.Command) *cobra.Command {
	var fromDateString string
	var toDateString string

	var day int8
	var month int8
	var year int

	var splitModes []string

	var roundFrames time.Duration
	var roundTotals time.Duration

	var roundModeFrames string
	var roundModeTotal string

	var cmd = &cobra.Command{
		Use:   "report",
		Short: "Generate reports about your tracked time",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			var start *time.Time
			var end *time.Time

			if fromDateString != "" {
				start, err = parseDate(&fromDateString)
				if err != nil {
					fatal(err)
				}
			}

			if toDateString != "" {
				end, err = parseDate(&toDateString)
				if err != nil {
					fatal(err)
				}
			}

			if cmd.Flag("day") != nil {
				now := time.Now()
				year, month, today := now.Date()
				start = createDate(year, month, today, day)
				end = createDate(year, month, today+int(day), 24)
			} else if cmd.Flag("month") != nil {
				now := time.Now()
				year, currentMonth, _ := now.Date()
				start = createDate(year, time.Month(int(currentMonth)+int(month)), 0, 0)
				end = createDate(year, time.Month(int(currentMonth)+int(month)+1), 0, 0)
			} else if cmd.Flag("year") != nil {
				now := time.Now()
				currentYear, _, _ := now.Date()
				start = createDate(currentYear+year, time.January, 0, 0)
				end = createDate(currentYear+year, time.December, 24, 0)
			}

			var frameRoundingMode = dateUtil.ParseRoundingMode(roundModeFrames)
			var totalsRoundingNode = dateUtil.ParseRoundingMode(roundModeTotal)

			var splitOperations []report.SplitOperation
			for _, mode := range splitModes {
				switch mode {
				case "year":
					splitOperations = append(splitOperations, report.SplitByYear)
				case "month":
					splitOperations = append(splitOperations, report.SplitByMonth)
				case "day":
					splitOperations = append(splitOperations, report.SplitByDay)
				case "project":
					splitOperations = append(splitOperations, report.SplitByProject)
				default:
					fatal(fmt.Errorf("unknown split value %s. Supported: year, month, day", mode))
				}
			}

			frameReport := report.NewBucketReport(context.Store.Frames())
			frameReport.FromDate = start
			frameReport.ToDate = end
			frameReport.RoundFramesTo = roundFrames
			frameReport.RoundTotalsTo = roundTotals
			frameReport.RoundingFrames = frameRoundingMode
			frameReport.RoundingTotals = totalsRoundingNode
			frameReport.SplitOperations = splitOperations
			frameReport.Update()

			if context.JsonOutput {
				data, err := json.MarshalIndent(frameReport.Result, "", "  ")
				if err != nil {
					fatal(err)
				}
				fmt.Println(string(data))
			} else {
				printReport(frameReport.Result, context, 1)
			}
		},
	}

	cmd.PersistentFlags().StringVarP(&fromDateString, "from", "f", "", "Optional start date")
	cmd.PersistentFlags().StringVarP(&toDateString, "to", "t", "", "Optional end date")
	cmd.PersistentFlags().Int8VarP(&day, "day", "", 0, "Select the date range of a given day. For example, 0 is today, -1 is one day ago, etc.")
	cmd.PersistentFlags().Int8VarP(&month, "month", "", 0, "Filter on a given month. For example, 0 is the current month, -1 is last month, etc.")
	cmd.PersistentFlags().IntVarP(&year, "year", "", 0, "Filter on a specific year. 0 is the current year, -1 is last year, etc.")

	cmd.PersistentFlags().StringArrayVarP(&splitModes, "group", "", []string{}, "Group frames into years, months and/or days. Possible values: year,month,day")

	cmd.PersistentFlags().DurationVarP(&roundFrames, "round-frames-to", "", time.Duration(0), "Round durations of each frame to the nearest multiple of this duration")
	cmd.PersistentFlags().StringVarP(&roundModeFrames, "round-frames", "", "up", "Rounding mode for sums of durations. Default: up. Possible values: up|nearest")

	cmd.PersistentFlags().DurationVarP(&roundTotals, "round-totals-to", "", time.Duration(0), "Round the overall duration of each project to the next matching multiple of this duration")
	cmd.PersistentFlags().StringVarP(&roundModeTotal, "round-totals", "", "up", "Rounding mode for sums of durations. Default: up. Possible values: up|nearest")

	parent.AddCommand(cmd)
	newReportHtmlCommand(context, cmd)

	return cmd
}

func createDate(year int, month time.Month, today int, day int8) *time.Time {
	date := time.Date(year, month, today+int(day), 0, 0, 0, 0, time.Local)
	return &date
}

func parseDate(dateString *string) (*time.Time, error) {
	result, err := time.Parse(time.RFC3339Nano, *dateString)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//noinspection ALL
func printReport(report *report.ResultBucket, ctx *context.GoTimeContext, level int) {
	title := report.Title(ctx)
	if title != "" {
		printlnIndenting(level-1, title)
	} else if (level == 1) {
		printlnIndenting(level-1, "Overall")
	}

	if report.Start != nil && report.End != nil {
		printfIndenting(level, "%s - %s\n", report.Start.String(), report.End.String())
	} else if report.Start != nil {
		printlnIndenting(level, report.Start.String())
	} else if (report.End != nil) {
		printfIndenting(level, report.End.String())
	}

	printfIndenting(level, "Duration: %s\n", report.Duration.String())
	printfIndenting(level, "Exact Duration: %s\n", report.ExactDuration.String())
	fmt.Println()

	for _, r := range report.Results {
		printReport(r, ctx, level+1)
	}
}

func printlnIndenting(level int, value string) (int, error) {
	fmt.Print(strings.Repeat("    ", level))
	return fmt.Println(value)
}

func printfIndenting(level int, format string, a ...interface{}) (int, error) {
	fmt.Print(strings.Repeat("    ", level))
	return fmt.Printf(format, a...)
}
