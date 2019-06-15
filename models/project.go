package models

import (
	"../enums"
)

type Project struct {
	Id string `json:"id"`
	Expand string `json:"expand"`
	Self string `json:"self"`
	Key string `json:"key"`
	Description string `json:"description"`
	Lead User `json:"lead"`
	Components []Component `json:"components"`
	IssueTypes []IssueTypeBean `json:"issueTypes"`
	Url string `json:"url"`
	Email string `json:"email"`
	AssigneeType enums.AssigneeType `json:"assigneeType"`
	Versions []Version `json:"versions"`
	Name string `json:"name"`
	Roles Any `json:"roles"`
	AvatarUrls AvatarUrlsBean `json:"avatarUrls"`
	ProjectCategory ProjectCategory `json:"projectCategory"`
	ProjectTypeKey enums.ProjectTypeKey `json:"projectTypeKey"`
	Simplified bool `json:"simplified"`
	Style enums.Style `json:"style"`
	IsPrivate bool `json:"isPrivate"`
	IssueTypeHierarchy Hierarchy `json:"issueTypeHierarchy"`
	Permissions ProjectPermissions `json:"permissions"`
	Properties Any `json:"properties"`
	Uuid string `json:"uuid"`
}
