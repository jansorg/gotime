package dateUtil

import (
	"fmt"
	"time"
)

func NewDateRange(start *time.Time, end *time.Time) DateRange {
	dateRange := DateRange{Start: start, End: end}
	dateRange.debug = dateRange.String()
	return dateRange
}

func NewYearRange(date time.Time) DateRange {
	start := time.Date(date.Year(), time.January, 1, 0, 0, 0, 0, date.Location())
	end := start.AddDate(1, 0, 0)

	return NewDateRange(&start, &end)
}

func NewMonthRange(date time.Time) DateRange {
	y, m, _ := date.Date()
	start := time.Date(y, m, 1, 0, 0, 0, 0, date.Location())
	end := start.AddDate(0, 1, 0)

	return NewDateRange(&start, &end)
}

func NewDayRange(date time.Time) DateRange {
	y, m, d := date.Date()
	start := time.Date(y, m, d, 0, 0, 0, 0, date.Location())
	end := start.AddDate(0, 0, 1)

	return NewDateRange(&start, &end)
}

type DateRange struct {
	Start *time.Time `json:"start"`
	End   *time.Time `json:"end"`
	debug string
}

func (r DateRange) String() string {
	var start, end string
	if r.Start != nil {
		start = r.Start.String()
	}
	if r.End != nil {
		end = r.End.String()
	}
	return fmt.Sprintf("%s - %s", start, end)
}

func (r DateRange) ShortString() string {
	var start, end string
	if r.Start != nil {
		start = ShortDateString(*r.Start)
	}
	if r.End != nil {
		end = ShortDateString(*r.End)
	}
	return fmt.Sprintf("%s - %s", start, end)
}

func (r DateRange) MinimalString() string {
	var y1, d1, y2, d2 int
	var m1, m2 time.Month

	var start, end string
	if r.Start != nil {
		start = ShortDateString(*r.Start)
		y1, m1, d1 = r.Start.Date()
	}
	if r.End != nil {
		end = ShortDateString(*r.End)
		y2, m2, d2 = r.End.Date()
	}

	if y1 == y2 {
		if m1 == m2 {
			if d1 == d2 {
				// print just the year
				return fmt.Sprintf("%04d", y1)
			}
		}
	}

	return fmt.Sprintf("%s - %s", start, end)
}

func (r DateRange) Shift(years, months, days int) DateRange {
	if r.Start != nil && !r.Start.IsZero() {
		*r.Start = r.Start.AddDate(years, months, days)
	}
	if r.End != nil && !r.End.IsZero() {
		*r.End = r.End.AddDate(years, months, days)
	}
	r.debug = r.String()
	return r
}

func (r DateRange) IsClosed() bool {
	return r.Start != nil && r.End != nil
}

func (r DateRange) IsOpen() bool {
	return !r.IsClosed()
}

func (r DateRange) Empty() bool {
	return r.Start == nil && r.End == nil
}
