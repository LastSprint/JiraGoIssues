package models

// IssueSearchWrapperEntity is object that contains issues ans paging information
type IssueSearchWrapperEntity struct {
	StartAt    int           `json:"startAt"`
	MaxResults int           `json:"maxResults"`
	Total      int           `json:"total"`
	Issues     []IssueEntity `json:"issues"`
}
