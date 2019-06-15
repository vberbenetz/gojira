package models

type SimpleListWrapperGroupName struct {
	Size int32 `json:"size"`
	Items []GroupName `json:"items"`
	PagingCallback string `json:"pagingCallback"`
	Callback string `json:"callback"`
	MaxResults int32 `json:"max-results"`
}
