package report

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/language"

	"github.com/jansorg/tom/go-tom/model"
	"github.com/jansorg/tom/go-tom/test_setup"
)

func Test_Report(t *testing.T) {
	ctx, err := test_setup.CreateTestContext(language.German)
	require.NoError(t, err)
	defer test_setup.CleanupTestContext(ctx)

	start := newDate(2018, time.March, 10, 10, 0)
	end := newDate(2018, time.March, 10, 12, 0)

	frameList := []*model.Frame{{Start: start, End: end}}
	report := NewBucketReport(model.NewFrameList(frameList), ctx)
	report.Calculate()
	assert.EqualValues(t, 1, report.Result.FrameCount)
	assert.EqualValues(t, 2*time.Hour, report.Result.Duration.Get())
	assert.EqualValues(t, 2*time.Hour, report.Result.Duration.GetExact())
	assert.EqualValues(t, start, report.Result.TrackedDateRange.Start)
	assert.EqualValues(t, end, report.Result.TrackedDateRange.End)
	assert.EqualValues(t, frameList, report.source.Frames())
}

func Test_ReportSplitYear(t *testing.T) {
	ctx, err := test_setup.CreateTestContext(language.German)
	require.NoError(t, err)
	defer test_setup.CleanupTestContext(ctx)

	// two hours
	start := newDate(2018, time.March, 10, 10, 0)
	end := newDate(2018, time.March, 10, 12, 0)

	// one hour
	start2 := newDate(2019, time.March, 10, 9, 0)
	end2 := newDate(2019, time.March, 10, 10, 0)

	frameList := []*model.Frame{
		{Start: start2, End: end2},
		{Start: start, End: end},
	}

	report := NewBucketReport(model.NewSortedFrameList(frameList), ctx)
	report.SplitOperations = []SplitOperation{SplitByYear}
	report.Calculate()

	require.NotNil(t, report.Result, "expected one top-level group (containing two years)")
	assert.EqualValues(t, 2, report.Result.FrameCount)
	assert.EqualValues(t, 3*time.Hour, report.Result.Duration.Get())
	assert.EqualValues(t, 3*time.Hour, report.Result.Duration.GetExact())
	assert.EqualValues(t, newDate(2018, time.January, 1, 0, 0), report.Result.DateRange.Start)
	assert.EqualValues(t, newDate(2020, time.January, 1, 0, 0), report.Result.DateRange.End)
	assert.EqualValues(t, frameList, report.source.Frames())

	require.EqualValues(t, 2, len(report.Result.Results), "expected a sub-report for each year")

	firstYear := report.Result.Results[0]
	assert.EqualValues(t, newDate(2018, time.January, 1, 0, 0), firstYear.DateRange.Start)
	assert.EqualValues(t, newDate(2019, time.January, 1, 0, 0), firstYear.DateRange.End)
	assert.EqualValues(t, 1, firstYear.FrameCount)
	assert.EqualValues(t, 2*time.Hour, firstYear.Duration.Get())
	assert.EqualValues(t, 2*time.Hour, firstYear.Duration.GetExact())

	secondYear := report.Result.Results[1]
	assert.EqualValues(t, newDate(2019, time.January, 1, 0, 0), secondYear.DateRange.Start)
	assert.EqualValues(t, newDate(2020, time.January, 1, 0, 0), secondYear.DateRange.End)
	assert.EqualValues(t, 1, secondYear.FrameCount)
	assert.EqualValues(t, 1*time.Hour, secondYear.Duration.Get())
	assert.EqualValues(t, 1*time.Hour, secondYear.Duration.GetExact())
}

// tests that frame dates in different time zones are not ending up in different split intervals
func Test_ReportSplitDifferentZones(t *testing.T) {
	ctx, err := test_setup.CreateTestContext(language.German)
	require.NoError(t, err)
	defer test_setup.CleanupTestContext(ctx)

	// two hours, UTC+0
	start := newDate(2018, time.March, 10, 10, 0).UTC()
	end := newDate(2018, time.March, 10, 12, 0).UTC()

	// one hour, UTC+2
	utc2, err := time.LoadLocation("Europe/Berlin")
	require.NoError(t, err)
	start2 := newDate(2018, time.March, 10, 9, 0).In(utc2)
	end2 := newDate(2018, time.March, 10, 10, 0).In(utc2)

	frameList := []*model.Frame{
		{Start: &start, End: &end},
		{Start: &start2, End: &end2},
	}

	report := NewBucketReport(model.NewSortedFrameList(frameList), ctx)
	report.SplitOperations = []SplitOperation{SplitByDay}
	report.Calculate()

	require.NotNil(t, report.Result, "expected one top-level group (containing two days)")
	assert.EqualValues(t, 2, report.Result.FrameCount)

	assert.EqualValues(t, 1, len(report.Result.Results), "expected a single day bucket, even if different time zones were used")
}

// tests that frame dates in different time zones are not ending up in different split intervals
func Test_ReportSplitDifferentZonesYear(t *testing.T) {
	ctx, err := test_setup.CreateTestContext(language.German)
	require.NoError(t, err)
	defer test_setup.CleanupTestContext(ctx)

	// two hours, UTC+0
	start := newDate(2018, time.January, 10, 10, 0).UTC()
	end := newDate(2018, time.January, 10, 12, 0).UTC()

	// one hour, UTC+2
	utc2, err := time.LoadLocation("Europe/Berlin")
	require.NoError(t, err)
	start2 := newDate(2018, time.January, 10, 9, 0).In(utc2)
	end2 := newDate(2018, time.January, 10, 10, 0).In(utc2)

	frameList := []*model.Frame{
		{Start: &start, End: &end},
		{Start: &start2, End: &end2},
	}

	report := NewBucketReport(model.NewSortedFrameList(frameList), ctx)
	report.SplitOperations = []SplitOperation{SplitByYear}
	report.Calculate()

	require.NotNil(t, report.Result, "expected one top-level group (containing two days)")
	assert.EqualValues(t, 2, report.Result.FrameCount)

	assert.EqualValues(t, 1, len(report.Result.Results), "expected a single day bucket, even if different time zones were used")
}

func newDate(year int, month time.Month, day, hour, minute int) *time.Time {
	date := time.Date(year, month, day, hour, minute, 0, 0, time.Local)
	return &date
}
