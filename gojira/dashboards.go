package gojira

import (
	"context"
	"fmt"
	"net/http"
)

type DashboardsService service

type Dashboard struct {
	Id string `json:"id"`
	IsFavourite bool `json:"isFavourite"`
	Name string `json:"name"`
	Owner UserBean `json:"owner"`
	Popularity int64 `json:"popularity"`
	Rank int32 `json:"rank"`
	Self string `json:"self"`
	SharePermissions []SharePermission `json:"sharePermissions"`
	View string `json:"view"`
}

type PageOfDashboards struct {
	StartAt int32 `json:"startAt"`
	MaxResults int32 `json:"maxResults"`
	Total int32 `json:"total"`
	Prev string `json:"prev"`
	Next string `json:"next"`
	Dashboards []Dashboard `json:"dashboards"`
}

type PageBeanDashboard struct {
	Self string `json:"self"`
	NextPage string `json:"nextPage"`
	MaxResults int32 `json:"maxResults"`
	StartAt int64 `json:"startAt"`
	Total int64 `json:"total"`
	IsLast bool `json:"isLast"`
	Values []Dashboard `json:"values"`
}

type PropertyKey struct {
	Self string `json:"self"`
	Key string `json:"key"`
}

type PropertyKeys struct {
	Keys []PropertyKey `json:"keys"`
}

type EntityProperty struct {
	Key string `json:"key"`
	Value interface{} `json:"value"`
}

type UserBean struct {
	Key string `json:"key"`
	Self string `json:"self"`
	Name string `json:"name"`
	DisplayName string `json:"displayName"`
	Active string `json:"active"`
	AccountId string `json:"accountId"`
	AvatarUrls map[string]string `json:"avatarUrls"`
}

type SharePermission struct {
	Id int64 `json:"id"`
	Type string `json:"type"`
	Project Project `json:"project"`
	Role ProjectRole `json:"role"`
	Group GroupName `json:"groupName"`
}

type GroupName struct {
	Name string `json:"name,omitempty"`
	Self string `json:"self,omitempty"`
}

type GetAllDashboardsQueryParams struct {
	Filter string `url:"filter,omitempty"`
	StartAt int32 `url:"startAt,omitempty"`
	MaxResults int32 `url:"maxResults,omitempty"`
}

type SearchForDashboardsQueryParams struct {
	DashboardName string `url:"dashboardName,omitempty"`
	AccountId string `url:"accountId,omitempty"`
	Owner string `url:"owner,omitempty"`
	GroupName string `url:"groupname,omitempty"`
	ProjectID int64 `url:"projectId,omitempty"`
	OrderBy string `url:"orderBy,omitempty"`
	StartAt int64 `url:"startAt,omitempty"`
	MaxResults int32 `url:"maxResults,omitempty"`
	Expand string `url:"expand,omitempty"`
}

func (s *DashboardsService) Get (ctx context.Context, dashboardID string) (*Dashboard, *http.Response, error) {
	endpoint := fmt.Sprintf("dashboard/%v", dashboardID)
	req, err := s.client.NewRequest("GET", endpoint, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var d *Dashboard
	resp, err := s.client.Do(ctx, req, &d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, err
}

func (s *DashboardsService) List (ctx context.Context, queryParams *GetAllDashboardsQueryParams) (*PageOfDashboards, *http.Response, error) {
	u, err := addQueryParams("dashboard", queryParams)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var pod *PageOfDashboards
	resp, err := s.client.Do(ctx, req, &pod)
	if err != nil {
		return nil, resp, err
	}

	return pod, resp, nil
}

func (s *DashboardsService) SearchForDashboards (ctx context.Context, queryParams *SearchForDashboardsQueryParams) (*PageBeanDashboard, *http.Response, error) {
	u, err := addQueryParams("dashboard/search", queryParams)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var pbd *PageBeanDashboard
	resp, err := s.client.Do(ctx, req, &pbd)
	if err != nil {
		return nil, resp, err
	}

	return pbd, resp, nil
}

func (s *DashboardsService) GetDashboardItemPropertyKeys (ctx context.Context, dashboardId string, itemId string) (*PropertyKeys, *http.Response, error) {
	endpoint := fmt.Sprintf("dashboard/%v/items/%v/properties", dashboardId, itemId)
	req, err := s.client.NewRequest("GET", endpoint, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var pk *PropertyKeys
	resp, err := s.client.Do(ctx, req, &pk)
	if err != nil {
		return nil, resp, err
	}

	return pk, resp, err
}

func (s *DashboardsService) GetDashboardItemProperty (ctx context.Context, dashboardId string, itemId string, propertyKey string) (*EntityProperty, *http.Response, error) {
	endpoint := fmt.Sprintf("dashboard/%v/items/%v/properties/%v", dashboardId, itemId, propertyKey)
	req, err := s.client.NewRequest("GET", endpoint, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var ep *EntityProperty
	resp, err := s.client.Do(ctx, req, &ep)
	if err != nil {
		return nil, resp, err
	}

	return ep, resp, err
}

func (s *DashboardsService) SetDashboardItemProperty (ctx context.Context, dashboardId string, itemId string, propertyKey string, body interface{}) (*interface{}, *http.Response, error) {
	endpoint := fmt.Sprintf("dashboard/%v/items/%v/properties/%v", dashboardId, itemId, propertyKey)
	req, err := s.client.NewRequest("POST", endpoint, nil, body)
	if err != nil {
		return nil, nil, err
	}

	var i *interface{}
	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, err
}

func (s *DashboardsService) DeleteDashboardItemProperty (ctx context.Context, dashboardId string, itemId string, propertyKey string) (*http.Response, error) {
	endpoint := fmt.Sprintf("dashboard/%v/items/%v/properties/%v", dashboardId, itemId, propertyKey)
	req, err := s.client.NewRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
