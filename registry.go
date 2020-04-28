package harbor

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/url"
)

// RegistryClient handles communication with the registry related methods of the Harbor API
type RegistryClient struct {
	*Client
}

// CreateRegistry
// Create a registry
func (s *RegistryClient) CreateRegistry(r Registry) error {
	resp, _, errs := s.NewRequest(gorequest.POST, "").
		Send(r).
		End()
	return CheckResponse(errs, resp, 201)
}

// ListRegistries
// Get a list of registries
func (s *RegistryClient) ListRegistries(name string) ([]Registry, error) {
	var r []Registry
	resp, _, errs := s.NewRequest(gorequest.GET, "").
		Query(map[string]string{"name": name}).
		EndStruct(&r)
	return r, CheckResponse(errs, resp, 200)
}

// GetRegistryByID
// Get a registry by ID
func (s *RegistryClient) GetRegistryByID(id int64) (Registry, error) {
	var v Registry
	resp, _, errs := s.NewRequest(gorequest.GET, "/"+I64toA(id)).
		EndStruct(&v)
	return v, CheckResponse(errs, resp, 200)
}

// GetRegistryInfoByID
// Get a registry's info by ID
func (s *RegistryClient) GetRegistryInfoByID(id int64) (RegistryInfo, error) {
	var v RegistryInfo
	resp, _, errs := s.NewRequest(gorequest.GET,fmt.Sprintf("/%d/info", id)).
		EndStruct(&v)
	return v, CheckResponse(errs, resp, 200)
}

// DeleteRegistryByID
// Delete a registry by ID
func (s *RegistryClient) DeleteRegistryByID(id int64) error {
	resp, _, errs := s.NewRequest(gorequest.DELETE, "/"+I64toA(id)).
		End()
	return CheckResponse(errs, resp, 200)
}

// UpdateRegistryByID
// Update a registry by ID
func (s *RegistryClient) UpdateRegistryByID(registryName string, r Registry) error {
	resp, _, errs := s.NewRequest(gorequest.PUT,"/"+url.PathEscape(registryName)).
		Send(r).
		End()
	return CheckResponse(errs, resp, 200)
}

// PingRegistry
// Ping a registry and return it's health status
func (s *RegistryClient) PingRegistry(r Registry) error {
	resp, _, errs := s.NewRequest(gorequest.POST, "/ping").
		Send(r).
		End()
	return CheckResponse(errs, resp, 200)
}
