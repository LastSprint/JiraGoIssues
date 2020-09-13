package worklog

import (
	"encoding/json"
	"github.com/LastSprint/JiraGoIssues/models/worklog"
	"net/http"
)

// IssueWorklogService сервис для получения информации о ворклоге issue
type IssueWorklogService interface {
	// Get получает ворклог об одном issue
	// Данные пагинируемы
	//
	// `issueKeyOrId` - идентификатор issue который можно получить из IssueEntity.Fields.
	// Это может быть модификатор задачи `IDPT-123`либо `id` задачи из модели issue.
	Get(issueKeyOrId string) (error, *worklog.RootWorklog)
}

// IssueWorklogLoader реализация IssueWorklogService по-умолчанию.
// Делает запрос `GET /rest/api/2/issue/{issueIdOrKey}/worklog`
// В качестве авторизации использует BasicAuth
type IssueWorklogLoader struct {
	// Базовый URL - адрес хоста Jira
	BaseUrl string
	// Имя пользователя от которого делается запрос. Используется для авторизации
	Username string
	// Пароль пользователя. Используется для авторизации
	Password string
}

func (loader *IssueWorklogLoader) Get(issueKeyOrId string) (*worklog.RootWorklog, error) {
	request, err := http.NewRequest("GET", loader.BaseUrl + "rest/api/2/issue/"+ issueKeyOrId +"/worklog", nil)

	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(loader.Username, loader.Password)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	var result worklog.RootWorklog

	err = json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}