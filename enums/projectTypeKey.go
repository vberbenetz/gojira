package enums

type ProjectTypeKey int

const (
	ProjectKeyType_Ops ProjectTypeKey = iota
	ProjectKeyType_Software
	ProjectKeyType_ServiceDesk
	ProjectKeyType_Business
)
