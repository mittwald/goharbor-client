package harbor

import "github.com/parnurzeal/gorequest"

// SystemClient handles communication with the system related methods of the Harbor API
type SystemClient struct {
	*Client
}

// GetSystemGarbageCollection
// Get the system's configured garbage collection schedule
func (s *SystemClient) GetSystemGarbageCollectionSchedule() (AdminJobReq, error) {
	var gc AdminJobReq
	resp, _, errs := s.NewRequest(gorequest.GET, "/gc/schedule").
		EndStruct(&gc)
	return gc, CheckResponse(errs, resp, 200)
}

// CreateSystemGarbageCollectionSchedule
// Create a garbage collection schedule for the system
func (s *SystemClient) CreateSystemGarbageCollectionSchedule(gc AdminJobReq) error {
	resp, _, errs := s.NewRequest(gorequest.POST, "/gc/schedule").
		Send(gc).
		End()
	return CheckResponse(errs, resp, 201)
}

// UpdateSystemGarbageCollectionSchedule
// Update the system's configured garbage collection schedule
func (s *SystemClient) UpdateSystemGarbageCollectionSchedule(gc AdminJobReq) error {
	resp, _, errs := s.NewRequest(gorequest.PUT, "/gc/schedule").
		Send(gc).
		End()
	return CheckResponse(errs, resp, 200)
}
