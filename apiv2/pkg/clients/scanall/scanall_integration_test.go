//go:build integration

package scanall

import (
	"context"
	"testing"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/stretchr/testify/require"
)

func TestAPIScanAllCreateScanAllSchedule(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	schedule, err := c.GetScanAllSchedule(ctx)
	require.NoError(t, err)
	if schedule.Schedule != nil {
		err = c.UpdateScanAllSchedule(ctx, &model.Schedule{
			Schedule: &model.ScheduleObj{
				Type: "None",
				Cron: "",
			},
		})
		require.NoError(t, err)
	}
	err = c.CreateScanAllSchedule(ctx, &model.Schedule{
		Schedule: &model.ScheduleObj{
			Type: "Daily",
			Cron: "0 0 0 * * *",
		},
		Parameters: nil,
	})

	defer c.UpdateScanAllSchedule(ctx, &model.Schedule{
		Schedule: &model.ScheduleObj{
			Type: "None",
			Cron: "",
		},
	})

	require.NoError(t, err)
}

func TestAPIScanAllGetScanAllSchedule(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	schedule, err := c.GetScanAllSchedule(ctx)
	require.NoError(t, err)

	if schedule.Schedule == nil {
		err = c.CreateScanAllSchedule(ctx, &model.Schedule{
			Schedule: &model.ScheduleObj{
				Type: "Daily",
				Cron: "0 0 0 * * *",
			},
			Parameters: nil,
		})
		require.NoError(t, err)
	}

	schedule, err = c.GetScanAllSchedule(ctx)

	defer c.UpdateScanAllSchedule(ctx, &model.Schedule{
		Schedule: &model.ScheduleObj{
			Type: "None",
			Cron: "",
		},
	})

	require.NoError(t, err)
	require.NotEmpty(t, schedule)
}

func TestAPIScanAllUpdateScanAllSchedule(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	schedule, err := c.GetScanAllSchedule(ctx)
	require.NoError(t, err)

	if schedule.Schedule == nil {
		err = c.CreateScanAllSchedule(ctx, &model.Schedule{
			Schedule: &model.ScheduleObj{
				Type: "Daily",
				Cron: "0 0 0 * * *",
			},
			Parameters: nil,
		})
		require.NoError(t, err)
	}

	err = c.UpdateScanAllSchedule(ctx, &model.Schedule{
		Schedule: &model.ScheduleObj{
			Type: "Weekly",
			Cron: "0 0 0 * * 0",
		},
		Parameters: nil,
	})
	require.NoError(t, err)

	schedule, err = c.GetScanAllSchedule(ctx)

	defer c.UpdateScanAllSchedule(ctx, &model.Schedule{
		Schedule: &model.ScheduleObj{
			Type: "None",
			Cron: "",
		},
	})

	require.NoError(t, err)
	require.NotEmpty(t, schedule)
}
