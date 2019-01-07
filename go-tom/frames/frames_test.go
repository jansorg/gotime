package frames

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jansorg/tom/go-tom/store"
)

func TestSplitByDay(t *testing.T) {
	buckets := NewFrameList([]*store.Frame{
		{
			Start: newDay(10, 12, 0),
			End:   newDay(10, 13, 30),
		},
		{
			Start: newDay(10, 14, 0),
			End:   newDay(10, 16, 0),
		},
		{
			Start: newDay(11, 12, 0),
			End:   newDay(11, 13, 30),
		},
		{
			Start: newDay(13, 1, 0),
			End:   newDay(13, 3, 0),
		},
	}).SplitByDay()

	assert.EqualValues(t, 3, len(buckets))
	assert.EqualValues(t, 2, len(buckets[0].Frames))
	assert.EqualValues(t, 1, len(buckets[1].Frames))
	assert.EqualValues(t, 1, len(buckets[1].Frames))
}

func TestSplitByMonth(t *testing.T) {
	buckets := NewFrameList([]*store.Frame{
		{
			Start: newDate(2017, time.February, 10, 12, 0),
			End:   newDate(2017, time.February, 11, 20, 0),
		},
		{
			Start: newDate(2018, time.February, 10, 12, 0),
			End:   newDate(2018, time.February, 13, 12, 0),
		},
		{
			Start: newDate(2018, time.February, 14, 12, 0),
			End:   newDate(2018, time.February, 16, 12, 0),
		},
	}).SplitByMonth()

	assert.EqualValues(t, 2, len(buckets))
	assert.EqualValues(t, 1, len(buckets[0].Frames))
	assert.EqualValues(t, 2, len(buckets[1].Frames))
}

func TestSplitByYear(t *testing.T) {
	buckets := NewFrameList([]*store.Frame{
		{
			Start: newDate(2017, time.February, 10, 12, 0),
			End:   newDate(2017, time.February, 11, 20, 0),
		},
		{
			Start: newDate(2018, time.February, 10, 12, 0),
			End:   newDate(2018, time.February, 13, 12, 0),
		},
		{
			Start: newDate(2018, time.February, 14, 12, 0),
			End:   newDate(2018, time.February, 16, 12, 0),
		},
		{
			Start: newDate(2018, time.March, 14, 12, 0),
			End:   newDate(2018, time.April, 16, 12, 0),
		},
	}).SplitByYear()

	assert.EqualValues(t, 2, len(buckets))
	assert.EqualValues(t, 1, len(buckets[0].Frames))
	assert.EqualValues(t, 3, len(buckets[1].Frames))
}

func newDay(day, hour, minute int) *time.Time {
	date := time.Date(2018, time.December, day, hour, minute, 0, 0, time.Local)
	return &date
}

func newDate(year int, month time.Month, day, hour, minute int) *time.Time {
	date := time.Date(year, month, day, hour, minute, 0, 0, time.Local)
	return &date
}
