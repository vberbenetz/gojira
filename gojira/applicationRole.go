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
	Groups []string `json:"groups,omitempty"`
	Name string `json:"name,omitempty"`
	DefaultGroups []string `json:"defaultGroups,omitempty"`
	SelectedByDefault bool `json:"selectedByDefault,omitempty"`
	Defined bool `json:"defined,omitempty"`
	NumberOfSeats int32 `json:"numberOfSeats,omitempty"`
	RemainingSeats int32 `json:"remainingSeats,omitempty"`
	UserCount int32 `json:"userCount,omitempty"`
	UserCountDescription string `json:"userCountDescription,omitempty"`
	HasUnlimitedSeats bool `json:"hasUnlimitedSeats,omitempty"`
	Platform bool `json:"platform,omitempty"`
}

// List returns a list of ApplicationRoles via the GET endpoint
func (s *ApplicationRoleService) List (ctx context.Context) ([]*ApplicationRole, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "applicationrole", nil, nil)
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
