package worklog

type RootWorklog struct {
	Total int `json:"total"`
	MaxResults int `json:"maxResults"`
	StartAt int `json:"startAt"`
	Worklogs []WorklogItem `json:"worklogs"`
}