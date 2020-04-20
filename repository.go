package harbor

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

// RepositoryClient handles communication with the repository related methods of the Harbor API
type RepositoryClient struct {
	client *Client
}

// ListRepository
// Get repositories filtered by the relevant project ID and repository name
func (s *RepositoryClient) ListRepository(opt *RepositoryQuery) ([]RepoRecord, gorequest.Response, []error) {
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
func (s *RepositoryClient) GetRepositoryTop(top interface{}) ([]RepoRecord, gorequest.Response, []error) {
	var v []RepoRecord
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
