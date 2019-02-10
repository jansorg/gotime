package report

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/jansorg/tom/go-tom/context"
	"github.com/jansorg/tom/go-tom/htmlreport"
	"github.com/jansorg/tom/go-tom/model"
	"github.com/jansorg/tom/go-tom/report"
	"github.com/jansorg/tom/go-tom/util"
)

type flags struct {
	title             string
	description       string
	showEmpty         bool
	showMatrixTables  bool
	fromDateString    string
	toDateString      string
	projectFilter     []string
	includeSubproject bool
	day               int
	month             int
	year              int
	splitModes        string
	roundFrames       time.Duration
	roundTotals       time.Duration
	roundModeFrames   string
	roundModeTotal    string
	decimalDurations  bool
	showSales         bool
	showSummary       bool
	templateName      string
	templateFilePath  string
	archivedFrames    bool
	properties        []string
}

var defaultFlags = flags{
	templateName:     "default",
	showMatrixTables: true,
	archivedFrames:   true,
	showSales:        false,
}

func NewCommand(ctx *context.TomContext, parent *cobra.Command) *cobra.Command {
	opts := defaultFlags

	var configFile string
	var saveConfigFile string
	var jsonOutput bool
	var htmlOutputFile string

	var cmd = &cobra.Command{
		Use:   "report",
		Short: "Generate reports about your tracked time",
		Run: func(cmd *cobra.Command, args []string) {
			config := htmlreport.DefaultOptions
			var err error

			if configFile != "" {
				config, err = loadJsonConfig(ctx, configFile)
				if err != nil {
					util.Fatal(err)
				}

				if flagConfig, err := configByFlags(opts, cmd, ctx); err == nil {
					applyFlags(cmd, flagConfig, &config)
				}
			} else {
				config, err = configByFlags(opts, cmd, ctx)
				if err != nil {
					util.Fatal(err)
				}
			}

			if saveConfigFile != "" {
				data, err := json.MarshalIndent(config, "", "  ")
				if err != nil {
					util.Fatal(err)
				}
				if err = ioutil.WriteFile(saveConfigFile, data, 0600); err != nil {
					util.Fatal(err)
				}
			}

			frameReport := report.NewBucketReport(model.NewSortedFrameList(ctx.Store.Frames()), config.Report, ctx)
			result := frameReport.Update()

			var data []byte
			if jsonOutput {
				data, err = json.MarshalIndent(result, "", "  ")
				if err != nil {
					log.Fatal(err)
				}
			} else {
				if data, err = renderReport(ctx, frameReport, config); err != nil {
					util.Fatal(fmt.Errorf("error while rendering: %s", err.Error()))
				}
			}

			if htmlOutputFile != "" {
				err = ioutil.WriteFile(htmlOutputFile, data, 0600)
				if err != nil {
					util.Fatal(err)
				}
			} else {
				fmt.Println(string(data))
			}
		},
	}

	cmd.Flags().BoolVarP(&jsonOutput, "json", "", false, "Prints JSON instead of plain text")
	cmd.Flags().StringVarP(&configFile, "config", "", "", "Path to a json configuration")
	cmd.Flags().StringVarP(&saveConfigFile, "save-config", "", "", "Path where the options are saved as a template")
	cmd.Flags().StringVarP(&htmlOutputFile, "output-file", "o", "", "Path where the rendered data will be written")

	cmd.Flags().StringVarP(&opts.templateName, "template", "", opts.templateName, "Template to use for rendering. This may either be a full path to a template file or the name (without extension) of a template shipped with gotime.")

	templateAnnotations := make(map[string][]string)
	templateAnnotations[cobra.BashCompFilenameExt] = []string{"gohtml"}
	cmd.Flags().StringVarP(&opts.templateFilePath, "template-file", "", "", "Custom gohtml template file to use for rendering. See the website for more details.")
	cmd.Flag("template-file").Annotations = templateAnnotations

	// fixme add defaults?
	// cmd.Flags().BoolVarP(&includeActiveFrames, "current", "c", false, "(Don't) Include currently running frame in report.")
	cmd.Flags().StringVarP(&opts.fromDateString, "from", "f", "", "The date when the report should start.")
	cmd.Flags().StringVarP(&opts.toDateString, "to", "t", "", "Optional end date")

	cmd.Flags().IntVarP(&opts.year, "year", "y", 0, "Filter on a specific year. 0 is the current year, -1 is last year, etc.")
	cmd.Flags().IntVarP(&opts.month, "month", "m", 0, "Filter on a given month. For example, 0 is the current month, -1 is last month, etc.")
	cmd.Flag("month").NoOptDefVal = "0"
	cmd.Flags().IntVarP(&opts.day, "day", "d", 0, "Select the date range of a given day. For example, 0 is today, -1 is one day ago, etc.")

	cmd.Flags().StringSliceVarP(&opts.projectFilter, "project", "p", []string{}, "--project ID | NAME . Reports activities only for the given project. You can add other projects by using this option multiple times.")
	cmd.Flags().BoolVarP(&opts.includeSubproject, "subprojects", "", true, "Automatically add the subprojects of the selected projects.")

	cmd.Flags().StringVarP(&opts.splitModes, "split", "s", "project", "Split the report into groups. Multiple values are possible. Possible values: year,month,week,day,project")

	cmd.Flags().DurationVarP(&opts.roundFrames, "round-frames-to", "", time.Minute, "Round durations of each frame to the nearest multiple of this duration")
	cmd.Flags().StringVarP(&opts.roundModeFrames, "round-frames", "", "", "Rounding mode for sums of durations. Default: no rounding. Possible values: up|nearest")

	cmd.Flags().DurationVarP(&opts.roundTotals, "round-totals-to", "", time.Minute, "Round durations of each frame to the nearest multiple of this duration")
	cmd.Flags().StringVarP(&opts.roundModeTotal, "round-totals", "", "", "Rounding mode for sums of durations. Default: no rounding. Possible values: up|nearest.")

	cmd.Flags().BoolVarP(&opts.showEmpty, "show-empty", "", false, "Show empty groups")
	cmd.Flags().BoolVarP(&opts.decimalDurations, "decimal", "", false, "Print durations as decimals 1.5h instead of 1:30h")
	cmd.Flags().BoolVarP(&opts.showSummary, "show-summary", "", defaultFlags.showSummary, "Show a report summary at the top of the report")
	cmd.Flags().BoolVarP(&opts.showMatrixTables, "matrix-tables", "", defaultFlags.showMatrixTables, "Show matrix tables when applicable instead of a list of tables")
	cmd.Flags().StringVarP(&opts.title, "title", "", "", "This will be displayed as the reports title when you're using the default templates")
	cmd.Flags().StringVarP(&opts.description, "description", "", "", "This will be displayed as the reports description when you're using the default templates")
	cmd.Flags().BoolVarP(&opts.archivedFrames, "include-archived", "", defaultFlags.archivedFrames, "Include archived frames in the reported times")
	cmd.Flags().StringSliceVarP(&opts.properties, "property", "", defaultFlags.properties, "Project properties to include in the report")
	cmd.Flags().BoolVarP(&opts.showSales, "show-sales", "", defaultFlags.showSales, "Show sales summaries")

	parent.AddCommand(cmd)
	return cmd
}

