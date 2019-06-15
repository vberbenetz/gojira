package models

/**
 Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/#api-rest-api-3-avatar-type-system-get
 */
type Avatar struct {
	Id string `json:"id"`
	Owner string `json:"owner"`
	IsSystemAvatar bool `json:"isSystemAvatar"`
	IsSelected bool `json:"isSelected"`
	IsDeletable bool `json:"isDeletable"`
	FileName string `json:"fileName"`
	Urls Any `json:"urls"`
}

type SystemAvatars struct {
	System []Avatar `json:"system"`
}

type Any interface {}