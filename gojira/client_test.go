package gojira

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	// Base path used during testing
	baseURLPath = "/rest-api-v3"
)

// Setup a test HTTP server with a jira client which will interact with the test server.
// The mux will send mock responses for the API endpoints being tested.
func setup() (client *Client, mux *http.ServeMux, serverURL string, destructor func()) {

	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath + "/", http.StripPrefix(baseURLPath, mux))
	apiHandler.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("FAIL: Client.BaseURL path prefix is not in the request URL:")
		fmt.Println(req.URL.String())
		fmt.Println("Use a relative endpoint URL")
		http.Error(w, "BaseURL path prefix: " + baseURLPath + " not in request URL", http.StatusInternalServerError)
	})

	// Test HTTP server used to provide mock API responses
	mockServer := httptest.NewServer(apiHandler)

	// Jira client being tested against the mock server
	client, _ = NewClient(nil, mockServer.URL)
	mockServerURL, _ := url.Parse(mockServer.URL + baseURLPath + "/")
	client.BaseURL = mockServerURL

	return client, mux, mockServer.URL, mockServer.Close
}

func testMethod(t *testing.T, r *http.Request, method string) {
	if reqMethod := r.Method; reqMethod != method {
		t.Errorf("Request method: %v, expected %v", reqMethod, method)
	}
}

// TODO: TestNewRequest

// TODO: TestNewRequest Invalid JSON

// TODO: TestNewRequest Bad URL

// TODO: TestNewRequest Trailing Backslash

// TODO: TestDo

// TODO: TestDo HttpError


