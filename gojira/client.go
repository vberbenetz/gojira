package gojira

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "atlassian.net/rest/api/3/"
)

// API Token: vygWwlzHVm2yRgMxbnZo2151
// Authorization : Basic dmFsLmJlcmJlbmV0ekBnbWFpbC5jb206dnlnV3dsekhWbTJ5UmdNeGJuWm8yMTUx

type Client struct {
	client *http.Client

	BaseUrl *url.URL

	// Removes requirement for structs for each service
	common service

	// Services for different parts of the API
	ApplicationRole *ApplicationRoleService
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client, domain string) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	var baseUrlStr strings.Builder
	baseUrlStr.WriteString("https://")
	baseUrlStr.WriteString(domain)
	baseUrlStr.WriteString(".")
	baseUrlStr.WriteString(defaultBaseURL)

	baseUrl, err := url.Parse(baseUrlStr.String())
	if err != nil {
		return nil, err
	}

	c := &Client{client: httpClient, BaseUrl: baseUrl}
	c.common.client = c
	c.ApplicationRole = (*ApplicationRoleService)(&c.common)
	return c, nil
}

func (c *Client) Close() error {
	c.client = nil
	return nil
}

func (c *Client) NewRequest (method, urlEndpoint string, body interface{}) (*http.Request, error) {
	u, err := c.BaseUrl.Parse(urlEndpoint)
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

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	return req, nil
}

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

type ResponseError struct {
	Response *http.Response
	ErrorMessages []string `json:"errorMessages"`
	Errors interface{} `json:"errors"`
	StatusCode int32 `json:"status"`
}

func (r *ResponseError) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.ErrorMessages, r.Errors)
}

// A client type which will attach the API key via basic auth to each request
type BasicAuth struct {
	Username string
	ApiKey string

	Transport http.RoundTripper
}

// Adds basic auth to the RoundTripper
func (t *BasicAuth) RoundTrip(req *http.Request) (*http.Response, error) {
	newReq := copyRequest(req)

	newReq.SetBasicAuth(t.Username, t.ApiKey)
	return t.transport().RoundTrip(newReq)
}

// Returns a client object that is already configured with the basic auth.
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