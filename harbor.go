package harbor

import (
	"github.com/parnurzeal/gorequest"
	"net/http"
	"net/url"
	"strings"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-harbor/" + libraryVersion
)

type Client struct {
	// HTTP client used to communicate with the Harbor API
	client *gorequest.SuperAgent
	// Base URL for Harbor API requests
	baseURL *url.URL
	// User agent used when communicating with the Harbor API
	UserAgent string
	// XSRFKey used when communicating with the Harbor API
	XSRFKey string
	// Services used for talking to different parts of the Harbor API
	Projects     *ProjectClient
	Repositories *RepositoryClient
	Users        *UserClient
}

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty" json:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.
	PageSize int `url:"page_size,omitempty" json:"page_size,omitempty"`
}

func NewClient(harborClient *gorequest.SuperAgent, baseURL, username, password, xsrfKey string) *Client {
	return newClient(harborClient, baseURL, username, password, xsrfKey)
}

// SetBaseURL sets the base URL for API requests to a custom endpoint. urlStr
// should always be specified with a trailing slash.
func (c *Client) CheckBaseURL(urlStr string) error {
	// Make sure the given URL end with a slash
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}
	var err error
	c.baseURL, err = url.Parse(urlStr)
	return err
}

// SetXsrfKey sets the XSRF Key for API requests
func (c *Client) CheckXsrfKey(xsrfKey string) error {
	var err error
	if len(xsrfKey) > 0 {
		return nil
	}

	return err
}

func newClient(harborClient *gorequest.SuperAgent, baseURL, username, password, xsrfKey string) *Client {
	if harborClient == nil {
		harborClient = gorequest.New()
	}
	harborClient.SetBasicAuth(username, password)

	c := &Client{client: harborClient, UserAgent: userAgent}
	if err := c.CheckBaseURL(baseURL); err != nil {
		// Should never happen since defaultBaseURL is our constant.
		panic(err)
	}

	if err := c.CheckXsrfKey(xsrfKey); err != nil {
		panic(err)
	}
	c.XSRFKey = xsrfKey

	// Create all the public services.
	c.Projects = &ProjectClient{client: c}
	c.Repositories = &RepositoryClient{client: c}
	c.Users = &UserClient{client: c}
	return c
}

// NewRequest
// creates an API request. A relative URL path can be provided in
// urlStr, in which case it is resolved relative to the base URL of the Client.
// Relative URL paths should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, subPath string) *gorequest.SuperAgent {
	u := c.baseURL.String() + "api/" + subPath
	h := c.client.Set("Accept", "application/json")

	xsrf := http.Cookie{Name: "_xsrf", Value: c.XSRFKey}

	if c.XSRFKey != "" {
		h.Set("X-Xsrftoken", c.XSRFKey)
	}

	if c.UserAgent != "" {
		h.Set("User-Agent", c.UserAgent)
	}
	switch method {
	case gorequest.PUT:
		return c.client.Put(u).Set("Content-Type", "application/json").Set("X-Xsrftoken", c.XSRFKey).AddCookie(&xsrf)
	case gorequest.POST:
		return c.client.Post(u).Set("Content-Type", "application/json").Set("X-Xsrftoken", c.XSRFKey).AddCookie(&xsrf)
	case gorequest.GET:
		return c.client.Get(u)
	case gorequest.HEAD:
		return c.client.Head(u)
	case gorequest.DELETE:
		return c.client.Delete(u)
	case gorequest.PATCH:
		return c.client.Patch(u)
	case gorequest.OPTIONS:
		return c.client.Options(u)
	default:
		return c.client.Get(u)
	}
}

type SearchRepository struct {
	// The ID of the project that the repository belongs to
	ProjectId int32 `json:"project_id,omitempty"`
	// The name of the project that the repository belongs to
	ProjectName string `json:"project_name,omitempty"`
	// The flag to indicate the publicity of the project that the repository belongs to
	ProjectPublic bool `json:"project_public,omitempty"`
	// The name of the repository
	RepositoryName string `json:"repository_name,omitempty"`
	PullCount      int32  `json:"pull_count,omitempty"`
	TagsCount      int32  `json:"tags_count,omitempty"`
}

type Search struct {
	// Search results of the projects that matched the filter keywords.
	Projects Project `json:"project,omitempty"`
	// Search results of the repositories that matched the filter keywords.
	Repositories []SearchRepository `json:"repository,omitempty"`
}

// Search for projects and repositories
//
// The Search endpoint returns information about the projects and repositories
// offered at public status or related to the current logged in user. The
// response includes the project and repository list in a proper
// display order.
//
// Harbor API docs: https://github.com/vmware/harbor/blob/release-1.4.0/docs/swagger.yaml#L17
func (c *Client) Search() (Search, *gorequest.Response, []error) {
	var search Search
	resp, _, errs := c.NewRequest(gorequest.GET, "search").
		EndStruct(&search)
	return search, &resp, errs
}

type StatisticMap struct {
	// The count of the private projects which the user is a member of.
	PrivateProjectCount int `json:"private_project_count,omitempty"`
	// The count of the private repositories belonging to the projects which the user is a member of.
	PrivateRepoCount int `json:"private_repo_count,omitempty"`
	// The count of the public projects.
	PublicProjectCount int `json:"public_project_count,omitempty"`
	// The count of the public repositories belonging to the public projects which the user is a member of.
	PublicRepoCount int `json:"public_repo_count,omitempty"`
	// The count of the total projects, only be seen when the is admin.
	TotalProjectCount int `json:"total_project_count,omitempty"`
	// The count of the total repositories, only be seen when the user is admin.
	TotalRepoCount int `json:"total_repo_count,omitempty"`
}

// GetStatistics
// Get project and repository statistics
func (c *Client) GetStatistics() (StatisticMap, *gorequest.Response, []error) {
	var statistics StatisticMap
	resp, _, errs := c.NewRequest(gorequest.GET, "statistics").
		EndStruct(&statistics)
	return statistics, &resp, errs
}
