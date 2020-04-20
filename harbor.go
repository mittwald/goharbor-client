package harbor

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"strconv"
)

type StatusCodeError struct {
	StatusCode   int
	ExpectedCode int
}

func (e *StatusCodeError) Error() string {
	return fmt.Sprintf("unexpected status code: %d, expected: %d", e.StatusCode, e.ExpectedCode)
}

func I64toA(in int64) string {
	return strconv.FormatInt(in, 10)
}

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty" json:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.
	PageSize int `url:"page_size,omitempty" json:"page_size,omitempty"`
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
