package worklog

import (
	"github.com/LastSprint/JiraGoIssues/models"
	"time"
)

// WorklogItem сущность, хранящая информацию об одном логе времени
type WorklogItem struct {
	Author models.UserReferenceEntity `json:"author"`
	UpdateAuthor models.UserReferenceEntity `json:"updateAuthor"`
	TimeSpentSeconds int `json:"timeSpentSeconds"`
	Updated time.Time `json:"updated"`
}