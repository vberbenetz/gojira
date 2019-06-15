package models

import (
	"../enums"
)

type SharePermission struct {
	Id int64 `json:"id"`
	Type enums.ProjectType `json:"type"`
	Project Project `json:"project"`
	Role ProjectRole `json:"role"`
	Group GroupName `json:"groupName"`
}
