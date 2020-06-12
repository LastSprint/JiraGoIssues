package Stubs

import (
	"github.com/LastSprint/pkg_go_jira_service/models"
	"github.com/LastSprint/pkg_go_jira_service/services"
)

// JiraIssueServiceStubs заглушка для тестов
type JiraIssueServiceStubs struct {
	// Request Тут будет указан запрос, который был передан в `LoadIssues`
	Request *services.RequestConvertible
	// Result ответ LoadIssues. Если равен nil, то вернется ошибка
	Result *models.IssueSearchWrapperEntity
	// Err ошибка, котора вернется из LoadIssues
	Err error
}

func (j *JiraIssueServiceStubs) LoadIssues(request services.RequestConvertible) (models.IssueSearchWrapperEntity, error) {
	j.Request = &request

	if j.Result == nil {
		return models.IssueSearchWrapperEntity{}, j.Err
	}

	return *j.Result, nil
}