func applyFlags(cmd *cobra.Command, source htmlreport.Options, target *htmlreport.Options) {
	if cmd.Flag("template").Changed {
		target.TemplateName = source.TemplateName
	}
	if cmd.Flag("template-file").Changed {
		target.TemplateFilePath = source.TemplateFilePath
	}
	if cmd.Flag("from").Changed || cmd.Flag("to").Changed || cmd.Flag("year").Changed || cmd.Flag("month").Changed || cmd.Flag("day").Changed {
		target.Report.DateFilterRange = source.Report.DateFilterRange
	}
	if cmd.Flag("project").Changed {
		target.Report.ProjectIDs = source.Report.ProjectIDs
	}
	if cmd.Flag("subprojects").Changed {
		target.Report.IncludeSubprojects = source.Report.IncludeSubprojects
	}
	if cmd.Flag("split").Changed {
		target.Report.Splitting = source.Report.Splitting
	}
	if cmd.Flag("round-frames-to").Changed {
		target.Report.EntryRounding.Size = source.Report.EntryRounding.Size
	}
	if cmd.Flag("round-frames").Changed {
		target.Report.EntryRounding.Mode = source.Report.EntryRounding.Mode
	}
	if cmd.Flag("round-totals-to").Changed {
		target.Report.SumRounding.Size = source.Report.SumRounding.Size
	}
	if cmd.Flag("round-totals").Changed {
		target.Report.SumRounding.Mode = source.Report.SumRounding.Mode
	}
	if cmd.Flag("decimal").Changed {
		target.DecimalDuration = source.DecimalDuration
	}
	if cmd.Flag("show-summary").Changed {
		target.ShowSummary = source.ShowSummary
	}
	if cmd.Flag("matrix-tables").Changed {
		target.ShowMatrixTables = source.ShowMatrixTables
	}
	if cmd.Flag("title").Changed {
		target.CustomTitle = source.CustomTitle
	}
	if cmd.Flag("description").Changed {
		target.CustomDescription = source.CustomDescription
	}
	if cmd.Flag("include-archived").Changed {
		target.Report.IncludeArchived = source.Report.IncludeArchived
	}
	if cmd.Flag("show-sales").Changed {
		target.ShowSales = source.ShowSales
	}
}

