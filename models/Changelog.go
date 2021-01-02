package models

import (
	"strings"
	"time"
)

type FieldType string

type IssueChangelogDate struct {
	Time time.Time
}

const (
	WorkLog string = "WorklogId"
	Status  string = "status"
	TimeSpent string = "timespent"
)

type ChangeLogWrapper struct {
	Histories []HistoryEntity `json:"histories"`
}

type HistoryEntity struct {
	CreatedDateTime IssueChangelogDate `json:"created"`
	HistoryItems    []HistoryItem      `json:"items"`
}

type HistoryItem struct {
	FieldType string  `json:"field"`
	From      string  `json:"fromString"`
	To        string  `json:"toString"`
}

func (date *IssueChangelogDate) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)
	newTime, err := time.Parse("2006-01-02T15:04:05.000-0700", strInput)
	if err != nil {
		return err
	}

	date.Time = newTime
	return nil
}