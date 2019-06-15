package models

type UserBean struct {
	Key string `json:"key"`
	Self string `json:"self"`
	Name string `json:"name"`
	DisplayName string `json:"displayName"`
	Active string `json:"active"`
	AccountId string `json:"accountId"`
	AvatarUrls UserBeanAvatarUrls `json:"avatarUrls"`
}
