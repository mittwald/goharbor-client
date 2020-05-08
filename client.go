package harbor

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/url"
	"strings"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "goharbor-client/" + libraryVersion
)

// Implementations of Client provide methods for managing Harbor instances.
type Client interface {
	// Projects returns an implementation of ProjectClient.
	Projects() ProjectClient

	// Users returns an implementation of UserClient.
	Users() UserClient

	// Repositories returns an implementation of RepositoryClient.
	Repositories() RepositoryClient

	// Registries returns an implementation of RegistryClient.
	Registries() RegistryClient

	// Replications returns an implementation of ReplicationClient.
	Replications() ReplicationClient
}

// RestClient implements the Client interface by communicating with Harbor via REST api.
type RestClient struct {
	// base URL for Harbor API requests
	baseURL *url.URL

	// client specific baseURLSuffix the client will append to the baseURL
	// default "", for RestProjectClient this will be "projects"
	baseURLSuffix string

	// basic auth username
	username string
	// basic auth password
	password string
}

// NewClient returns a new REST based harbor client.
func NewClient(baseURL, username, password string) (Client, error) {
	c := &RestClient{
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
func (c *RestClient) SetBaseURL(urlStr string) error {
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	var err error
	c.baseURL, err = url.Parse(urlStr)
	return err
}

// NewRequest creates an API request.
// A relative URL path can be provided in urlStr,
// in which case it is resolved relative to the base URL of the RestClient.
// Relative URL paths should always be specified without a preceding slash.
func (c *RestClient) NewRequest(method, subPath string) *gorequest.SuperAgent {
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

// CheckResponse flattens errors and checks the response status code.
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

func (c *RestClient) withURLSuffix(suffix string) *RestClient {
	return &RestClient{
		baseURL:       c.baseURL,
		username:      c.username,
		password:      c.password,
		baseURLSuffix: suffix,
	}
}

// Projects returns a REST based project client.
// It satisfies the Client interface.
func (c *RestClient) Projects() ProjectClient {
	return &RestProjectClient{c.withURLSuffix("projects")}
}

// Users returns a REST based user client.
// It satisfies the Client interface.
func (c *RestClient) Users() UserClient {
	return &RestUserClient{c.withURLSuffix("users")}
}

// Repositories returns a REST based repository client.
// It satisfies the Client interface.
func (c *RestClient) Repositories() RepositoryClient {
	return &RestRepositoryClient{c.withURLSuffix("repositories")}
}

// Registries returns a REST based registry client.
// It satisfies the Client interface.
func (c *RestClient) Registries() RegistryClient {
	return &RestRegistryClient{c.withURLSuffix("registries")}
}

// Replications returns a REST based replication client.
// It satisfies the Client interface.
func (c *RestClient) Replications() ReplicationClient {
	return &RestReplicationClient{c.withURLSuffix("replication")}
}
