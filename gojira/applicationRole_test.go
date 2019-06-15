package gojira

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestApplicationRoleService_List(t *testing.T) {
	client, mux, _, destructor := setup()
	defer destructor()

	mux.HandleFunc("/applicationrole", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(appRolesJson)
	})

	received, _, err := client.ApplicationRole.List(context.Background())
	if err != nil {
		t.Errorf("ApplicationRole.List returned error: %v", err)
	}
	if expected := expectedAppRoles; !reflect.DeepEqual(received, expected) {
		t.Errorf("ApplicationRole.List = %v, expected %+v", received, expected)
	}
}

var appRolesJson = []byte(`[
   {
      "key":"jira-software",
      "groups":[
         "jira-software-users",
         "jira-testers"
      ],
      "name":"Jira Software",
      "defaultGroups":[
         "jira-software-users"
      ],
      "selectedByDefault":false,
      "defined":false,
      "numberOfSeats":10,
      "remainingSeats":5,
      "userCount":5,
      "userCountDescription":"5 developers",
      "hasUnlimitedSeats":false,
      "platform":false
   },
   {
      "key":"jira-core",
      "groups":[
         "jira-core-users"
      ],
      "name":"Jira Core",
      "defaultGroups":[
         "jira-core-users"
      ],
      "selectedByDefault":false,
      "defined":false,
      "numberOfSeats":1,
      "remainingSeats":1,
      "userCount":0,
      "userCountDescription":"0 users",
      "hasUnlimitedSeats":false,
      "platform":true
   }
]`)

var expectedAppRoles = []*ApplicationRole{
	expectedJiraSoftwareRole,
	expectedJiraCoreRole,
}

var expectedJiraSoftwareRole = &ApplicationRole{
	Key: "jira-software",
	Groups: []string{
		"jira-software-users",
		"jira-testers",
	},
	Name: "Jira Software",
	DefaultGroups: []string{
		"jira-software-users",
	},
	SelectedByDefault: false,
	Defined: false,
	NumberOfSeats: 10,
	RemainingSeats: 5,
	UserCount: 5,
	UserCountDescription: "5 developers",
	HasUnlimitedSeats: false,
	Platform: false,
}

var expectedJiraCoreRole = &ApplicationRole{
	Key: "jira-core",
	Groups: []string{
		"jira-core-users",
	},
	Name: "Jira Core",
	DefaultGroups: []string{
		"jira-core-users",
	},
	SelectedByDefault: false,
	Defined: false,
	NumberOfSeats: 1,
	RemainingSeats: 1,
	UserCount: 0,
	UserCountDescription: "0 users",
	HasUnlimitedSeats: false,
	Platform: true,
}
