package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	mj "github.com/LastSprint/JiraGoIssues/models"
)

// JiraIssueService сервис для работы с задачами из Jira.
type JiraIssueService interface {
	// LoadIssues Загружает все задачи для конкретного пользователя
	// Parameters:
	// 	- request: Запрос в джиру.
	LoadIssues(request RequestConvertible) (mj.IssueSearchWrapperEntity, error)
}

// JiraIssueLoader загрузчик задач из Jira.
type JiraIssueLoader struct {
	authToken string
	url       string
}

// NewJiraIssueLoader создает экземпляр загрузчка задач из Jira.
// Parameters:
//	- url: URL-адрес jira API.
//	- authToken: BasicAuth токен
func NewJiraIssueLoader(url, authToken string) *JiraIssueLoader {
	return &JiraIssueLoader{
		url:       url,
		authToken: authToken,
	}
}

// LoadIssues Загружает все задачи для конкретного пользователя
// Parameters:
// 	- user: Пользователь, для которого нужно выгрузить задачи
//	- statuses: Статусы, в которых должны быть задачи
func (service *JiraIssueLoader) LoadIssues(request RequestConvertible) (mj.IssueSearchWrapperEntity, error) {
	result := mj.IssueSearchWrapperEntity{}

	req, err := http.NewRequest(
		"GET", service.url, nil,
	)

	if err != nil {
		return result, err
	}

	h := req.Header

	h.Set("Authorization", "Basic "+service.authToken)
	req.Header = h

	qr := req.URL.Query()

	jql := request.MakeJiraRequest()
	fmt.Println(jql)

	var fields []JiraField

	if request.GetUseOnlyAdditionalFields() {
		fields = request.GetAdditionFields()
	} else {
		fields = append(acceptedFields, request.GetAdditionFields()...)
	}

	qr.Set("jql", jql)
	qr.Set("startAt", strconv.Itoa(request.GetStartAtIndex()))
	qr.Set("maxResults", strconv.Itoa(request.GetPageSize()))
	qr.Set("fields", joinByCharacter(Str(fields), ",", ""))
	qr.Set("expand", "changelog")

	req.URL.RawQuery = qr.Encode()
	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	bd, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return result, err
	}

	err = json.Unmarshal(bd, &result)

	if err != nil {
		return result, err
	}

	return result, err
}

// joinByCharacter объединяет массив строк через символ `character`
// и может "окружить" каждый элемент массива значением`surroundBy`
func joinByCharacter(arr []string, delim string, surroundBy string) string {
	result := ""

	arrLen := len(arr)

	for i, item := range arr {

		result += surroundBy + item + surroundBy

		if i < arrLen-1 {
			result += delim
		}
	}

	return result
}
