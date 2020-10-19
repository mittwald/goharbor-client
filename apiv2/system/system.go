package system

import (
	"context"

	"github.com/go-openapi/runtime"
	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client"

	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client/products"
	model "github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"
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
	NewSystemGarbageCollection(ctx context.Context, cron, scheduleType string) (*model.AdminJobSchedule, error)
	UpdateSystemGarbageCollection(ctx context.Context, newGcSchedule *model.AdminJobScheduleObj) error
	GetSystemGarbageCollection(ctx context.Context) (*model.AdminJobSchedule, error)
	ResetSystemGarbageCollection(ctx context.Context) error
}

// NewSystemGarbageCollection creates a new garbage collection schedule.
func (c *RESTClient) NewSystemGarbageCollection(ctx context.Context, cron,
	scheduleType string) (*model.AdminJobSchedule, error) {
	gcReq := &model.AdminJobSchedule{
		Schedule: &model.AdminJobScheduleObj{
			Cron: cron,
			Type: scheduleType,
		},
	}

	_, err := c.LegacyClient.Products.PostSystemGcSchedule(
		&products.PostSystemGcScheduleParams{
			Schedule: gcReq,
			Context:  ctx,
		}, c.AuthInfo)

	err = handleSwaggerSystemErrors(err)
	if err != nil {
		return nil, err
	}

	systemGc, err := c.GetSystemGarbageCollection(ctx)
	if err != nil {
		return nil, err
	}

	return systemGc, nil
}

// UpdateSystemGarbageCollection updates the system GC schedule.
func (c *RESTClient) UpdateSystemGarbageCollection(ctx context.Context,
	newGcSchedule *model.AdminJobScheduleObj) error {
	if newGcSchedule == nil {
		return &ErrSystemGcScheduleNotProvided{}
	}

	systemGc, err := c.GetSystemGarbageCollection(ctx)
	if err != nil {
		return err
	}

	if systemGc.Schedule == newGcSchedule {
		return &ErrSystemGcScheduleIdentical{}
	}

	_, err = c.LegacyClient.Products.PutSystemGcSchedule(
		&products.PutSystemGcScheduleParams{
			Schedule: &model.AdminJobSchedule{Schedule: newGcSchedule},
			Context:  ctx,
		}, c.AuthInfo)

	return handleSwaggerSystemErrors(err)
}

// GetSystemGarbageCollection returns the system GC schedule.
func (c *RESTClient) GetSystemGarbageCollection(ctx context.Context) (*model.AdminJobSchedule, error) {
	systemGc, err := c.LegacyClient.Products.GetSystemGcSchedule(
		&products.GetSystemGcScheduleParams{
			Context: ctx,
		}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerSystemErrors(err)
	}

	if systemGc.Payload.Schedule == nil {
		return nil, &ErrSystemGcUndefined{}
	}

	return systemGc.Payload, nil
}

// ResetSystemGarbageCollection resets the system GC schedule to it's default values
// containing "None" as the Schedule Type, which effectively deactivates the schedule.
// For this to work correctly, a GC schedule must exist beforehand.
func (c *RESTClient) ResetSystemGarbageCollection(ctx context.Context) error {
	_, err := c.LegacyClient.Products.PutSystemGcSchedule(
		&products.PutSystemGcScheduleParams{
			Schedule: &model.AdminJobSchedule{
				Schedule: &model.AdminJobScheduleObj{
					Type: "None",
				},
			},
			Context: ctx,
		}, c.AuthInfo)

	return handleSwaggerSystemErrors(err)
}

// Health reports Harbor system health information.
func (c *RESTClient) Health(ctx context.Context) (*model.OverallHealthStatus, error) {
	resp, err := c.LegacyClient.Products.GetHealth(&products.GetHealthParams{
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
