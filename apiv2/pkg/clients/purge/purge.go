package purge

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/runtime"
	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/purge"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
	"github.com/pkg/errors"
)

// RESTClient is a subclient for handling purge related actions.
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
	CreatePurgeSchedule(ctx context.Context, schedule *model.Schedule) error
	RunPurge(ctx context.Context, dryRun bool) error
	ListPurgeHistory(ctx context.Context) ([]*model.ExecHistory, error)
	GetPurgeJob(ctx context.Context, id int64) (*model.ExecHistory, error)
	GetPurgeJobLog(ctx context.Context, id int64) (string, error)
	GetPurgeSchedule(ctx context.Context) (*model.ExecHistory, error)
	StopPurge(ctx context.Context, id int64) error
	UpdatePurgeSchedule(ctx context.Context, schedule *model.Schedule) error
}

// CreatePurgeSchedule creates a new purge schedule.
func (c *RESTClient) CreatePurgeSchedule(ctx context.Context, schedule *model.Schedule) error {
	params := &purge.CreatePurgeScheduleParams{
		Schedule: schedule,
		Context:  ctx,
	}
	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Purge.CreatePurgeSchedule(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerPurgeErrors(err)
	}

	return nil
}

// RunPurge runs a manual purge job.
func (c *RESTClient) RunPurge(ctx context.Context, dryRun bool) error {
	schedule, err := c.GetPurgeSchedule(ctx)
	if err != nil {
		return handleSwaggerPurgeErrors(err)
	}
	if schedule.Schedule == nil {
		return errors.New("no schedule found")
	}
	s := schedule.Schedule
	s.Type = "Manual"

	// schedule.JobParameters is a string containing json
	parameters := make(map[string]interface{})
	err = json.Unmarshal([]byte(schedule.JobParameters), &parameters)
	parameters["dry_run"] = dryRun

	if err != nil {
		return errors.Wrap(err, "failed to unmarshal job parameters")
	}

	params := &purge.CreatePurgeScheduleParams{
		Schedule: &model.Schedule{
			ID:         schedule.ID,
			Parameters: parameters,
			Schedule:   s,
		},
		Context: ctx,
	}
	params.WithTimeout(c.Options.Timeout)

	_, err = c.V2Client.Purge.CreatePurgeSchedule(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerPurgeErrors(err)
	}

	return nil
}

// ListPurgeHistory lists all purge history entries.
// While the APIs purge service exposes a method called
// 'GetPurgeHistory', it technically returns a list of purge schedules.
func (c *RESTClient) ListPurgeHistory(ctx context.Context) ([]*model.ExecHistory, error) {
	var history []*model.ExecHistory
	page := c.Options.Page

	params := purge.NewGetPurgeHistoryParams()
	params.WithPage(&page)
	params.WithContext(ctx)
	params.WithTimeout(c.Options.Timeout)
	params.WithPageSize(&c.Options.PageSize)
	params.WithQ(&c.Options.Query)
	params.WithSort(&c.Options.Sort)

	for {
		resp, err := c.V2Client.Purge.GetPurgeHistory(params, c.AuthInfo)
		if err != nil {
			return nil, handleSwaggerPurgeErrors(err)
		}

		if len(resp.Payload) == 0 {
			break
		}

		totalCount := resp.XTotalCount

		history = append(history, resp.Payload...)

		if int64(len(history)) >= totalCount {
			break
		}

		page++
	}

	return history, nil
}

func (c *RESTClient) GetPurgeJob(ctx context.Context, id int64) (*model.ExecHistory, error) {
	params := &purge.GetPurgeJobParams{
		PurgeID: id,
		Context: ctx,
	}
	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Purge.GetPurgeJob(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerPurgeErrors(err)
	}

	return resp.Payload, nil
}

func (c *RESTClient) GetPurgeJobLog(ctx context.Context, id int64) (string, error) {
	params := &purge.GetPurgeJobLogParams{
		PurgeID: id,
		Context: ctx,
	}
	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Purge.GetPurgeJobLog(params, c.AuthInfo)
	if err != nil {
		return "", handleSwaggerPurgeErrors(err)
	}

	return resp.Payload, nil
}

func (c *RESTClient) GetPurgeSchedule(ctx context.Context) (*model.ExecHistory, error) {
	params := &purge.GetPurgeScheduleParams{
		Context: ctx,
	}
	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Purge.GetPurgeSchedule(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerPurgeErrors(err)
	}

	return resp.Payload, nil
}

func (c *RESTClient) StopPurge(ctx context.Context, id int64) error {
	params := &purge.StopPurgeParams{
		PurgeID: id,
		Context: ctx,
	}
	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Purge.StopPurge(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerPurgeErrors(err)
	}

	return nil
}

func (c *RESTClient) UpdatePurgeSchedule(ctx context.Context, schedule *model.Schedule) error {
	params := &purge.UpdatePurgeScheduleParams{
		Schedule: schedule,
		Context:  ctx,
	}
	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Purge.UpdatePurgeSchedule(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerPurgeErrors(err)
	}

	return nil
}
