package gc

import (
	"context"

	"github.com/go-openapi/runtime"

	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/gc"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
)

// RESTClient is a subclient for handling garbage collection related actions.
type RESTClient struct {
	// Options contains optional configuration when making API calls.
	Options *config.Options

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(v2Client *v2client.Harbor, opts *config.Options, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Options:  opts,
		V2Client: v2Client,
		AuthInfo: authInfo,
	}
}

type Client interface {
	NewGarbageCollection(ctx context.Context, gcSchedule *model.Schedule) error
	UpdateGarbageCollection(ctx context.Context,
		newGCSchedule *model.Schedule) error
	GetGarbageCollectionExecutions(ctx context.Context) ([]*model.GCHistory, error)
	GetGarbageCollectionExecution(ctx context.Context, id int64) (*model.GCHistory, error)
	GetGarbageCollectionSchedule(ctx context.Context) (*model.GCHistory, error)
	ResetGarbageCollection(ctx context.Context) error
}

// NewGarbageCollection creates a new garbage collection schedule.
func (c *RESTClient) NewGarbageCollection(ctx context.Context, gcSchedule *model.Schedule) error {
	if gcSchedule == nil {
		return &errors.ErrSystemGcScheduleNotProvided{}
	}

	if gcSchedule.Parameters == nil {
		gcSchedule.Parameters = map[string]interface{}{
			"delete_untagged": false,
		}
	}

	_, err := c.V2Client.GC.CreateGCSchedule(&gc.CreateGCScheduleParams{
		Schedule: gcSchedule,
		Context:  ctx,
	}, c.AuthInfo)

	err = handleSwaggerSystemErrors(err)
	if err != nil {
		return err
	}

	return nil
}

// UpdateGarbageCollection updates the system GC schedule.
func (c *RESTClient) UpdateGarbageCollection(ctx context.Context,
	newGCSchedule *model.Schedule,
) error {
	if newGCSchedule == nil {
		return &errors.ErrSystemGcScheduleNotProvided{}
	}
	if newGCSchedule.Parameters == nil {
		newGCSchedule.Parameters = map[string]interface{}{
			"delete_untagged": false,
		}
	}

	_, err := c.V2Client.GC.UpdateGCSchedule(&gc.UpdateGCScheduleParams{
		Schedule: newGCSchedule,
		Context:  ctx,
	}, c.AuthInfo)

	return handleSwaggerSystemErrors(err)
}

// GetGarbageCollectionExecutions Returns the garbage collection executions.
func (c *RESTClient) GetGarbageCollectionExecutions(ctx context.Context) ([]*model.GCHistory, error) {
	resp, err := c.V2Client.GC.GetGCHistory(&gc.GetGCHistoryParams{
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerSystemErrors(err)
	}

	return resp.Payload, nil
}

// GetGarbageCollectionExecution Returns a garbage collection execution identified by its id.
func (c *RESTClient) GetGarbageCollectionExecution(ctx context.Context, id int64) (*model.GCHistory, error) {
	resp, err := c.V2Client.GC.GetGC(&gc.GetGCParams{
		Context: ctx,
		GCID:    id,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerSystemErrors(err)
	}

	if resp.Payload.Schedule == nil {
		return nil, &errors.ErrSystemGcUndefined{}
	}

	return resp.Payload, nil
}

// GetGarbageCollectionSchedule returns the system GC schedule.
func (c *RESTClient) GetGarbageCollectionSchedule(ctx context.Context) (*model.GCHistory, error) {
	resp, err := c.V2Client.GC.GetGCSchedule(&gc.GetGCScheduleParams{
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerSystemErrors(err)
	}

	if resp.Payload.Schedule != nil {
		return resp.Payload, nil
	}

	return nil, &errors.ErrSystemGcScheduleUndefined{}
}

// ResetGarbageCollection resets the system GC schedule to it's default values
// containing "None" as the Schedule Type, which effectively deactivates the schedule.
// For this to work correctly, a GC schedule must exist beforehand.
func (c *RESTClient) ResetGarbageCollection(ctx context.Context) error {
	_, err := c.V2Client.GC.UpdateGCSchedule(&gc.UpdateGCScheduleParams{
		Schedule: &model.Schedule{
			Parameters: map[string]interface{}{
				"delete_untagged": false,
			},
			Schedule: &model.ScheduleObj{
				Type: "None",
			},
		},
		Context: ctx,
	}, c.AuthInfo)

	return handleSwaggerSystemErrors(err)
}
