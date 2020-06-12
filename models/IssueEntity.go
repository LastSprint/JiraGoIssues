package models

// IssueEntity represent Jira issue with needed (for me) parameters
type IssueEntity struct {
	Key    string            `json:"key"`
	Fields IssueFieldsEntity `json:"fields"`
	Changelog ChangeLogWrapper  `json:"changelog"`
}
