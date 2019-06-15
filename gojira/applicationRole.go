package gojira

import (
	"context"
	"net/http"
)

type ApplicationRoleService service

/**
 Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/#api-rest-api-3-applicationrole-get
 */
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
