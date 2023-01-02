//go:build integration

package purge

import (
	"context"
	"testing"
	"time"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/stretchr/testify/require"
)

func TestAPICreatePurgeSchedule(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.CreatePurgeSchedule(ctx, &model.Schedule{
		Parameters: map[string]interface{}{
			"audit_retention_hour": 168,
			"dry_run":              false,
			"include_operations":   "create,delete,pull",
		},
		Schedule: &model.ScheduleObj{
			Cron: "0 0 * * * *",
			Type: "Hourly",
		},
	})

	require.NoError(t, err)
}

func TestAPIRunPurge(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.RunPurge(ctx, true)
	require.NoError(t, err)
}

func TestAPIListPurgeHistory(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	purgeHistory, err := c.ListPurgeHistory(ctx)
	require.NoError(t, err)
	require.NotNil(t, purgeHistory)
}

func TestAPIGetPurgeJob(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	hist, err := c.ListPurgeHistory(ctx)
	require.NoError(t, err)

	purgeJob, err := c.GetPurgeJob(ctx, hist[0].ID)
	require.NoError(t, err)
	require.NotNil(t, purgeJob)
}

func TestAPIGetPurgeSchedule(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	purgeSchedule, err := c.GetPurgeSchedule(ctx)
	require.NoError(t, err)
	require.NotNil(t, purgeSchedule)
}
func TestAPIStopPurge(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.RunPurge(ctx, true)
	require.NoError(t, err)

	purgeJobs, err := c.ListPurgeHistory(ctx)
	require.NoError(t, err)
	require.NotNil(t, purgeJobs)

	err = c.StopPurge(ctx, purgeJobs[0].ID)
	require.NoError(t, err)
}
func TestAPIUpdatePurgeSchedule(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.UpdatePurgeSchedule(ctx, &model.Schedule{
		Parameters: map[string]interface{}{
			"audit_retention_hour": 168,
			"dry_run":              true,
			"include_operations":   "create,delete,pull",
		},
		Schedule: &model.ScheduleObj{
			Cron: "0 0 * * * *",
			Type: "Hourly",
		},
	})

	require.NoError(t, err)
}

func TestAPIGetPurgeJobLog(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.RunPurge(ctx, true)
	require.NoError(t, err)
	jobs, err := c.ListPurgeHistory(ctx)
	require.NoError(t, err)

	require.Eventually(t, func() bool {
		purgeJobLog, err := c.GetPurgeJobLog(ctx, jobs[0].ID)
		logExists := purgeJobLog != ""
		done := err == nil && logExists
		return done
	}, 10*time.Second, 1*time.Second)
}
