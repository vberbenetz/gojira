package models

import (
	"../enums"
)

type User struct {
	Self string `json:"self"`
	Key string `json:"key"`
	AccountId string `json:"accountId"`
	AccountType enums.AccountType `json:"accountType"`
	Name string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	AvatarUrls AvatarUrlsBean `json:"avatarUrls"`
	DisplayName string `json:"displayName"`
	Active bool `json:"active"`
	TimeZone string `json:"timeZone"`
	Locale string `json:"locale"`
	Groups SimpleListWrapperGroupName `json:"groups"`
	ApplicationRoles SimpleListWrapperApplicationRole `json:"applicationRoles"`
	Expand string `json:"expand"`
}
