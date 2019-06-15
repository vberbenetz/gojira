package models

/**
 Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/#api-rest-api-3-auditing-record-get
 */

type AuditRecords struct {
	Offset int32 `json:"offset"`
	Limit int32 `json:"limit"`
	Total int64 `json:"total"`
	Records []AuditRecordBean `json:"records"`
}

type AuditRecordBean struct {
	Id int64 `json:"id"`
	Summary string `json:"summary"`
	RemoteAccess string `json:"remoteAccess"`
	AuthorKey string `json:"authorKey"`
	Created string `json:"created"`
	Category string `json:"category"`
	EventSource string `json:"eventSource"`
	Description string `json:"description"`
	ObjectItem AssociatedItemBean `json:"objectItem"`
	ChangedValues []ChangedValueBean `json:"changedValues"`
	AssociatedItems []AssociatedItemBean `json:"associatedItemBean"`
}

type AssociatedItemBean struct {
	Id string `json:"id"`
	Name string `json:"name"`
	TypeName string `json:"typeName"`
	ParentId string `json:"parentId"`
	ParentName string `json:"parentName"`
}

type ChangedValueBean struct {
	FieldName string `json:"fieldName"`
	ChangedFrom string `json:"changedFrom"`
	ChangedTo string `json:"changedTo"`
}