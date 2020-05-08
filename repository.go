package harbor

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/url"
)

// Implementations of RepositoryClient handle communication with
// repository related methods of Harbor.
type RepositoryClient interface {
	// ListRepository lists repositories filtered by the
	// relevant project ID and repository name.
	ListRepository(opt *RepositoryQuery) ([]RepoRecord, error)

	// DeleteRepository deletes a repository by repository name.
	DeleteRepository(repoName string) error

	// UpdateRepository updates the description of a repository by repository name.
	UpdateRepository(repoName string, d RepositoryDescription) error

	// GetRepositoryTag retrieves the tag of a repository.
	// NOTE: If deployed with Notary, the signature property of
	// response represents whether the image is signed or not
	// If the property is null, the image is unsigned
	GetRepositoryTag(repoName, tag string) (TagResp, error)

	// DeleteRepositoryTag deletes a repository tag by repository and tag name.
	DeleteRepositoryTag(repoName, tag string) error

	// ListRepositoryTags retrieves tags from a repository.
	// NOTE: If deployed with Notary, the signature property of response represents whether the image is signed or not
	// If the property is null, the image is unsigned
	ListRepositoryTags(repoName string) ([]TagResp, error)

	// GetRepositoryTagManifests retrieves manifests from a relevant repository.
	GetRepositoryTagManifests(repoName, tag string, version string) (ManifestResp, error)

	// ScanImage triggers the jobservice component to call the Clair API to scan the image.
	// Only accessible for project admins.
	ScanImage(repoName, tag string) error

	// GetImageScan retrieves information from the Clair API containing vulnerability
	// information based on the previous successful scan.
	GetImageScan(repoName, tag string) ([]VulnerabilityItem, error)

	// GetRepositorySignature retrieves signature information of a repository,
	// originating from the notary component of Harbor.
	// NOTE: If the repository does not have any signature information in notary, this API will
	// return an empty list with response code 200, instead of 404
	GetRepositorySignature(repoName string) ([]Signature, error)

	// GetRepositoryTop retrieves the most popular public repositories
	GetRepositoryTop(top interface{}) ([]RepoRecord, error)
}

// RestRepositoryClient implements the RepositoryClient interface by communicating via Rest api.
type RestRepositoryClient struct {
	*RestClient
}

// ListRepository satisfies the RepositoryClient interface.
func (s *RestRepositoryClient) ListRepository(opt *RepositoryQuery) ([]RepoRecord, error) {
	var v []RepoRecord
	resp, _, errs := s.NewRequest(gorequest.GET, "").
		Query(*opt).
		EndStruct(&v)
	return v, CheckResponse(errs, resp, 200)
}

// DeleteRepository satisfies the RepositoryClient interface.
func (s *RestRepositoryClient) DeleteRepository(repoName string) error {
	resp, _, errs := s.NewRequest(gorequest.DELETE,"/"+url.PathEscape(repoName)).
		End()
	return CheckResponse(errs, resp, 200)
}

// UpdateRepository satisfies the RepositoryClient interface.
func (s *RestRepositoryClient) UpdateRepository(repoName string, d RepositoryDescription) error {
	resp, _, errs := s.NewRequest(gorequest.PUT,"/"+url.PathEscape(repoName)).
		Send(d).
		End()
	return CheckResponse(errs, resp, 200)
}

// GetRepositoryTag satisfies the RepositoryClient interface.
func (s *RestRepositoryClient) GetRepositoryTag(repoName, tag string) (TagResp, error) {
	var v TagResp
	resp, _, errs := s.NewRequest(gorequest.GET,
		fmt.Sprintf("/%s/tags/%s", url.PathEscape(repoName), url.PathEscape(tag))).
		EndStruct(&v)
	return v, CheckResponse(errs, resp, 200)
}

// DeleteRepositoryTag satisfies the RepositoryClient interface.
func (s *RestRepositoryClient) DeleteRepositoryTag(repoName, tag string) error {
	resp, _, errs := s.NewRequest(gorequest.DELETE,fmt.Sprintf("/%s/tags/%s",
		url.PathEscape(repoName), url.PathEscape(tag))).
		End()
	return CheckResponse(errs, resp, 200)
}

// ListRepositoryTags satisfies the RepositoryClient interface.
func (s *RestRepositoryClient) ListRepositoryTags(repoName string) ([]TagResp, error) {
	var v []TagResp
	resp, _, errs := s.NewRequest(gorequest.GET,
		fmt.Sprintf("/%s/tags", url.PathEscape(repoName))).
		EndStruct(&v)
	return v, CheckResponse(errs, resp, 200)
}

// GetRepositoryTagManifests satisfies the RepositoryClient interface.
func (s *RestRepositoryClient) GetRepositoryTagManifests(repoName, tag string, version string) (ManifestResp, error) {
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

// ScanImage satisfies the RepositoryClient interface.
func (s *RestRepositoryClient) ScanImage(repoName, tag string) error {
	resp, _, errs := s.NewRequest(gorequest.POST,fmt.Sprintf("/%s/tags/%s/scan",
		url.PathEscape(repoName), url.PathEscape(tag))).
		End()
	return CheckResponse(errs, resp, 202)
}

// GetImageScan satisfies the RepositoryClient interface.
func (s *RestRepositoryClient) GetImageScan(repoName, tag string) ([]VulnerabilityItem, error) {
	var v []VulnerabilityItem
	resp, _, errs := s.NewRequest(gorequest.GET,fmt.Sprintf("/%s/tags/%s/scan",
		url.PathEscape(repoName), url.PathEscape(tag))).
		EndStruct(&v)
	return v, CheckResponse(errs, resp, 200)
}

// GetRepositorySignature satisfies the RepositoryClient interface.
func (s *RestRepositoryClient) GetRepositorySignature(repoName string) ([]Signature, error) {
	var v []Signature
	resp, _, errs := s.NewRequest(gorequest.GET,
		fmt.Sprintf("/%s/signatures", url.PathEscape(repoName))).
		EndStruct(&v)
	return v, CheckResponse(errs, resp, 200)
}

// GetRepositoryTop satisfies the RepositoryClient interface.
func (s *RestRepositoryClient) GetRepositoryTop(top interface{}) ([]RepoRecord, error) {
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
