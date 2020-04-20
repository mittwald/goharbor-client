package harbor

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type cfg struct {
	Labels map[string]string `json:"labels"`
}

type Signature struct {
	Tag    string            `json:"tag"`
	Hashes map[string][]byte `json:"hashes"`
}

type ManifestResp struct {
	Manifest interface{} `json:"manifest"`
	Config   interface{} `json:"config,omitempty" `
}

// ComponentsOverview holds information about the total number of components with a certain CVE severity
type ComponentsOverview struct {
	Total   int                        `json:"total"`
	Summary []*ComponentsOverviewEntry `json:"summary"`
}

//ComponentsOverviewEntry ...
type ComponentsOverviewEntry struct {
	Sev   int `json:"severity"`
	Count int `json:"count"`
}

// VulnerabilityItem is an item in the vulnerability result returned by the vulnerability details API
type VulnerabilityItem struct {
	ID          string `json:"id"`
	Severity    int64  `json:"severity"`
	Pkg         string `json:"package"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Fixed       string `json:"fixedVersion,omitempty"`
}

type RepoResp struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	ProjectID   int64  `json:"project_id"`
	Description string `json:"description"`
	PullCount   int64  `json:"pull_count"`
	StarCount   int64  `json:"star_count"`
	TagsCount   int64  `json:"tags_count"`
}

// RepoRecord holds the record of an repository, held by the database
type RepoRecord struct {
	RepositoryID int64  `json:"repository_id"`
	Name         string `json:"name"`
	ProjectID    int64  `json:"project_id"`
	Description  string `json:"description"`
	PullCount    int64  `json:"pull_count"`
	StarCount    int64  `json:"star_count"`
}

// ImgScanOverview maps the record of an image scan overview.
type ImgScanOverview struct {
	ID              int64               `json:"-"`
	Digest          string              `json:"image_digest"`
	Status          string              `json:"scan_status"`
	JobID           int64               `json:"job_id"`
	Sev             int                 `json:"severity"`
	CompOverviewStr string              `json:"-"`
	CompOverview    *ComponentsOverview `json:"components,omitempty"`
	DetailsKey      string              `json:"details_key"`
}

type tagDetail struct {
	Digest        string `json:"digest"`
	Name          string `json:"name"`
	Size          int64  `json:"size"`
	Architecture  string `json:"architecture"`
	OS            string `json:"os"`
	DockerVersion string `json:"docker_version"`
	Author        string `json:"author"`
	Config        *cfg   `json:"config"`
}

type TagResp struct {
	tagDetail
	Signature    *Signature       `json:"signature"`
	ScanOverview *ImgScanOverview `json:"scan_overview,omitempty"`
}

// RepositoryClient handles communication with the repository related methods of the Harbor API
type RepositoryClient struct {
	client *Client
}

type ListRepositoriesOption struct {
	ListOptions
	ProjectId int64  `url:"project_id,omitempty" json:"project_id,omitempty"`
	Q         string `url:"q,omitempty" json:"q,omitempty"`
	Sort      string `url:"sort,omitempty" json:"sort,omitempty"`
}

// ListRepository
// Get repositories filtered by the relevant project ID and repository name
func (s *RepositoryClient) ListRepository(opt *ListRepositoriesOption) ([]RepoRecord, gorequest.Response, []error) {
	var v []RepoRecord
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, "repositories").
		Query(*opt).
		EndStruct(&v)
	return v, resp, errs
}

// DeleteRepository
// Delete a repository
func (s *RepositoryClient) DeleteRepository(repoName string) (gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.DELETE, fmt.Sprintf("repositories/%s", repoName)).
		End()
	return resp, errs
}

type RepositoryDescription struct {
	Description string `url:"description,omitempty" json:"description,omitempty"`
}

// UpdateRepository
// Update the description of a repository
func (s *RepositoryClient) UpdateRepository(repoName string, d RepositoryDescription) (gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.PUT, fmt.Sprintf("repositories/%s", repoName)).
		Send(d).
		End()
	return resp, errs
}

// GetRepositoryTag
// Get the tag of a repository
// NOTE: If deployed with Notary, the signature property of response represents whether the image is signed or not
// If the property is null, the image is unsigned
func (s *RepositoryClient) GetRepositoryTag(repoName, tag string) (TagResp, gorequest.Response, []error) {
	var v TagResp
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("repositories/%s/tags/%s", repoName, tag)).
		EndStruct(&v)
	return v, resp, errs
}

// DeleteRepositoryTag
// Delete tags of a repository
func (s *RepositoryClient) DeleteRepositoryTag(repoName, tag string) (gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.DELETE, fmt.Sprintf("repositories/%s/tags/%s", repoName, tag)).
		End()
	return resp, errs
}

// ListRepositoryTags
// Get tags from a repository
// NOTE: If deployed with Notary, the signature property of response represents whether the image is signed or not
// If the property is null, the image is unsigned

func (s *RepositoryClient) ListRepositoryTags(repoName string) ([]TagResp, gorequest.Response, []error) {
	var v []TagResp
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("repositories/%s/tags", repoName)).
		EndStruct(&v)
	return v, resp, errs
}

// GetRepositoryTagManifests
// Get manifests from a relevant repository
func (s *RepositoryClient) GetRepositoryTagManifests(repoName, tag string, version string) (ManifestResp, gorequest.Response, []error) {
	var v ManifestResp
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, func() string {
			if version == "" {
				return fmt.Sprintf("repositories/%s/tags/%s/manifest", repoName, tag)
			}
			return fmt.Sprintf("repositories/%s/tags/%s/manifest?version=%s", repoName, tag, version)
		}()).
		EndStruct(&v)
	return v, resp, errs
}

// ScanImage
// Trigger the jobservice component to call the Clair API to scan the image
// Only accessible for project admins
func (s *RepositoryClient) ScanImage(repoName, tag string) (gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.POST, fmt.Sprintf("repositories/%s/tags/%s/scan", repoName, tag)).
		End()
	return resp, errs
}

// GetImageDetails
// Get information from the Clair API containing vulnerability information based on the previous successful scan
func (s *RepositoryClient) GetImageDetails(repoName, tag string) ([]VulnerabilityItem, gorequest.Response, []error) {
	var v []VulnerabilityItem
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("repositories/%s/tags/%s/vulnerability/details", repoName, tag)).
		EndStruct(&v)
	return v, resp, errs
}

// GetRepositorySignature
// Get signature information of a repository, originating from the notary component of Harbor
// NOTE: If the repository does not have any signature information in notary, this API will
// return an empty list with response code 200, instead of 404
func (s *RepositoryClient) GetRepositorySignature(repoName string) ([]Signature, gorequest.Response, []error) {
	var v []Signature
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("repositories/%s/signatures", repoName)).
		EndStruct(&v)
	return v, resp, errs
}

// GetRepositoryTop
// Get the most popular public repositories
func (s *RepositoryClient) GetRepositoryTop(top interface{}) ([]RepoResp, gorequest.Response, []error) {
	var v []RepoResp
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, func() string {
			if t, ok := top.(int); ok {
				return fmt.Sprintf("repositories/top?count=%d", t)
			}
			return fmt.Sprintf("repositories/top")
		}()).
		EndStruct(&v)
	return v, resp, errs
}
