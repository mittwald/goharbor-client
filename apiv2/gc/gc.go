package gc

import (
	"context"

	"github.com/go-openapi/runtime"
	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client/gc"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client"

	modelv2 "github.com/mittwald/goharbor-client/v3/apiv2/model"
)

// RESTClient is a subclient for handling system related actions.
type RESTClient struct {
	// The legacy swagger client
	LegacyClient *client.Harbor

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(legacyClient *client.Harbor, v2Client *v2client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		LegacyClient: legacyClient,
		V2Client:     v2Client,
		AuthInfo:     authInfo,
	}
}

type Client interface {
	NewSystemGarbageCollection(ctx context.Context, gcSchedule *modelv2.Schedule) error
	UpdateSystemGarbageCollection(ctx context.Context, newGCSchedule *modelv2.Schedule) error
	GetSystemGarbageCollectionSchedule(ctx context.Context) (*modelv2.GCHistory, error)
	ResetSystemGarbageCollection(ctx context.Context) error
}

// NewSystemGarbageCollection creates a new garbage collection schedule.
func (c *RESTClient) NewSystemGarbageCollection(ctx context.Context, gcSchedule *modelv2.Schedule) error {
	if gcSchedule == nil {
		return &ErrSystemGcScheduleNotProvided{}
	}

	if gcSchedule.Parameters == nil {
		gcSchedule.Parameters = make(map[string]interface{})
		gcSchedule.Parameters["delete_untagged"] = false
	}

	_, err := c.V2Client.Gc.CreateGCSchedule(&gc.CreateGCScheduleParams{
		Schedule: gcSchedule,
		Context:  ctx,
	}, c.AuthInfo)

	err = handleSwaggerSystemErrors(err)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSystemGarbageCollection updates the system GC schedule.
func (c *RESTClient) UpdateSystemGarbageCollection(ctx context.Context,
	newGCSchedule *modelv2.Schedule) error {
	if newGCSchedule == nil {
		return &ErrSystemGcScheduleNotProvided{}
	}
	if newGCSchedule.Parameters == nil {
		newGCSchedule.Parameters = make(map[string]interface{})
		newGCSchedule.Parameters["delete_untagged"] = false
	}

	_, err := c.V2Client.Gc.UpdateGCSchedule(&gc.UpdateGCScheduleParams{
		Schedule: newGCSchedule,
		Context:  ctx,
	}, c.AuthInfo)

	return handleSwaggerSystemErrors(err)
}

// GetSystemGarbageCollectionExecution Returns a garbage collection execution identified by its id.
func (c *RESTClient) GetSystemGarbageCollectionExecution(ctx context.Context, id int64) (*modelv2.GCHistory, error) {
	resp, err := c.V2Client.Gc.GetGC(&gc.GetGCParams{
		Context: ctx,
		GcID:    id,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerSystemErrors(err)
	}

	if resp.Payload.Schedule == nil {
		return nil, &ErrSystemGcUndefined{}
	}

	return resp.Payload, nil
}

// GetSystemGarbageCollectionSchedule returns the system GC schedule.
func (c *RESTClient) GetSystemGarbageCollectionSchedule(ctx context.Context) (*modelv2.GCHistory, error) {
	resp, err := c.V2Client.Gc.GetGCSchedule(&gc.GetGCScheduleParams{
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerSystemErrors(err)
	}

	if resp.Payload.Schedule != nil {
		return resp.Payload, nil
	}

	return nil, &ErrSystemGcScheduleUndefined{}
}

// ResetSystemGarbageCollection resets the system GC schedule to it's default values
// containing "None" as the Schedule Type, which effectively deactivates the schedule.
// For this to work correctly, a GC schedule must exist beforehand.
func (c *RESTClient) ResetSystemGarbageCollection(ctx context.Context) error {
	_, err := c.V2Client.Gc.UpdateGCSchedule(&gc.UpdateGCScheduleParams{
		Schedule: &modelv2.Schedule{
			Parameters: map[string]interface{}{
				"delete_untagged": false,
			},
			Schedule: &modelv2.ScheduleObj{
				Type: "None",
			},
		},
		Context: ctx,
	}, c.AuthInfo)

	return handleSwaggerSystemErrors(err)
}
