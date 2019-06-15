package models

import (
	"../enums"
)

type FoundUsersAndGroups struct {
	Users FoundUsers `json:"users"`
	Groups FoundGroups `json:"groups"`
}

type FoundUsers struct {
	Users []UserPickerUser `json:"users"`
	Total int32 `json:"total"`
	Header string `json:"header"`
}

type FoundGroups struct {
	Header string `json:"header"`
	Total int32 `json:"total"`
	Groups []FoundGroup `json:"groups"`
}

type UserPickerUser struct {
	AccountId string `json:"accountId"`
	Name string `json:"name"`
	Key string `json:"key"`
	Html string `json:"html"`
	DisplayName string `json:"displayName"`
	AvatarUrl string `json:"avatarUrl"`
}

type FoundGroup struct {
	Name string `json:"name"`
	Html string `json:"html"`
	Labels []GroupLabel `json:"labels"`
}

type GroupLabel struct {
	Text string `json:"text"`
	Title string `json:"title"`
	Type enums.GroupLabelType `json:"type"`
}
