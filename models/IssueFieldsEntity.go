package models

import "fmt"

// IssueFieldsEntity contains needed issue fields
type IssueFieldsEntity struct {
	Assignee UserReferenceEntity `json:"assignee"`
	Creator  UserReferenceEntity `json:"creator"`
	Reporter UserReferenceEntity `json:"reporter"`

	Issuetype IssueTypeEntity `json:"issuetype"`
	Status    StatusEntity    `json:"status"`
	Priority  *IssuePriority  `json:"priority"`
	Project   ProjectEntity   `json:"project"`

	Summary   string `json:"summary"`
	Remaining int    `json:"timeestimate"`
	Estimate  int    `json:"timeoriginalestimate"`
	TimeSpend int    `json:"timespent"`

	Reason []string `json:"customfield_11402"`

	Comments CommentWrapper `json:"comment"`

	VCSInfo *string `json:"customfield_10100"`

	EpicLink string `json:"customfield_10008"`
	EpicName string `json:"customfield_10005"`

	IsGitCommitExist string `json:"customfield_11701"`

	Labels []string `json:"labels"`
}

func (model IssueFieldsEntity) FormatedRemaining() string {
	return formateTime(model.Remaining)
}

func (model IssueFieldsEntity) FormatedEstimate() string {
	return formateTime(model.Estimate)
}

func (model IssueFieldsEntity) FormatedTimeSpend() string {
	return formateTime(model.TimeSpend)
}

func formateTime(seconds int) string {
	var minutes = seconds / 60

	if minutes >= 60 {
		hours := float32(minutes) / 60.0
		return fmt.Sprintf("%.2f ч.", hours)
	}
	return fmt.Sprintf("%d мин.", minutes)
}


