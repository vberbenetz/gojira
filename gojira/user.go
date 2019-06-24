package gojira

type User struct {
	Self string `json:"self"`
	Key string `json:"key"`
	AccountID string `json:"accountId"`
	AccountType string `json:"accountType"`
	Name string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	AvatarUrls map[string]string `json:"avatarUrls"`
	DisplayName string `json:"displayName"`
	Active bool `json:"active"`
	TimeZone string `json:"timeZone"`
	Locale string `json:"locale"`
	Groups SimpleListWrapperGroupName `json:"groups"`
	ApplicationRoles SimpleListWrapperApplicationRole `json:"applicationRoles"`
	Expand string `json:"expand"`
}

type SimpleListWrapperApplicationRole struct {
	Size int32 `json:"size"`
	Items []ApplicationRole `json:"items"`
	PagingCallback string `json:"pagingCallback"`
	Callback string `json:"callback"`
	MaxResults int32 `json:"max-results"`
}

type SimpleListWrapperGroupName struct {
	Size int32 `json:"size"`
	Items []GroupName `json:"items"`
	PagingCallback string `json:"pagingCallback"`
	Callback string `json:"callback"`
	MaxResults int32 `json:"max-results"`
}
