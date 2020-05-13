package harbor

import (
	"fmt"
	"net/url"

	"github.com/parnurzeal/gorequest"
)

// RepositoryClient handles communication with the repository related methods of the Harbor API
type RepositoryClient struct {
	*Client
}

// ListRepository
// Get repositories filtered by the relevant project ID and repository name
func (s *RepositoryClient) ListRepository(opt *RepositoryQuery) ([]RepoRecord, error) {
	var v []RepoRecord
	resp, _, errs := s.NewRequest(gorequest.GET, "").
		Query(*opt).
		EndStruct(&v)
	return v, CheckResponse(errs, resp, 200)
}

// DeleteRepository
// Delete a repository
func (s *RepositoryClient) DeleteRepository(repoName string) error {
	resp, _, errs := s.NewRequest(gorequest.DELETE, "/"+url.PathEscape(repoName)).
		End()
	return CheckResponse(errs, resp, 200)
}

// UpdateRepository
// Update the description of a repository
func (s *RepositoryClient) UpdateRepository(repoName string, d RepositoryDescription) error {
	resp, _, errs := s.NewRequest(gorequest.PUT, "/"+url.PathEscape(repoName)).
		Send(d).
		End()
	return CheckResponse(errs, resp, 200)
}

// GetRepositoryTag
// Get the tag of a repository
// NOTE: If deployed with Notary, the signature property of response represents whether the image is signed or not
// If the property is null, the image is unsigned
func (s *RepositoryClient) GetRepositoryTag(repoName, tag string) (TagResp, error) {
	var v TagResp
	resp, _, errs := s.NewRequest(gorequest.GET,
		fmt.Sprintf("/%s/tags/%s", url.PathEscape(repoName), url.PathEscape(tag))).
		EndStruct(&v)
	return v, CheckResponse(errs, resp, 200)
}

// DeleteRepositoryTag
// Delete tags of a repository
func (s *RepositoryClient) DeleteRepositoryTag(repoName, tag string) error {
	resp, _, errs := s.NewRequest(gorequest.DELETE, fmt.Sprintf("/%s/tags/%s",
		url.PathEscape(repoName), url.PathEscape(tag))).
		End()
	return CheckResponse(errs, resp, 200)
}

// ListRepositoryTags
// Get tags from a repository
// NOTE: If deployed with Notary, the signature property of response represents whether the image is signed or not
// If the property is null, the image is unsigned

func (s *RepositoryClient) ListRepositoryTags(repoName string) ([]TagResp, error) {
	var v []TagResp
	resp, _, errs := s.NewRequest(gorequest.GET,
		fmt.Sprintf("/%s/tags", url.PathEscape(repoName))).
		EndStruct(&v)
	return v, CheckResponse(errs, resp, 200)
}

// GetRepositoryTagManifests
// Get manifests from a relevant repository
func (s *RepositoryClient) GetRepositoryTagManifests(repoName, tag string, version string) (ManifestResp, error) {
	var v ManifestResp
	resp, _, errs := s.NewRequest(gorequest.GET, func() string {
		if version == "" {
			return fmt.Sprintf("/%s/tags/%s/manifest",
				url.PathEscape(repoName), url.PathEscape(tag))
		}
		return fmt.Sprintf("/%s/tags/%s/manifest?version=%s",
			url.PathEscape(repoName), url.PathEscape(tag), url.PathEscape(version))
	}()).
		EndStruct(&v)
	return v, CheckResponse(errs, resp, 200)
}

// ScanImage
// Trigger the jobservice component to call the Clair API to scan the image
// Only accessible for project admins
func (s *RepositoryClient) ScanImage(repoName, tag string) error {
	resp, _, errs := s.NewRequest(gorequest.POST, fmt.Sprintf("/%s/tags/%s/scan",
		url.PathEscape(repoName), url.PathEscape(tag))).
		End()
	return CheckResponse(errs, resp, 202)
}

// GetImageDetails
// Get information from the Clair API containing vulnerability information based on the previous successful scan
func (s *RepositoryClient) GetImageScan(repoName, tag string) ([]VulnerabilityItem, error) {
	var v []VulnerabilityItem
	resp, _, errs := s.NewRequest(gorequest.GET, fmt.Sprintf("/%s/tags/%s/scan",
		url.PathEscape(repoName), url.PathEscape(tag))).
		EndStruct(&v)
	return v, CheckResponse(errs, resp, 200)
}

// GetRepositorySignature
// Get signature information of a repository, originating from the notary component of Harbor
// NOTE: If the repository does not have any signature information in notary, this API will
// return an empty list with response code 200, instead of 404
func (s *RepositoryClient) GetRepositorySignature(repoName string) ([]Signature, error) {
	var v []Signature
	resp, _, errs := s.NewRequest(gorequest.GET,
		fmt.Sprintf("/%s/signatures", url.PathEscape(repoName))).
		EndStruct(&v)
	return v, CheckResponse(errs, resp, 200)
}

// GetRepositoryTop
// Get the most popular public repositories
func (s *RepositoryClient) GetRepositoryTop(top interface{}) ([]RepoRecord, error) {
	var v []RepoRecord
	resp, _, errs := s.NewRequest(gorequest.GET, func() string {
		if t, ok := top.(int); ok {
			return fmt.Sprintf("/top?count=%d", t)
		}
		return fmt.Sprintf("/top")
	}()).
		EndStruct(&v)
	return v, CheckResponse(errs, resp, 200)
}
