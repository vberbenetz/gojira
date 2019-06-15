package models

type Group struct {
	Name string `json:"name"`
	Self string `json:"self"`
	Users PagedListUserDetailsApplicationUser `json:"users"`
	Expand string `json:"expand"`
}

type PagedListUserDetailsApplicationUser struct {
	Size int32 `json:"size"`
	Items []UserDetails `json:"items"`
	MaxResults int32 `json:"max-results"`
	StartIndex int32 `json:"start-index"`
	EndIndex int32 `json:"end-index"`
}

type UserDetails struct {
	Self string `json:"self"`
	Name string `json:"name"`
	Key string `json:"key"`
	AccountId string `json:"accountId"`
	EmailAddress string `json:"emailAddress"`
	AvatarUrls AvatarUrlsBean `json:"avatarUrls"`
	DisplayName string `json:"displayName"`
	Active bool `json:"active"`
	TimeZone string `json:"timeZone"`
	AccountType string `json:"accountType"`
}

type PageBeanUserDetails struct {
	Self string `json:"self"`
	NextPage string `json:"nextPage"`
	MaxResults int32 `json:"maxResults"`
	StartAt int64 `json:"startAt"`
	Total int64 `json:"total"`
	IsLast bool `json:"isLast"`
	Values []UserDetails `json:"values"`
}
