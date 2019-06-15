package models

type ProjectRole struct {
	Id int64 `json:"id"`
	Self string `json:"self"`
	Name string `json:"name"`
	Description string `json:"description"`
	Actors []RoleActor `json:"actors"`
	Scope Scope `json:"scope"`
}
