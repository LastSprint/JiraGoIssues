package models

type Comment struct {
	Author UserReferenceEntity `json:"author"`
	Body   string              `json:"body"`
}
