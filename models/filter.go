package models

type Filter struct {
	Id string `json:"id"`
	Self string `json:"self"`
	Name string `json:"name"`
	Description string `json:"description"`
	Owner User `json:"owner"`
	Jql string `json:"jql"`
	ViewUrl string `json:"viewUrl"`
	SearchUrl string `json:"searchUrl"`
	Favourite bool `json:"favourite"`
	FavouritedCount int64 `json:"favouritedCount"`
	SharePermissions []SharePermission `json:"sharePermissions"`
	SharedUsers UserList `json:"sharedUsers"`
	Subscriptions FilterSubscriptionsList `json:"subscriptions"`
}

type FilterSubscriptionsList struct {
	Size int32 `json:"size"`
	Items []FilterSubscription `json:"items"`
	MaxResults int32 `json:"max-results"`
	StartIndex int32 `json:"start-index"`
	EndIndex int32 `json:"end-index"`
}

type FilterSubscription struct {
	Id int64 `json:"id"`
	User User `json:"user"`
	Group GroupName `json:"group"`
}

type GroupName struct {
	Name string `json:"name"`
	Self string `json:"self"`
}

type PageBeanFoundFilter struct {
	Self string `json:"self"`
	NextPage string `json:"nextPage"`
	MaxResults int32 `json:"maxResults"`
	StartAt int64 `json:"startAt"`
	Total int64 `json:"total"`
	IsLast bool `json:"isLast"`
	Values []FoundFilter `json:"values"`
}

type FoundFilter struct {
	Id string `json:"id"`
	Self string `json:"self"`
	Name string `json:"name"`
	Description string `json:"description"`
	Owner User `json:"owner"`
	Jql string `json:"jql"`
	ViewUrl string `json:"viewUrl"`
	SearchUrl string `json:"searchUrl"`
	Favourite bool `json:"favourite"`
	FavouritedCount int64 `json:"favouritedCount"`
	SharePermissions []SharePermission `json:"sharePermissions"`
	Subscriptions []FilterSubscription `json:"subscriptions"`
}

type CountItem struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
