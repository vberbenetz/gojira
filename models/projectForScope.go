package models

import (
	"../enums"
)

type ProjectForScope struct {
	Id string `json:"id"`
	Self string `json:"self"`
	Key string `json:"key"`
	Name string `json:"name"`
	ProjectTypeKey enums.ProjectTypeKey `json:"projectTypeKey"`
	Simplified bool `json:"simplified"`
	AvatarUrls AvatarUrlsBean `json:"avatarUrls"`
	ProjectCategory UpdatedProjectCategory `json:"projectCategory"`
}