func loadJsonConfig(ctx *context.TomContext, filePath string) (htmlreport.Options, error) {
	var config htmlreport.Options
	if data, err := ioutil.ReadFile(filePath); err != nil {
		util.Fatal(err)
	} else if err := json.Unmarshal(data, &config); err != nil {
		util.Fatal(err)
	}

	// validate project IDs
	ids := []string{}
	for _, idOrName := range config.Report.ProjectIDs {
		project, err := ctx.Query.ProjectByFullNameOrID(idOrName, "/")
		if err != nil {
			return htmlreport.Options{}, fmt.Errorf("validating project %s: %s", idOrName, err.Error())
		}
		ids = append(ids, project.ID)
	}
	config.Report.ProjectIDs = ids
	return config, nil
}

func configByFlags(opts flags, cmd *cobra.Command, ctx *context.TomContext) (htmlreport.Options, error) {
	filterRange := util.NewDateRange(nil, nil, ctx.Locale)

	if opts.fromDateString != "" {
		start, err := parseDate(&opts.fromDateString)
		if err != nil {
			util.Fatal(err)
		}
		filterRange.Start = start
	}

	if opts.toDateString != "" {
		end, err := parseDate(&opts.toDateString)
		if err != nil {
			log.Fatal(err)
		}
		filterRange.End = end
	}

	// day, month, year params override the filter values
	if cmd.Flag("day").Changed {
		filterRange = util.NewDayRange(time.Now(), ctx.Locale, time.Local).Shift(0, 0, opts.day)
	} else if cmd.Flag("month").Changed {
		filterRange = util.NewMonthRange(time.Now(), ctx.Locale, time.Local).Shift(0, opts.month, 0)
	} else if cmd.Flag("year").Changed {
		filterRange = util.NewYearRange(time.Now(), ctx.Locale, time.Local).Shift(opts.year, 0, 0)
	}

	var splitOperations []report.SplitOperation
	if opts.splitModes != "" {
		for _, mode := range strings.Split(opts.splitModes, ",") {
			if op, err := report.SplitOperationByName(mode); err != nil {
				util.Fatal(err)
			} else {
				splitOperations = append(splitOperations, op)
			}
		}
	}

	// project filter
	var projectIDs []string
	// resolve names or IDs to IDs only
	for _, nameOrID := range opts.projectFilter {
		id := nameOrID
		// if it's a name resolve it to the ID
		if project, err := ctx.Query.ProjectByFullName(strings.Split(nameOrID, "/")); err == nil {
			id = project.ID
		} else if _, err := ctx.Query.ProjectByID(nameOrID); err != nil {
			util.Fatal(fmt.Errorf("project %s not found", opts.projectFilter))
		}
		projectIDs = append(projectIDs, id)
	}

	return htmlreport.Options{
		TemplateFilePath:  &opts.templateFilePath,
		TemplateName:      &opts.templateName,
		DecimalDuration:   opts.decimalDurations,
		ShowMatrixTables:  opts.showMatrixTables,
		ShowSummary:       opts.showSummary,
		CustomTitle:       &opts.title,
		CustomDescription: &opts.description,
		ShowSales:         opts.showSales,
		Report: report.Config{
			// fixme missing timezone
			ProjectIDs:         projectIDs,
			IncludeSubprojects: opts.includeSubproject,
			IncludeArchived:    opts.archivedFrames,
			DateFilterRange:    filterRange,
			Splitting:          splitOperations,
			ShowEmpty:          opts.showEmpty,
			EntryRounding: util.RoundingConfig{
				Mode: util.RoundingByName(opts.roundModeFrames),
				Size: opts.roundFrames,
			},
			SumRounding: util.RoundingConfig{
				Mode: util.RoundingByName(opts.roundModeTotal),
				Size: opts.roundTotals,
			},
		},
	}, nil
}

func renderReport(ctx *context.TomContext, report *report.BucketReport, opts htmlreport.Options) ([]byte, error) {
	dir, _ := os.Getwd()
	t := htmlreport.NewReport(dir, opts, ctx)
	return t.Render(report)
}

func parseDate(dateString *string) (*time.Time, error) {
	result, err := time.Parse(time.RFC3339Nano, *dateString)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
