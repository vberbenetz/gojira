package gojira

type Project struct {
	ID string `json:"id"`
	Expand string `json:"expand"`
	Self string `json:"self"`
	Key string `json:"key"`
	Description string `json:"description"`
	Lead User `json:"lead"`
	Components []Component `json:"components"`
	IssueTypes []IssueTypeBean `json:"issueTypes"`
	Url string `json:"url"`
	Email string `json:"email"`
	AssigneeType string `json:"assigneeType"`
	Versions []Version `json:"versions"`
	Name string `json:"name"`
	Roles interface{} `json:"roles"`
	AvatarUrls map[string]string `json:"avatarUrls"`
	ProjectCategory ProjectCategory `json:"projectCategory"`
	ProjectTypeKey string `json:"projectTypeKey"`
	Simplified bool `json:"simplified"`
	Style string `json:"style"`
	IsPrivate bool `json:"isPrivate"`
	IssueTypeHierarchy Hierarchy `json:"issueTypeHierarchy"`
	Permissions ProjectPermissions `json:"permissions"`
	Properties interface{} `json:"properties"`
	Uuid string `json:"uuid"`
}

type Component struct {
	ID string `json:"id"`
	Self string `json:"self,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Lead User `json:"lead,omitempty"`
	LeadUserName string `json:"leadUserName,omitempty"`
	LeadAccountId string `json:"leadAccountId,omitempty"`
	AssigneeType string `json:"assigneeType,omitempty"`
	Assignee User `json:"assignee,omitempty"`
	RealAssigneeType string `json:"realAssigneeType,omitempty"`
	RealAssignee User `json:"realAssignee,omitempty"`
	IsAssigneeTypeValid bool `json:"isAssigneeTypeValid,omitempty"`
	Project string `json:"project,omitempty"`
	ProjectID int64 `json:"projectId,omitempty"`
}

type IssueTypeBean struct {
	Self string `json:"self"`
	Id string `json:"id"`
	Description string `json:"description"`
	IconUrl string `json:"iconUrl"`
	Name string `json:"name"`
	Subtask bool `json:"subtask"`
	AvatarId int64 `json:"avatarId"`
	Scope Scope `json:"scope"`
}

type Version struct {
	Expand string `json:"expand,omitempty"`
	Self string `json:"self,omitempty"`
	ID string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	Archived bool `json:"archived,omitempty"`
	Released bool `json:"released,omitempty"`
	StartDate string `json:"startDate,omitempty"`
	ReleaseDate string `json:"releaseDate,omitempty"`
	Overdue bool `json:"overdue,omitempty"`
	UserStartDate string `json:"userStartDate,omitempty"`
	Project string `json:"project,omitempty"`
	ProjectId int64 `json:"projectId,omitempty"`
	MoveUnfixedIssuesTo string `json:"moveUnfixedIssuesTo,omitempty"`
	Operations []SimpleLink `json:"operations,omitempty"`
	IssuesStatusForFixVersion VersionIssuesStatus `json:"issuesStatusForFixVersion,omitempty"`
}

type SimpleLink struct {
	ID string `json:"id,omitempty"`
	StyleClass string `json:"styleClass,omitempty"`
	IconClass string `json:"iconClass,omitempty"`
	Label string `json:"label,omitempty"`
	Title string `json:"title,omitempty"`
	Href string `json:"href,omitempty"`
	Weight int32 `json:"weight,omitempty"`
}

type VersionIssuesStatus struct {
	Unmapped int64 `json:"unmapped,omitempty"`
	ToDo int64 `json:"toDo,omitempty"`
	InProgress int64 `json:"inProgress,omitempty"`
	Done int64 `json:"done,omitempty"`
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty"`
}

type ProjectCategory struct {
	Self string `json:"self,omitempty"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Hierarchy struct {
	Level []HierarchyLevel `json:"level,omitempty"`
}

type HierarchyLevel struct {
	ID int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	AboveLevelID int64 `json:"aboveLevelId,omitempty"`
	BelowLevelID int64 `json:"belowLevelId,omitempty"`
	ProjectConfigurationID int64 `json:"projectConfigurationId,omitempty"`
	IssueTypeIDs []int64 `json:"issueTypeIds,omitempty"`
	ExternalUuid string `json:"externalUuid,omitempty"`
}

type ProjectPermissions struct {
	CanEdit bool `json:"canEdit,omitempty"`
}

type ProjectRole struct {
	Id int64 `json:"id"`
	Self string `json:"self"`
	Name string `json:"name"`
	Description string `json:"description"`
	Actors []RoleActor `json:"actors"`
	Scope Scope `json:"scope"`
}

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

type Scope struct {
	Type string `json:"type"`
	Project ProjectForScope `json:"project"`
	AdditionalProperties interface{} `json:"additionalProperties"`
}

type DefaultShareScope struct {
	Scope string `json:"scope"`
}

type ProjectRoleUserBean struct {
	AccountId string `json:"accountId"`
}

type ProjectRoleGroupBean struct {
	DisplayName string `json:"displayName"`
	Name string `json:"name"`
}

type ProjectForScope struct {
	Id string `json:"id"`
	Self string `json:"self"`
	Key string `json:"key"`
	Name string `json:"name"`
	ProjectTypeKey string `json:"projectTypeKey"`
	Simplified bool `json:"simplified"`
	AvatarUrls map[string]string `json:"avatarUrls"`
	ProjectCategory UpdatedProjectCategory `json:"projectCategory"`
}

type UpdatedProjectCategory struct {
	Id string `json:"id"`
	Self string `json:"self"`
	Description string `json:"description"`
	Name string `json:"name"`
}