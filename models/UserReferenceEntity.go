package models

type Avatars struct {
	LargestUrl string `json:"48x48"`
}

// UserReferenceEntity is short user entity
type UserReferenceEntity struct {
	DisplayName string  `json:"displayName"`
	Key         string  `json:"key"`
	Name        string  `json:"name"`
	AvatarUrls  Avatars `json:"avatarUrls"`
}
