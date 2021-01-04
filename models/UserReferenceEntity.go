package models

type Avatars struct {
	LargestUrl  string `json:"48x48"`
	MidUrl      string `json:"32x32"`
	SmallUrl    string `json:"24x24"`
	SmallestUrl string `json:"16x16"`
}

// UserReferenceEntity is short user entity
type UserReferenceEntity struct {
	DisplayName string  `json:"displayName"`
	Key         string  `json:"key"`
	Name        string  `json:"name"`
	AvatarUrls  Avatars `json:"avatarUrls"`
}
