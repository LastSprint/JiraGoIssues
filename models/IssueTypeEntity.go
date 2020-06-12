package models

// IssueTypeEntity represent typeof Jira issue
type IssueTypeEntity struct {
	AvatarId    int    `json:"avatarId"`
	Description string `json:"description"`
	IconUrl     string `json:"iconUrl"`
	Name        string `json:"name"`
	ID          string `json:"id"`
}

const (
	// IssueTypeTestExecuted тип jira issue `Test Executed`
	IssueTypeTestExecuted string = "10102"
	// IssueTypeServiceTask тип jira issue `Service Task`
	IssueTypeServiceTask string = "10012"
	// IssueTypeTask тип jira issue `Task`
	IssueTypeTask string = "10002"
	// IssueTypeBug тип jira issue `Bug`
	IssueTypeBug string = "10004"

	IssueTypeEpic string = "10000"
)
