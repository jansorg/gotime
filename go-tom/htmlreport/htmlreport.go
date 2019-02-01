package htmlreport

import (
	"bytes"
	"fmt"
	htmlTemplate "html/template"
	"io/ioutil"
	"path"
	"path/filepath"
	"time"

	assetTemplate "github.com/arschles/go-bindata-html-template"

	"github.com/jansorg/tom/go-tom"
	"github.com/jansorg/tom/go-tom/context"
	"github.com/jansorg/tom/go-tom/report"
	"github.com/jansorg/tom/go-tom/util"
)

type Report struct {
	workingDir string
	options    Options
	ctx        *context.TomContext
}

type Options struct {
	CustomTitle        *string          `json:"title"`
	CustomDescription  *string          `json:"description"`
	ShowSummary        bool             `json:"show_summary"`
	ShowExactDurations bool             `json:"show_exact"`
	DecimalDuration    bool             `json:"decimal_duration"`
	TemplateName       *string          `json:"template_name"`
	TemplateFilePath   *string          `json:"template_path"`
	CustomCSS          htmlTemplate.CSS `json:"css"`
	CustomCSSFile      string           `json:"css_file"`
	Report             report.Config    `json:"report"`
}

func NewReport(workingDir string, opts Options, ctx *context.TomContext) *Report {
	return &Report{
		options:    opts,
		workingDir: workingDir,
		ctx:        ctx,
	}
}

func (r *Report) Render(results *report.BucketReport) (string, error) {
	functionMap := map[string]interface{}{
		"reportOptions": func() *Options {
			return &r.options
		},
		"i18n": func(key string) string {
			return r.ctx.LocalePrinter.Sprintf(key)
		},
		"inlineCSS": func(filename string) htmlTemplate.CSS {
			file, err := ioutil.ReadFile(filename)
			if err != nil {
				util.Fatal(fmt.Errorf("error reading CSS file %s", filename))
				return ""
			}
			return htmlTemplate.CSS(file)
		},
		"langBase": func() string {
			base, _ := r.ctx.Language.Base()
			return base.String()
		},
		"formatNumber": func(n interface{}) string {
			return r.ctx.LocalePrinter.Sprint(n)
		},
		"formatTime": func(date time.Time) string {
			return r.ctx.DateTimePrinter.Time(date)
		},
		"formatDate": func(date time.Time) string {
			return r.ctx.DateTimePrinter.Date(date)
		},
		"formatDateTime": func(date time.Time) string {
			return r.ctx.DateTimePrinter.DateTime(date)
		},
		"minDuration": func(duration time.Duration) string {
			if r.options.DecimalDuration {
				return r.ctx.DecimalDurationPrinter.Minimal(duration)
			}
			return r.ctx.DurationPrinter.Minimal(duration)
		},
		"shortDuration": func(duration time.Duration) string {
			if r.options.DecimalDuration {
				return r.ctx.DecimalDurationPrinter.Short(duration)
			}
			return r.ctx.DurationPrinter.Short(duration)
		},
		"longDuration": func(duration time.Duration) string {
			if r.options.DecimalDuration {
				return r.ctx.DecimalDurationPrinter.Long(duration)
			}
			return r.ctx.DurationPrinter.Long(duration)
		},
		"isMatrix": func(bucket report.ResultBucket) bool {
			if bucket.Depth() != 2 {
				return false
			}

			// all buckets must have the same number of children
			refCol := bucket.ChildBuckets[0].ChildBuckets
			for _, b := range bucket.ChildBuckets {
				if len(b.ChildBuckets) != len(refCol) {
					return false
				}

				for i, col := range b.ChildBuckets {
					other := refCol[i]

					// if both are the same month names, then accepot
					if col.SplitByType == report.SplitByMonth && other.SplitByType == report.SplitByMonth && col.DateRange().IsMonthRange() && other.DateRange().IsMonthRange() && col.DateRange().Start.Month() == other.DateRange().Start.Month() {
						continue
					}

					if col.Title() != other.Title() {
						return false
					}
				}
			}

			return true
		},
		"sumChildValues": func(parent report.ResultBucket, childIndex int) *util.DurationSum {
			sum := util.NewDurationSum()

			for _, b := range parent.ChildBuckets {
				if len(b.ChildBuckets) >= childIndex {
					sum.AddSum(b.ChildBuckets[childIndex].Duration)
				}
			}

			return sum
		},
	}

	if r.options.TemplateFilePath != nil && *r.options.TemplateFilePath != "" {
		templatePath := *r.options.TemplateFilePath
		if !filepath.IsAbs(templatePath) {
			templatePath = filepath.Join(r.workingDir, templatePath)
		}

		files, err := filepath.Glob(filepath.Join(filepath.Dir(templatePath), "*.gohtml"))

		tmpl, err := htmlTemplate.New(filepath.Base(templatePath)).Funcs(functionMap).ParseFiles(append(files, templatePath)...)
		if err != nil {
			return "", err
		}
		out := bytes.NewBuffer([]byte{})
		if err = tmpl.Execute(out, results); err != nil {
			return "", err
		}

		return out.String(), nil
	} else if r.options.TemplateName != nil && *r.options.TemplateName != "" {
		templatePath := *r.options.TemplateName

		baseDir := path.Join("reports", "html")
		templateFiles := []string{
			path.Join(baseDir, templatePath+".gohtml"),
			path.Join(baseDir, "commons.gohtml"),
		}

		tmpl, err := assetTemplate.New(templatePath, tom.Asset).Funcs(functionMap).ParseFiles(templateFiles...)
		if err != nil {
			return "", err
		}
		out := bytes.NewBuffer([]byte{})
		if err = tmpl.Execute(out, results); err != nil {
			return "", err
		}

		return out.String(), nil
	} else {
		return "", fmt.Errorf("template undefined")
	}
}
