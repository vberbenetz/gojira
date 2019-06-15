package models

type Dashboard struct {
	Id string `json:"id"`
	IsFavourite bool `json:"isFavourite"`
	Name string `json:"name"`
	Owner UserBean `json:"owner"`
	Popularity int64 `json:"popularity"`
	Rank int32 `json:"rank"`
	Self string `json:"self"`
	SharePermissions []SharePermission `json:"sharePermissions"`
	View string `json:"view"`
}

type PageOfDashboards struct {
	StartAt int32 `json:"startAt"`
	MaxResults int32 `json:"maxResults"`
	Total int32 `json:"total"`
	Prev string `json:"prev"`
	Next string `json:"next"`
	Dashboards []Dashboard `json:"dashboards"`
}

type PageBeanDashboard struct {
	Self string `json:"self"`
	NextPage string `json:"nextPage"`
	MaxResults int32 `json:"maxResults"`
	StartAt int64 `json:"startAt"`
	Total int64 `json:"total"`
	IsLast bool `json:"isLast"`
	Values []Dashboard `json:"values"`
}

type PropertyKey struct {
	Self string `json:"self"`
	Key string `json:"key"`
}

type PropertyKeys struct {
	Keys []PropertyKey `json:"keys"`
}

type EntityProperty struct {
	Key string `json:"key"`
	Value Any `json:"value"`
}