package gojira

import (
	"context"
	"net/http"
	"time"
)

// AuditRecordsService is used to execute requests pertaining to AuditRecords
type AuditRecordsService service

// AuditRecords defines the response body returned by retrieving all the Audit Records
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/#api-group-Audit-records
type AuditRecords struct {
	Offset int32 `json:"offset"`
	Limit int32 `json:"limit"`
	Total int64 `json:"total"`
	Records []AuditRecordBean `json:"records"`
}

// AuditRecordBean is the actual AuditRecord object returned from retrieving all the Audit Records
type AuditRecordBean struct {
	ID int64 `json:"id"`
	Summary string `json:"summary,omitempty"`
	RemoteAccess string `json:"remoteAccess,omitempty"`
	AuthorKey string `json:"authorKey,omitempty"`
	Created string `json:"created,omitempty"`
	Category string `json:"category,omitempty"`
	EventSource string `json:"eventSource,omitempty"`
	Description string `json:"description,omitempty"`
	ObjectItem AssociatedItemBean `json:"objectItem,omitempty"`
	ChangedValues []ChangedValueBean `json:"changedValues,omitempty"`
	AssociatedItems []AssociatedItemBean `json:"associatedItems,omitempty"`
}

// AssociatedItemBean is a nested object within AuditRecordBean
type AssociatedItemBean struct {
	ID string `json:"id"`
	Name string `json:"name,omitempty"`
	TypeName string `json:"typeName,omitempty"`
	ParentID string `json:"parentId,omitempty"`
	ParentName string `json:"parentName,omitempty"`
}

// ChangedValueBean is a nested object within AuditRecordBean
type ChangedValueBean struct {
	FieldName string `json:"fieldName,omitempty"`
	ChangedFrom string `json:"changedFrom,omitempty"`
	ChangedTo string `json:"changedTo,omitempty"`
}

// AuditRecordsQueryParams is the object defining the possible query parameters
// for the request retrieving all AuditRecords
//
// *** NOTE *** Since parameters are omitted if empty, Zero Valued variables will be excluded from the query param.
// Ex: Limit = 0 will not be present as a query parameter
type AuditRecordsQueryParams struct {
	Offset int32 `url:"offset,omitempty"`
	Limit int32 `url:"limit,omitempty"`
	Filter string `url:"filter,omitempty"`
	From time.Time `url:"from,omitempty"`
	To time.Time `url:"to,omitempty"`
}

// List returns a list of AuditRecords via the GET endpoint
// The response also has paging parameters in the top level object
func (s *AuditRecordsService) List (
	ctx context.Context, queryParams *AuditRecordsQueryParams) (*AuditRecords, *http.Response, error) {

	u, err := addQueryParams("auditing/record", queryParams)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	ar := &AuditRecords{}
	resp, err := s.client.Do(ctx, req, ar)
	if err != nil {
		return nil, resp, err
	}

	return ar, resp, nil
}
