package goharborclient

import (
	"context"

	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/client/products"
	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/model"
)

// SystemRESTClient is a subclient for RESTClient handling system related
// actions.
type SystemRESTClient struct {
	Parent *RESTClient
}

// NewSystemGarbageCollection creates a new garbage collection schedule.
func (c *SystemRESTClient) NewSystemGarbageCollection(ctx context.Context, cron, scheduleType string) (*model.AdminJobSchedule, error) {
	gcReq := &model.AdminJobSchedule{
		Schedule: &model.AdminJobScheduleObj{
			Cron: cron,
			Type: scheduleType,
		}}

	_, err := c.Parent.Client.Products.PostSystemGcSchedule(
		&products.PostSystemGcScheduleParams{
			Schedule: gcReq,
			Context:  ctx,
		}, c.Parent.AuthInfo)

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
func (c *SystemRESTClient) UpdateSystemGarbageCollection(ctx context.Context, newGcSchedule *model.AdminJobScheduleObj) error {
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

	_, err = c.Parent.Client.Products.PutSystemGcSchedule(
		&products.PutSystemGcScheduleParams{
			Schedule: &model.AdminJobSchedule{Schedule: newGcSchedule},
			Context:  ctx,
		}, c.Parent.AuthInfo)

	return handleSwaggerSystemErrors(err)
}

// GetSystemGarbageCollection returns the system GC schedule.
func (c *SystemRESTClient) GetSystemGarbageCollection(ctx context.Context) (*model.AdminJobSchedule, error) {
	systemGc, err := c.Parent.Client.Products.GetSystemGcSchedule(
		&products.GetSystemGcScheduleParams{
			Context: ctx,
		}, c.Parent.AuthInfo)

	err = handleSwaggerSystemErrors(err)
	if err != nil {
		return nil, err
	}

	if systemGc.Payload.Schedule == nil {
		return nil, &ErrSystemGcUndefined{}
	}

	return systemGc.Payload, nil
}

// ResetSystemGarbageCollection resets the system GC schedule to it's default values
// containing "None" as the Schedule Type, which effectively deactivates the schedule.
// For this to work correctly, a GC schedule must exist beforehand.
func (c *SystemRESTClient) ResetSystemGarbageCollection(ctx context.Context) error {
	_, err := c.Parent.Client.Products.PutSystemGcSchedule(
		&products.PutSystemGcScheduleParams{
			Schedule: &model.AdminJobSchedule{
				Schedule: &model.AdminJobScheduleObj{
					Type: "None",
				}},
			Context: ctx,
		}, c.Parent.AuthInfo)

	return handleSwaggerSystemErrors(err)
}
