package harbor

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

// Implementations of RegistryClient handle communication with
// registry related methods of Harbor.
type RegistryClient interface {
	// Create a new registry.
	CreateRegistry(r Registry) error

	// List all registries.
	ListRegistries(name string) ([]Registry, error)

	// Get a registry by id.
	GetRegistryByID(id int64) (Registry, error)

	// Get information about a registry by id.
	GetRegistryInfoByID(id int64) (RegistryInfo, error)

	// delete a registry by id.
	DeleteRegistryByID(id int64) error

	// update a registry by id.
	UpdateRegistryByID(r Registry) error

	// ping a registry and return it's health status.
	PingRegistry(r Registry) error
}

// RestRegistryClient implements the RegistryClient interface by communicating via Rest api.
type RestRegistryClient struct {
	*RestClient
}

// CreateRegistry satisfies the RegistryClient interface.
func (s *RestRegistryClient) CreateRegistry(r Registry) error {
	resp, _, errs := s.NewRequest(gorequest.POST, "").
		Send(r).
		End()
	return CheckResponse(errs, resp, 201)
}

// ListRegistries satisfies the RegistryClient interface.
func (s *RestRegistryClient) ListRegistries(name string) ([]Registry, error) {
	var r []Registry
	resp, _, errs := s.NewRequest(gorequest.GET, "").
		Query(map[string]string{"name": name}).
		EndStruct(&r)
	return r, CheckResponse(errs, resp, 200)
}

// GetRegistryByID satisfies the RegistryClient interface.
func (s *RestRegistryClient) GetRegistryByID(id int64) (Registry, error) {
	var v Registry
	resp, _, errs := s.NewRequest(gorequest.GET, "/"+I64toA(id)).
		EndStruct(&v)
	return v, CheckResponse(errs, resp, 200)
}

// GetRegistryInfoByID satisfies the RegistryClient interface.
func (s *RestRegistryClient) GetRegistryInfoByID(id int64) (RegistryInfo, error) {
	var v RegistryInfo
	resp, _, errs := s.NewRequest(gorequest.GET,fmt.Sprintf("/%d/info", id)).
		EndStruct(&v)
	return v, CheckResponse(errs, resp, 200)
}

// DeleteRegistryByID satisfies the RegistryClient interface.
func (s *RestRegistryClient) DeleteRegistryByID(id int64) error {
	resp, _, errs := s.NewRequest(gorequest.DELETE, "/"+I64toA(id)).
		End()
	return CheckResponse(errs, resp, 200)
}

// UpdateRegistryByID satisfies the RegistryClient interface.
func (s *RestRegistryClient) UpdateRegistryByID(r Registry) error {
	resp, _, errs := s.NewRequest(gorequest.PUT,"/"+I64toA(r.ID)).
		Send(r).
		End()
	return CheckResponse(errs, resp, 200)
}

// PingRegistry satisfies the RegistryClient interface.
func (s *RestRegistryClient) PingRegistry(r Registry) error {
	resp, _, errs := s.NewRequest(gorequest.POST, "/ping").
		Send(r).
		End()
	return CheckResponse(errs, resp, 200)
}
