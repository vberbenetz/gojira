package gojira

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestAvatarsService_Get(t *testing.T) {
	client, mux, _, destructor := setup()
	defer destructor()

	mux.HandleFunc("/universal_avatar/type/project/owner/1000", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(getAvatarsResponseJSON)
	})

	// Test successful key match and returned value
	received, _, err := client.AvatarsService.Get(context.Background(), "project", "1000")
	if err != nil {
		t.Errorf("AvatarsService.Get returned error: %v", err)
	}
	if expected := expectedAvatarsResponse; !reflect.DeepEqual(received, expected) {
		t.Errorf("AvatarsService.Get = %v,\n EXPECTED: %v", received, expected)
	}
}

func TestAvatarsService_GetSystemAvatarsByType(t *testing.T) {
	client, mux, _, destructor := setup()
	defer destructor()

	mux.HandleFunc("/avatar/project/system", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(getSystemAvatarsResponseJSON)
	})

	// Test successful key match and returned value
	received, _, err := client.AvatarsService.GetSystemAvatarsByType(context.Background(), "project")
	if err != nil {
		t.Errorf("AvatarsService.GetSystemAvatarsByType returned error: %v", err)
	}
	if expected := expectedSystemAvatarsResponse; !reflect.DeepEqual(received, expected) {
		t.Errorf("AvatarsService.GetSystemAvatarsByType = %v,\n EXPECTED: %v", received, expected)
	}
}

func TestAvatarsService_LoadAvatar(t *testing.T) {
	client, mux, _, destructor := setup()
	defer destructor()

	mux.HandleFunc("/universal_avatar/type/project/owner/1000", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		w.WriteHeader(http.StatusOK)
		w.Write(loadAvatarResponseJSON)
	})

	// Test successful key match and returned value
	received, _, err := client.AvatarsService.LoadAvatar(context.Background(), "project", "1000", "JPEG", nil, nil)
	if err != nil {
		t.Errorf("AvatarsService.LoadAvatar returned error: %v", err)
	}
	if expected := avatarBody1; !reflect.DeepEqual(received, expected) {
		t.Errorf("AvatarsService.LoadAvatar = %v,\n EXPECTED: %v", received, expected)
	}
}

func TestAvatarsService_Delete(t *testing.T) {
	client, mux, _, destructor := setup()
	defer destructor()

	mux.HandleFunc("/universal_avatar/type/project/owner/20/avatar/1000", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

		w.WriteHeader(http.StatusOK)
		w.Write(loadAvatarResponseJSON)
	})

	// Test successful key match and returned value
	_, err := client.AvatarsService.Delete(context.Background(), "project", "20", 1000)
	if err != nil {
		t.Errorf("AvatarsService.Delete returned error: %v", err)
	}
}

var getAvatarsResponseJSON = []byte(`
{
   "system":[
      {
         "id":"1000",
         "isSystemAvatar":true,
         "isSelected":false,
         "isDeletable":false,
         "urls":{
            "16x16":"https://your-domain.atlassian.net/secure/viewavatar?size=xsmall&avatarId=10040&avatarType=project",
            "24x24":"https://your-domain.atlassian.net/secure/viewavatar?size=small&avatarId=10040&avatarType=project",
            "32x32":"https://your-domain.atlassian.net/secure/viewavatar?size=medium&avatarId=10040&avatarType=project",
            "48x48":"https://your-domain.atlassian.net/secure/viewavatar?avatarId=10040&avatarType=project"
         }
      }
   ],
   "custom":[
      {
         "id":"1010",
         "isSystemAvatar":false,
         "isSelected":false,
         "isDeletable":true,
         "urls":{
            "16x16":"https://your-domain.atlassian.net/secure/viewavatar?size=xsmall&avatarId=10080&avatarType=project",
            "24x24":"https://your-domain.atlassian.net/secure/viewavatar?size=small&avatarId=10080&avatarType=project",
            "32x32":"https://your-domain.atlassian.net/secure/viewavatar?size=medium&avatarId=10080&avatarType=project",
            "48x48":"https://your-domain.atlassian.net/secure/viewavatar?avatarId=10080&avatarType=project"
         }
      }
   ]
}`)

var getSystemAvatarsResponseJSON = []byte(`
{
   "system":[
      {
         "id":"1000",
         "isSystemAvatar":true,
         "isSelected":false,
         "isDeletable":false,
         "urls":{
            "16x16":"https://your-domain.atlassian.net/secure/viewavatar?size=xsmall&avatarId=10040&avatarType=project",
            "24x24":"https://your-domain.atlassian.net/secure/viewavatar?size=small&avatarId=10040&avatarType=project",
            "32x32":"https://your-domain.atlassian.net/secure/viewavatar?size=medium&avatarId=10040&avatarType=project",
            "48x48":"https://your-domain.atlassian.net/secure/viewavatar?avatarId=10040&avatarType=project"
         }
      }
   ]
}`)

var loadAvatarResponseJSON = []byte(`
{
	"id":"1000",
    "isSystemAvatar":true,
    "isSelected":false,
    "isDeletable":false,
    "urls":{
       "16x16":"https://your-domain.atlassian.net/secure/viewavatar?size=xsmall&avatarId=10040&avatarType=project",
       "24x24":"https://your-domain.atlassian.net/secure/viewavatar?size=small&avatarId=10040&avatarType=project",
       "32x32":"https://your-domain.atlassian.net/secure/viewavatar?size=medium&avatarId=10040&avatarType=project",
       "48x48":"https://your-domain.atlassian.net/secure/viewavatar?avatarId=10040&avatarType=project"
	}
}`)

var expectedAvatarsResponse = &Avatars{
	System: []Avatar{
		*avatarBody1,
	},
	Custom: []Avatar{
		*avatarBody2,
	},
}

var expectedSystemAvatarsResponse = &SystemAvatars{
	System: []Avatar{
		*avatarBody1,
	},
}

var avatarBody1 = &Avatar{
	ID: "1000",
	IsSystemAvatar: true,
	IsSelected: false,
	IsDeletable: false,
	Urls: map[string]string{
		"16x16": "https://your-domain.atlassian.net/secure/viewavatar?size=xsmall&avatarId=10040&avatarType=project",
		"24x24": "https://your-domain.atlassian.net/secure/viewavatar?size=small&avatarId=10040&avatarType=project",
		"32x32": "https://your-domain.atlassian.net/secure/viewavatar?size=medium&avatarId=10040&avatarType=project",
		"48x48": "https://your-domain.atlassian.net/secure/viewavatar?avatarId=10040&avatarType=project",
	},
}

var avatarBody2 = &Avatar{
	ID: "1010",
	IsSystemAvatar: false,
	IsSelected: false,
	IsDeletable: true,
	Urls: map[string]string{
		"16x16":"https://your-domain.atlassian.net/secure/viewavatar?size=xsmall&avatarId=10080&avatarType=project",
		"24x24":"https://your-domain.atlassian.net/secure/viewavatar?size=small&avatarId=10080&avatarType=project",
		"32x32":"https://your-domain.atlassian.net/secure/viewavatar?size=medium&avatarId=10080&avatarType=project",
		"48x48":"https://your-domain.atlassian.net/secure/viewavatar?avatarId=10080&avatarType=project",
	},
}
