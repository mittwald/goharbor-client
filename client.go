package harbor

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/url"
	"strings"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-harbor/" + libraryVersion
)

type Client struct {
	// base URL for Harbor API requests
	baseURL *url.URL

	// client specific baseURLSuffix the client will append to the baseURL
	// default "", for ProjectClient this will be "projects"
	baseURLSuffix string

	// basic auth username
	username string
	// basic auth password
	password string
}

func NewClient(baseURL, username, password string) (*Client, error) {
	return newClient(baseURL, username, password)
}

func newClient(baseURL, username, password string) (*Client, error) {
	c := &Client{
		username: username,
		password: password,
	}

	if err := c.SetBaseURL(baseURL); err != nil {
		return nil, err
	}

	return c, nil
}

// SetBaseURL sets the base URL for API requests to a custom endpoint. urlStr
// should always be specified with a trailing slash.
func (c *Client) SetBaseURL(urlStr string) error {
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	var err error
	c.baseURL, err = url.Parse(urlStr)
	return err
}

// NewRequest
// creates an API request. A relative URL path can be provided in
// urlStr, in which case it is resolved relative to the base URL of the Client.
// Relative URL paths should always be specified without a preceding slash.
func (c *Client) NewRequest(method, subPath string) *gorequest.SuperAgent {
	r := gorequest.New()
	r.Set("Accept", "application/json")
	r.Set("User-Agent", userAgent)
	r.SetBasicAuth(c.username, c.password)

	r.Method = method

	u := c.baseURL.String() + "api/" + c.baseURLSuffix + subPath
	r.Url = u

	if method == gorequest.PUT || method == gorequest.POST {
		r.Set("Content-Type", "application/json")
	}

	return r
}

// CheckResponse flattens errors and checks the response status code
func CheckResponse(errs []error, resp gorequest.Response, expected int) error {
	var err error
	// flatten the gorequests error into a single one
	// TODO: either replace the underlying library or handle this in a better way
	for _, e := range errs {
		if err == nil {
			err = fmt.Errorf("%v", e)
		}
		err = fmt.Errorf("%v %v", e, err)
	}

	if err != nil {
		return err
	}

	if resp.StatusCode != expected {
		return &StatusCodeError{
			StatusCode:   resp.StatusCode,
			ExpectedCode: expected,
		}
	}
	return nil
}

func (c *Client) Projects() *ProjectClient {
	p := &ProjectClient{Client: &Client{
		baseURL:  c.baseURL,
		username: c.username,
		password: c.password,
	}}
	p.baseURLSuffix = "projects"
	return p
}

func (c *Client) Users() *UserClient {
	u := &UserClient{Client: &Client{
		baseURL:  c.baseURL,
		username: c.username,
		password: c.password,
	}}

	u.baseURLSuffix = "users"
	return u
}
