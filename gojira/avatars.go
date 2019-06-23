package gojira

import (
	"context"
	"fmt"
	"net/http"
)

// AvatarsService is used to interact with all Avatars related endpoints
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/#api-group-Avatars
type AvatarsService service

// Avatar is the object representing an avatar object
type Avatar struct {
	ID string `json:"id"`
	Owner string `json:"owner,omitempty"`
	IsSystemAvatar bool `json:"isSystemAvatar,omitempty"`
	IsSelected bool `json:"isSelected,omitempty"`
	IsDeletable bool `json:"isDeletable,omitempty"`
	FileName string `json:"fileName,omitempty"`
	Urls interface{} `json:"urls,omitempty"`
}

// Avatars is the the response object encompassing both System and Custom avatars
type Avatars struct {
	System []Avatar `json:"system,omitempty"`
	Custom []Avatar `json:"custom,omitempty"`
}

// SystemAvatars is the object returned when getting a SystemAvatar by type
type SystemAvatars struct {
	System []Avatar `json:"system"`
}

// LoadAvatarQueryParams are the query parameters for the Load Avatar request
// Parameters [X, Y] are the coordinates of the top-left corner of the crop region. Default (0,0)
// Size is the length of each side of the crop region
type LoadAvatarQueryParams struct {
	X int32 `json:"x,omitempty"`
	Y int32 `json:"y,omitempty"`
	Size int32 `json:"size,omitempty"`
}

// Get returns the Avatars (system & custom) based on the given type of entity and ID or entity item.
// Accepted values for entityType are: [ "project", "issueType" ]
func (s *AvatarsService) Get (ctx context.Context, entityType string, entityID string) (*Avatars, *http.Response, error) {

	endpoint := fmt.Sprintf("universal_avatar/type/%v/owner/%v", entityType, entityID)
	req, err := s.client.NewRequest("GET", endpoint, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var a *Avatars
	resp, err := s.client.Do(ctx, req, &a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, err
}

// GetSystemAvatarsByType retrieve a list of system avatars by the provided ownerType.
// Accepted values for ownerType are: [ "issuetype", "project", "user" ]
func (s *AvatarsService) GetSystemAvatarsByType (ctx context.Context, ownerType string) (*SystemAvatars, *http.Response, error) {

	endpoint := fmt.Sprintf("avatar/%v/system", ownerType)
	req, err := s.client.NewRequest("GET", endpoint, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var sa *SystemAvatars
	resp, err := s.client.Do(ctx, req, &sa)
	if err != nil {
		return nil, resp, err
	}

	return sa, resp, err
}

// LoadAvatar sends a user provided image to Jira to use as their avatar.
// Accepted values for entityType are: [ "project", "issueType" ]
// Accepted values for imageType are: [ "JPEG", "GIF", "PNG" ]
// The image needs to be converted to a byte array and passed in via imageData
func (s *AvatarsService) LoadAvatar (ctx context.Context, entityType string, entityID string, imageType string, imageData *[]byte, queryParams *LoadAvatarQueryParams) (*Avatar, *http.Response, error) {

	endpoint := fmt.Sprintf("universal_avatar/type/%v/owner/%v", entityType, entityID)

	u, err := addQueryParams(endpoint, queryParams)
	if err != nil {
		return nil, nil, err
	}

	reqHeaders := make(map[string]string)

	reqHeaders["Content-Type"] = fmt.Sprintf("image/%v", imageType)
	reqHeaders["X-Atlassian-Token"] = "no-check"

	req, err := s.client.NewRequest("POST", u, reqHeaders, imageData)
	if err != nil {
		return nil, nil, err
	}

	var a *Avatar
	resp, err := s.client.Do(ctx, req, &a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, err
}

// Delete removes an avatar from a project or issue type
// Accepted values for entityType are: [ "project", "issueType" ]
// OwningObjectID is the ID of the entity item
func (s *AvatarsService) Delete (ctx context.Context, entityType string, owningObjectID string, avatarID int64) (*http.Response, error) {
	endpoint := fmt.Sprintf("universal_avatar/type/%v/owner/%v/avatar/%v", entityType, owningObjectID, avatarID)
	req, err := s.client.NewRequest("GET", endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}