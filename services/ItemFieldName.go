package services

type JiraField string

const (
	// JiraFieldKey Ключ Jira-итема.
	JiraFieldKey JiraField = "key"
	// JiraFieldAssignee тот, на кого назначен Jira-итем.
	JiraFieldAssignee JiraField = "assignee"
	// JiraFieldCreator хз что это.
	JiraFieldCreator JiraField = "creator"
	// JiraFieldReporter тот, кто создал Jira-итем.
	JiraFieldReporter JiraField = "reporter"
	// JiraFieldIssueType тип Jira-итема.
	JiraFieldIssueType JiraField = "issuetype"
	// JiraFieldStatus статус Jira-итема.
	JiraFieldStatus JiraField = "status"
	// JiraFieldReason поле Reason Jira-итема
	JiraFieldReason JiraField = "Reason"
	// JiraFieldSummary Название Jira-итема.
	JiraFieldSummary JiraField = "summary"
	// JiraFieldRemaining оценка Jira-итема.
	JiraFieldRemaining JiraField = "timeestimate"
	// JiraFieldTimeOriginalEstimate оригинальная (первая) оценка Jira-итема.
	JiraFieldTimeOriginalEstimate JiraField = "timeoriginalestimate"
	// JiraFieldTimeSpent затраченное время.
	JiraFieldTimeSpent JiraField = "timespent"
	// JiraFieldPriority приоритет Jira-итема.
	JiraFieldPriority JiraField = "priority"
	// JiraFieldProject имя или ключ проекта в Jira.
	JiraFieldProject JiraField = "project"
	// JiraFieldEpicLink эпик к которому привязана сущность в Jira.
	JiraFieldEpicLink JiraField = "Epic Link"
	// JiraFieldComment это верхнеуровневое поле, которое хранит данные о комментариях.
	JiraFieldComment JiraField = "comment"
	// JiraFieldBoard доска отдела Jira
	JiraFieldBoard JiraField = "board"
	// JiraFieldVCSInfo поле, которое содержит информацию о работе плагина,
	// который показывает были ли в задачу ПРы, коммиты и т.д.
	//
	// Данные приходят в странном формате.
	JiraFieldVCSInfo JiraField = "customfield_10100"

	JiraFieldEpicName JiraField = "Epic Name"

	JiraFieldEpicNameField JiraField = "customfield_10005"

	JiraFieldUpdatedDate JiraField = "updatedDate"
	JiraFieldCreatedDate JiraField = "createdDate"

	JiraFieldAll JiraField = "*all"
	JiraFieldGitCommitReferenced = "customfield_11701"
)

var acceptedFields = []JiraField{
	JiraFieldKey,
	JiraFieldAssignee,
	JiraFieldCreator,
	JiraFieldReporter,
	JiraFieldIssueType,
	JiraFieldStatus,
	JiraFieldReason,
	JiraFieldSummary,
	JiraFieldRemaining,
	JiraFieldTimeOriginalEstimate,
	JiraFieldTimeSpent,
	JiraFieldPriority,
	JiraFieldProject,
}

func (field JiraField) Str() string {
	return string(field)
}

func Str(fields []JiraField) []string {
	res := make([]string, len(fields))

	for i, item := range fields {
		res[i] = item.Str()
	}

	return res
}
