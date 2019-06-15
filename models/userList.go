package models

type UserList struct {
	Size int32 `json:"size"`
	Items []User `json:"items"`
	MaxResults int32 `json:"max-results"`
	StartIndex int32 `json:"start-index"`
	EndIndex int32 `json:"end-index"`
}
