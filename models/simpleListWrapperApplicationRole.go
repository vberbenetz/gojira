package models

type SimpleListWrapperApplicationRole struct {
	Size int32 `json:"size"`
	Items []ApplicationRole `json:"items"`
	PagingCallback string `json:"pagingCallback"`
	Callback string `json:"callback"`
	MaxResults int32 `json:"max-results"`
}
