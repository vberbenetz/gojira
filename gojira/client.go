package gojira

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "atlassian.net/rest/api/3/"
)

// Client is used to access all the services offered by the Jira API
type Client struct {
	client *http.Client

	BaseURL *url.URL

	// Removes requirement for structs for each service
	common service

	// Services for different parts of the API
	ApplicationRole *ApplicationRoleService
	AuditRecords *AuditRecordsService
	AvatarsService *AvatarsService
}

type service struct {
	client *Client
}

// NewClient creates a new client with the project's Atlassian subdomain
func NewClient(httpClient *http.Client, atlasSubdomain string) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	var baseURLStr strings.Builder
	baseURLStr.WriteString("https://")
	baseURLStr.WriteString(atlasSubdomain)
	baseURLStr.WriteString(".")
	baseURLStr.WriteString(defaultBaseURL)

	baseURL, err := url.Parse(baseURLStr.String())
	if err != nil {
		return nil, err
	}

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.common.client = c
	c.ApplicationRole = (*ApplicationRoleService)(&c.common)
	c.AuditRecords = (*AuditRecordsService)(&c.common)
	c.AvatarsService = (*AvatarsService)(&c.common)
	return c, nil
}

// Close is the cleanup function used to remove association with the client
func (c *Client) Close() error {
	c.client = nil
	return nil
}

// NewRequest creates a new request to the given Jira endpoint.
// It formats and encodes the JSON body for the given request.
func (c *Client) NewRequest (method, urlEndpoint string, headers map[string]string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlEndpoint)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		jsonEncoder := json.NewEncoder(buf)
		jsonEncoder.SetEscapeHTML(false)
		err := jsonEncoder.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	// Set content-type default to application/json if no content-type provided.
	if _, ok := headers["Content-Type"]; !ok {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")

	return req, nil
}

// Do function executes the HTTP request.
// It handles any errors that are returned in the response.
// The response body is also decoded into an object and appended to the response.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)

	if err != nil {
		select {

		// Context error takes precedence over other errors
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			return nil, err
		}
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	err = CheckResponseError(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decodeErr := json.NewDecoder(resp.Body).Decode(v)

			// Ignore EOF errors
			if decodeErr != io.EOF {
				decodeErr = nil
			}

			if decodeErr != nil {
				err = decodeErr
			}
		}
	}

	return resp, err
}

// CheckResponseError logs and formats a response error when an HTTP status code not in the 200s is returned
func CheckResponseError(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	responseError := &ResponseError{ Response: r }
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		err := json.Unmarshal(data, responseError)
		if err != nil {
			log.Println(err)
		}
	}

	return responseError
}

// ResponseError is used to model the formatted error returned by the response.
type ResponseError struct {
	Response *http.Response
	ErrorMessages []string `json:"errorMessages"`
	Errors interface{} `json:"errors"`
	StatusCode int32 `json:"status"`
}

// Implements Error function in order for ResponseError to be an error type.
func (r *ResponseError) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.ErrorMessages, r.Errors)
}

// BasicAuth is used with a client type which will attach the API key via basic auth to each request
type BasicAuth struct {
	Username string
	APIKey string

	Transport http.RoundTripper
}

// RoundTrip adds basic auth to the RoundTripper
func (t *BasicAuth) RoundTrip(req *http.Request) (*http.Response, error) {
	newReq := copyRequest(req)

	newReq.SetBasicAuth(t.Username, t.APIKey)
	return t.transport().RoundTrip(newReq)
}

// Client returns a client object that is already configured with the basic auth.
// Removes the need for attaching the basic auth header in each request.
func (t *BasicAuth) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *BasicAuth) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func copyRequest(r *http.Request) *http.Request {
	r2 := &http.Request{}
	*r2 = *r

	r2.Header = make(http.Header, len(r.Header))
	for k, v := range r.Header {
		r2.Header[k] = append([]string(nil), v...)
	}

	return r2
}

// addQueryparams adds the values passed in query parameter interface to s (the url string)
func addQueryParams(s string, queryParams interface{}) (string, error) {
	v := reflect.ValueOf(queryParams)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(queryParams)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}