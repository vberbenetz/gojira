package models

type RoleActor struct {
	Id int64 `json:"id"`
	DisplayName string `json:"displayName"`
	Type string `json:"type"`
	Name string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
	ActorUser ProjectRoleUserBean `json:"actorUser"`
	ActorGroup ProjectRoleGroupBean `json:"actorGroup"`
	User string `json:"user"`
}
