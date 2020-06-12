package models

type ProjectEntity struct {
	Self     string         `json:"self"`
	Id       string         `json:"id"`
	Key      string         `json:"key"`
	Name     string         `json:"name"`
	Category CategoryEntity `json:"projectCategory"`
}

type CategoryEntity struct {
	Self        string `json:"self"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}