package scanall

import (
	"context"

	"github.com/go-openapi/runtime"

	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/scan_all"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
)

// RESTClient is a subclient for handling project related actions.
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
	CreateScanAllSchedule(ctx context.Context, schedule *model.Schedule) error
	GetScanAllSchedule(ctx context.Context) (*model.Schedule, error)
	UpdateScanAllSchedule(ctx context.Context, schedule *model.Schedule) error
}

// CreateScanAllSchedule creates a new scan all schedule.
func (c *RESTClient) CreateScanAllSchedule(ctx context.Context, schedule *model.Schedule) error {
	params := &scan_all.CreateScanAllScheduleParams{
		Context:  ctx,
		Schedule: schedule,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.ScanAll.CreateScanAllSchedule(params, c.AuthInfo)

	return handleSwaggerScanallErrors(err)
}

// GetScanAllSchedule returns the scan all schedule.
func (c *RESTClient) GetScanAllSchedule(ctx context.Context) (*model.Schedule, error) {
	params := &scan_all.GetScanAllScheduleParams{
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.ScanAll.GetScanAllSchedule(params, c.AuthInfo)
	if err != nil {
		return nil, err
	}
	if resp.Payload == nil {
		return nil, &errors.ErrNotFound{}
	}
	return resp.Payload, handleSwaggerScanallErrors(err)
}

// CreateScanAllSchedule creates a new scan all schedule.
func (c *RESTClient) UpdateScanAllSchedule(ctx context.Context, schedule *model.Schedule) error {
	params := &scan_all.UpdateScanAllScheduleParams{
		Context:  ctx,
		Schedule: schedule,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.ScanAll.UpdateScanAllSchedule(params, c.AuthInfo)

	return handleSwaggerScanallErrors(err)
}
