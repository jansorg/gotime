package report

import (
	"time"

	"github.com/jansorg/tom/go-tom/dateTime"
)

type Config struct {
	Splitting          []SplitOperation        `json:"split"`
	Timezone           *time.Location          `json:"timezone"`
	ProjectIDs         []string                `json:"projects"`
	ProjectDelimiter   string                  `json:"project_delimiter"`
	IncludeSubprojects bool                    `json:"show_subprojects"`
	DateFilterRange    dateTime.DateRange      `json:"date_range"`
	ShowEmpty          bool                    `json:"show_empty"`
	IncludeArchived    bool                    `json:"include_archived"`
	ShortTitles        bool                    `json:"short_titles"`
	EntryRounding      dateTime.RoundingConfig `json:"rounding_entry"`
	SumRounding        dateTime.RoundingConfig `json:"rounding_total"`
}
