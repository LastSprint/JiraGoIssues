package services

// RequestConvertible интерфейс для сущности, которая может быть представлена как запрос в Jira.
type RequestConvertible interface {
	// GetUseOnlyAdditionalFields возвращает значение, указывающее,
	// использовать ли общие поля, или же использовать только те, что были переданы в `GetAdditionFields`
	GetUseOnlyAdditionalFields() bool
	// MakeJiraRequest явно вызывает конвертацию сущности в запрос.
	MakeJiraRequest() string
	// AdditionFields возвращает список полей, которые ожидается получить от запроса. Этот список дополняется к списку по-умолчанию.
	GetAdditionFields() []JiraField
}

// OrderingType тип сортировки результата из Jira
type OrderingType string

const (
	// Ascending сортировать по возрастанию.
	Ascending OrderingType = "ASC"
	// Descending сортировать по убыванию.
	Descending OrderingType = "DESC"
)

// OrderingModel модель для определения сортировки для запроса в Jira
type OrderingModel struct {
	// OrderTarget поле, по значению которого нужно сортировать.
	OrderTarget string
	// Type тип сортировки.
	Type OrderingType
}

// MakeJiraRequest конвертирует модель в запись вида `ORDER BY $field$ $orderType$`
func (ordering OrderingModel) MakeJiraRequest() string {
	return "ORDER BY " + ordering.OrderTarget + " " + string(ordering.Type)
}
