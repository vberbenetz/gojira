package gojira

import (
	"context"
	"fmt"
	"net/http"
)

// ApplicationRoleService is used to execute requests pertaining to ApplicationRoles
type ApplicationRoleService service

// ApplicationRole is the response body as outlined by the Jira Cloud API
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/#api-group-Application-roles
type ApplicationRole struct {
	Key string `json:"key"`
	Groups []string `json:"groups"`
	Name string `json:"name"`
	DefaultGroups []string `json:"defaultGroups"`
	SelectedByDefault bool `json:"selectedByDefault"`
	Defined bool `json:"defined"`
	NumberOfSeats int32 `json:"numberOfSeats"`
	RemainingSeats int32 `json:"remainingSeats"`
	UserCount int32 `json:"userCount"`
	UserCountDescription string `json:"userCountDescription"`
	HasUnlimitedSeats bool `json:"hasUnlimitedSeats"`
	Platform bool `json:"platform"`
}

// List returns a list of ApplicationRoles via the GET endpoint
func (s *ApplicationRoleService) List (ctx context.Context) ([]*ApplicationRole, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "applicationrole", nil)
	if err != nil {
		return nil, nil, err
	}

	var ar []*ApplicationRole
	resp, err := s.client.Do(ctx, req, &ar)
	if err != nil {
		return nil, resp, err
	}

	return ar, resp, nil
}

// Get returns a single ApplicationRole by the provided key parameter
func (s *ApplicationRoleService) Get (ctx context.Context, key string) (*ApplicationRole, *http.Response, error) {
	endpoint := fmt.Sprintf("applicationrole/%v", key)
	req, err := s.client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var ar *ApplicationRole
	resp, err := s.client.Do(ctx, req, &ar)
	if err != nil {
		return nil, resp, err
	}

	return ar, resp, err
}
