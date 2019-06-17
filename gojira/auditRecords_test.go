package gojira

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestAuditRecordsService_List(t *testing.T) {
	client, mux, _, destructor := setup()
	defer destructor()

	mux.HandleFunc("/auditing/record", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"offset": "1",
			"limit": "1000",
			"filter": "no-filter",
			"from": "2019-01-01T12:00:00Z",
			"to": "2019-06-15T12:00:00Z",
		})

		w.WriteHeader(http.StatusOK)
		w.Write(auditRecordsJSON)
	})

	queryParams := &AuditRecordsQueryParams{
		Offset: 1,
		Limit: 1000,
		Filter: "no-filter",
		From: time.Date(2019, time.January, 1, 12, 0, 0, 0, time.UTC),
		To: time.Date(2019, time.June, 15, 12, 0, 0, 0, time.UTC),
	}

	received, _, err := client.AuditRecords.List(context.Background(), queryParams)
	if err != nil {
		t.Errorf("AuditRecords.List returned error: %v", err)
	}

	if expected := expectedAuditRecordResponse; !reflect.DeepEqual(received, expected) {
		t.Errorf("AuditRecords.List = %v, expected: %+v", received, expected)
	}
}

var auditRecordsJSON = []byte(`{
   "offset":1,
   "limit":1000,
   "total":1,
   "records":[
      {
         "id":10126,
         "summary":"User added to group",
         "created":"2019-06-09T23:56:43.901+0000",
         "category":"group management",
         "eventSource":"",
         "objectItem":{
            "name":"jira-software-users",
            "typeName":"GROUP",
            "parentId":"10000",
            "parentName":"IDP Directory"
         },
         "associatedItems":[
            {
               "id":"addon_jira-workplace-integration",
               "name":"addon_jira-workplace-integration",
               "typeName":"USER",
               "parentId":"10000",
               "parentName":"IDP Directory"
            }
         ]
      }
   ]
}`)

var expectedAuditRecordResponse = &AuditRecords{
	Offset: 1,
	Limit: 1000,
	Total: 1,
	Records: []AuditRecordBean{
		{
			ID: 10126,
			Summary: "User added to group",
			Created: "2019-06-09T23:56:43.901+0000",
			Category: "group management",
			EventSource: "",
			ObjectItem: AssociatedItemBean{
				Name: "jira-software-users",
				TypeName: "GROUP",
				ParentID: "10000",
				ParentName: "IDP Directory",
			},
			AssociatedItems: []AssociatedItemBean{
				{
					ID: "addon_jira-workplace-integration",
					Name: "addon_jira-workplace-integration",
					TypeName: "USER",
					ParentID: "10000",
					ParentName: "IDP Directory",
				},
			},
		},
	},
}
