package services

type ChangelogOperation string

const (
	WasIn ChangelogOperation = "was in"
	WasNotIn ChangelogOperation = "was not in"
)

type ChangelogSearch struct {
	FieldName JiraField
	Values []string
	Operation ChangelogOperation
}

func (val ChangelogSearch) QueryString() string {
	result := string(val.Operation) + " "
	values := joinByCharacter(val.Values, ",", `"`)

	return val.FieldName.Str() + " " + result + "(" + values + ")"
}