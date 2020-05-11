package harbor

import (
	"github.com/parnurzeal/gorequest"
)

// ReplicationClient abstracts away the communication implementation of
// replication related methods of Harbor.
type ReplicationClient interface {
	// List all replication adapters.
	ListReplicationAdapters() ([]string, error)

	// List all replication policies.
	ListReplicationPolicies(name string) ([]ReplicationPolicy, error)

	// Retrieves a replication policy by id.
	GetReplicationPolicyByID(id int64) (ReplicationPolicy, error)

	// Update a replication policy by id.
	UpdateReplicationPolicyByID(id int64, policy ReplicationPolicy) error

	// Delete a replication policy by id.
	DeleteReplicationPolicyByID(id int64) error

	// Create a replication policy by id.
	CreateReplicationPolicy(rp ReplicationPolicy) error

	// Get a replication execution by id.
	GetReplicationExecutionsByID(id int64) (ReplicationExecution, error)

	// Trigger a new replication execution.
	TriggerReplicationExecution(e ReplicationExecution) error

	// Get a replication execution by policy id.
	GetReplicationExecutions(policyID int64) ([]ReplicationExecution, error)
}

// RestReplicationClient satisfies the ReplicationClient interface by communicating via Rest api.
type RestReplicationClient struct {
	*RestClient
}

// ListReplicationAdapters satisfies the ReplicationClient interface.
func (s *RestReplicationClient) ListReplicationAdapters() ([]string, error) {
	var r []string
	resp, _, errs := s.NewRequest(gorequest.GET, "/adapters").
		EndStruct(&r)
	return r, CheckResponse(errs, resp, 200)
}

// ListReplicationPolicies satisfies the ReplicationClient interface.
func (s *RestReplicationClient) ListReplicationPolicies(name string) ([]ReplicationPolicy, error) {
	var rp []ReplicationPolicy
	resp, _, errs := s.NewRequest(gorequest.GET, "/policies").
		Query(map[string]string{"name": name}).
		EndStruct(&rp)
	return rp, CheckResponse(errs, resp, 200)
}

// GetReplicationPolicyByID satisfies the ReplicationClient interface.
func (s *RestReplicationClient) GetReplicationPolicyByID(id int64) (ReplicationPolicy, error) {
	var r ReplicationPolicy
	resp, _, errs := s.NewRequest(gorequest.GET, "/policies/"+I64toA(id)).
		EndStruct(&r)
	return r, CheckResponse(errs, resp, 200)
}

// UpdateReplicationPolicyByID satisfies the ReplicationClient interface.
func (s *RestReplicationClient) UpdateReplicationPolicyByID(id int64, policy ReplicationPolicy) error {
	resp, _, errs := s.NewRequest(gorequest.PUT, "/policies/"+I64toA(id)).
		Send(policy).
		End()
	return CheckResponse(errs, resp, 200)
}

// DeleteReplicationPolicyByID satisfies the ReplicationClient interface.
func (s *RestReplicationClient) DeleteReplicationPolicyByID(id int64) error {
	resp, _, errs := s.NewRequest(gorequest.DELETE, "/policies/"+I64toA(id)).
		End()
	return CheckResponse(errs, resp, 200)
}

// CreateReplicationPolicy satisfies the ReplicationClient interface.
func (s *RestReplicationClient) CreateReplicationPolicy(rp ReplicationPolicy) error {
	resp, _, errs := s.NewRequest(gorequest.POST, "/policies").
		Send(rp).
		End()
	return CheckResponse(errs, resp, 201)
}

// GetReplicationExecutionsByID satisfies the ReplicationClient interface.
func (s *RestReplicationClient) GetReplicationExecutionsByID(id int64) (ReplicationExecution, error) {
	var r ReplicationExecution
	resp, _, errs := s.NewRequest(gorequest.GET, "/executions/"+I64toA(id)).
		EndStruct(&r)
	return r, CheckResponse(errs, resp, 200)
}

// TriggerReplicationExecution satisfies the ReplicationClient interface.
func (s *RestReplicationClient) TriggerReplicationExecution(e ReplicationExecution) error {
	resp, _, errs := s.NewRequest(gorequest.POST, "/executions").
		Send(e).
		End()
	return CheckResponse(errs, resp, 201)
}

// GetReplicationExecutions satisfies the ReplicationClient interface.
func (s *RestReplicationClient) GetReplicationExecutions(policyID int64) ([]ReplicationExecution, error) {
	var r []ReplicationExecution
	resp, _, errs := s.NewRequest(gorequest.GET, "/executions/").
		Query(map[string]string{"policy_id": I64toA(policyID)}).
		EndStruct(&r)
	return r, CheckResponse(errs, resp, 200)
}
