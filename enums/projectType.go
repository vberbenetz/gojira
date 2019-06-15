package enums

type ProjectType int

const (
	ProjectType_Authenticated ProjectType = iota
	ProjectType_Group
	ProjectType_Project
	ProjectType_ProjectRole
	ProjectType_Global
	ProjectType_LoggedIn
	ProjectType_ProjectUnknown
)