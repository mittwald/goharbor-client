//go:build integration

package gc

import (
	"context"
	"testing"

	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"

	"github.com/stretchr/testify/require"

	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
)

func TestAPINewGarbageCollection(t *testing.T) {
	ctx := context.Background()

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.NewGarbageCollection(ctx, &modelv2.Schedule{
		Schedule: &modelv2.ScheduleObj{
			Cron: "0 * * * * *",
			Type: "Hourly",
		},
	})

	defer c.ResetGarbageCollection(ctx)

	require.NoError(t, err)
}

func TestAPIUpdateGarbageCollection(t *testing.T) {
	ctx := context.Background()

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.UpdateGarbageCollection(ctx, &modelv2.Schedule{
		Schedule: &modelv2.ScheduleObj{
			Cron: "0 * * * * *",
			Type: "Hourly",
		},
	})

	defer c.ResetGarbageCollection(ctx)

	require.NoError(t, err)
}

func TestResetGarbageCollection(t *testing.T) {
	ctx := context.Background()

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.ResetGarbageCollection(ctx)

	require.NoError(t, err)
}

func TestGetGarbageCollectionSchedule(t *testing.T) {
	ctx := context.Background()

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	sched := &modelv2.ScheduleObj{
		Cron: "0 * * * * *",
		Type: "Hourly",
	}

	err := c.NewGarbageCollection(ctx, &modelv2.Schedule{
		Schedule: sched,
	})

	require.NoError(t, err)

	gc, err := c.GetGarbageCollectionSchedule(ctx)

	require.NoError(t, err)
	require.NotNil(t, gc)

	require.Equal(t, gc.Schedule, sched)

	defer c.ResetGarbageCollection(ctx)
}
