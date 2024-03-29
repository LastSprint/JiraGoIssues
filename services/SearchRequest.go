package services

type JiraFiledCompareOperation string

const (
	Equal         JiraFiledCompareOperation = "="
	GraterOrEqual JiraFiledCompareOperation = ">="
	LessOrEqual   JiraFiledCompareOperation = "<="
	NotEqual      JiraFiledCompareOperation = "!="
)

type ComparableValue struct {
	Value            string
	CompareOperation JiraFiledCompareOperation
}

func (val *ComparableValue) JqlString() string {
	return string(val.CompareOperation) + " " + val.Value
}

// SearchRequest Модель для запроса в Jira.
type SearchRequest struct {
	PageSize     int
	StartAtIndex int

	// UseOnlyAdditionalFields если установлено значение `true`,
	// то в запросе будут использованы ТОЛЬКО поля, которые содержаться в этой модели реквеста
	// в поле `AdditionFields`
	//
	// В противном случае поведение стандартное.
	UseOnlyAdditionalFields bool

	// 	Board это доска определенного отдела у проекта в Jira.
	Board string
	// IncludedStatuses статусы Jira-итемов, которые НУЖНО получить в ответ на этот запрос.
	// Константны описаны в пакете jira/models
	IncludedStatuses []string
	// ExcludedStatuses статусы Jira-итемов, которые НЕ НУЖНО получать в ответ на этот запрос.
	// Константны описаны в пакете jira/models
	ExcludedStatuses []string
	// IncludedTypes типы Jira-итемов, которые НУЖНО получить в ответ на этот запрос.
	// Константны описаны в пакете jira/models
	IncludedTypes []string
	// ExcludedTypes типы Jira-итемов, которые НЕ НУЖНО получать в этот на этот запрос.
	// Константны описаны в пакете jira/models
	ExcludedTypes []string
	// Значения поля Reason (причина непопадания в оценки), задачи с которыми нужно включить в выдачу
	IncludedReasons []string
	// Priorities приортеты Jira-итемов, которые нужно включить в выдачу.
	Priorities []string
	// Assignee тот, на кого назначен Jira-итем.
	Assignee string
	// ProjectID id проекта в Jira.
	ProjectID string
	// ProjectsToExclude содержит ключи проектов в Jira которые НЕ нужно включать в выдачу. Если список пуст, то он не будет использован в запросе.
	ProjectsToExclude []string
	// EpicLink название эпика, к которому относятся запрошенные задачи.
	EpicLink string
	// Ordering опциональная сортировка.
	Ordering *OrderingModel
	// AdditionFields дополнительные поля, которые ожидается получить в ответе.
	AdditionFields []JiraField
	// EpicName будет выполнять поиск по операции CONTAINS
	EpicName string
	// Дата обновления issue
	// Указывается в формате yyy-mm-dd
	UpdatedDate *ComparableValue

	Labels []string

	// Дата создания issue
	// Указывается в формате yyy-mm-dd
	CreatedDate *ComparableValue

	ChangelogSearch []ChangelogSearch
}

func (req SearchRequest) GetAdditionFields() []JiraField {
	return req.AdditionFields
}

func (req SearchRequest) GetUseOnlyAdditionalFields() bool {
	return req.UseOnlyAdditionalFields
}

// MakeJiraRequest конвертирует структуру в строку JQL запроса.
func (req SearchRequest) MakeJiraRequest() string {

	result := []string{}

	if len(req.Board) != 0 {
		result = append(result, JiraFieldBoard.Str()+" = "+req.Board)
	}

	if len(req.Assignee) != 0 {
		result = append(result, JiraFieldAssignee.Str()+" = "+req.Assignee)
	}

	if len(req.ProjectID) != 0 {
		result = append(result, JiraFieldProject.Str()+" = "+req.ProjectID)
	}

	if len(req.EpicLink) != 0 {
		result = append(result, "\""+JiraFieldEpicLink.Str()+"\""+" = "+req.EpicLink)
	}

	if len(req.EpicName) != 0 {
		result = append(result, "\""+JiraFieldEpicName.Str()+"\""+" ~ "+"\""+req.EpicName+"\"")
	}

	if str := joinByCharacter(req.ProjectsToExclude, ",", "\""); len(str) != 0 {
		result = append(result, JiraFieldProject.Str()+" not in ("+str+")")
	}

	if str := joinByCharacter(req.IncludedStatuses, ",", "\""); len(str) != 0 {
		result = append(result, JiraFieldStatus.Str()+" in ("+str+")")
	}

	if str := joinByCharacter(req.ExcludedStatuses, ",", "\""); len(str) != 0 {
		result = append(result, JiraFieldStatus.Str()+" not in ("+str+")")
	}

	if str := joinByCharacter(req.IncludedReasons, ",", "\""); len(str) != 0 {
		result = append(result, JiraFieldReason.Str()+" in ("+str+")")
	}

	if str := joinByCharacter(req.IncludedTypes, ",", "\""); len(str) != 0 {
		result = append(result, JiraFieldIssueType.Str()+" in ("+str+")")
	}

	if str := joinByCharacter(req.ExcludedTypes, ",", "\""); len(str) != 0 {
		result = append(result, JiraFieldIssueType.Str()+" not in ("+str+")")
	}

	if str := joinByCharacter(req.Labels, ",", "\""); len(str) != 0 {
		result = append(result, JiraFiledLabels.Str()+" in ("+str+")")
	}

	if str := joinByCharacter(req.Priorities, ",", "\""); len(str) != 0 {
		result = append(result, JiraFieldPriority.Str()+" in ("+str+")")
	}

	if req.UpdatedDate != nil {
		result = append(result, JiraFieldUpdatedDate.Str()+" "+req.UpdatedDate.JqlString())
	}

	if req.CreatedDate != nil {
		result = append(result, JiraFieldCreatedDate.Str()+" "+req.CreatedDate.JqlString())
	}

	if len(req.ChangelogSearch) != 0 {
		for _, item := range req.ChangelogSearch {
			result = append(result, item.QueryString())
		}
	}

	str := joinByCharacter(result, " and ", "")

	if order := req.Ordering; order != nil {
		str += " " + order.MakeJiraRequest()
	}

	return str
}

func (req SearchRequest) GetPageSize() int {
	return req.PageSize
}

func (req SearchRequest) GetStartAtIndex() int {
	return req.StartAtIndex
}
