package harbor

import (
	"github.com/parnurzeal/gorequest"
)

// ReplicationClient handles communication with the replication related methods of the Harbor API
type ReplicationClient struct {
	*Client
}

// ListReplicationAdapters
// Get all replication adapters
func (s *ReplicationClient) ListReplicationAdapters() ([]string, error) {
	var r []string
	resp, _, errs := s.NewRequest(gorequest.GET, "/adapters").
		EndStruct(&r)
	return r, CheckResponse(errs, resp, 200)
}

// ListReplicationPolicies
// Get an array of matching replication policies
func (s *ReplicationClient) ListReplicationPolicies(name string) ([]ReplicationPolicy, error) {
	var rp []ReplicationPolicy
	resp, _, errs := s.NewRequest(gorequest.GET, "/policies").
		Query(map[string]string{"name": name}).
		EndStruct(&rp)
	return rp, CheckResponse(errs, resp, 200)
}

// GetReplicationPolicies
// Get a replication policy by ID
func (s *ReplicationClient) GetReplicationPolicyByID(id int64) (ReplicationPolicy, error) {
	var r ReplicationPolicy
	resp, _, errs := s.NewRequest(gorequest.GET, "/policies"+I64toA(id)).
		EndStruct(&r)
	return r, CheckResponse(errs, resp, 200)
}

// UpdateReplicationPolicyByID
// Update a replication policy by ID
func (s *ReplicationClient) UpdateReplicationPolicyByID(id int64, policy ReplicationPolicy) error {
	resp, _, errs := s.NewRequest(gorequest.PUT, "/policies/"+I64toA(id)).
		Send(policy).
		End()
	return CheckResponse(errs, resp, 200)
}

// DeleteReplicationPolicyByID
// Delete a replication policy by ID
func (s *ReplicationClient) DeleteReplicationPolicyByID(id int64) error {
	resp, _, errs := s.NewRequest(gorequest.DELETE, "/policies/"+I64toA(id)).
		End()
	return CheckResponse(errs, resp, 200)
}

// CreateReplicationPolicy
// Create a replication policy
func (s *ReplicationClient) CreateReplicationPolicy(rp ReplicationPolicy) error {
	resp, _, errs := s.NewRequest(gorequest.POST, "/policies").
		Send(rp).
		End()
	return CheckResponse(errs, resp, 201)
}

// GetReplicationExecutionsByID
// Get replication executions filtered by replication ID
func (s *ReplicationClient) GetReplicationExecutionsByID(id int64) (ReplicationExecution, error) {
	var r ReplicationExecution
	resp, _, errs := s.NewRequest(gorequest.GET, "/executions/"+I64toA(id)).
		EndStruct(&r)
	return r, CheckResponse(errs, resp, 200)
}

// TriggerReplicationExecution
// Trigger a replication execution
func (s *ReplicationClient) TriggerReplicationExecution(e ReplicationExecution) error {
	resp, _, errs := s.NewRequest(gorequest.POST, "/executions").
		Send(e).
		End()
	return CheckResponse(errs, resp, 201)
}

// GetReplicationPolicies
// Get an array of matching replication policies
func (s *ReplicationClient) GetReplicationExecutions(policyID int64) ([]ReplicationExecution, error) {
	var r []ReplicationExecution
	resp, _, errs := s.NewRequest(gorequest.GET, "/executions/").
		Query(map[string]string{"policy_id": I64toA(policyID)}).
		EndStruct(&r)
	return r, CheckResponse(errs, resp, 200)
}
