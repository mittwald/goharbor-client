// +build integration

package gc

import (
	"context"
	"net/url"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client"
	modelv2 "github.com/mittwald/goharbor-client/v3/apiv2/model"
	integrationtest "github.com/mittwald/goharbor-client/v3/apiv2/testing"
	"github.com/stretchr/testify/require"
)

var (
	u, _                = url.Parse(integrationtest.Host)
	legacySwaggerClient = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient     = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo            = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
)

func TestAPINewSystemGarbageCollection(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	err := c.NewSystemGarbageCollection(ctx, &modelv2.Schedule{
		Schedule: &modelv2.ScheduleObj{
			Cron: "0 * * * * *",
			Type: "Hourly",
		},
	})

	defer c.ResetSystemGarbageCollection(ctx)

	require.NoError(t, err)
}

func TestAPIUpdateSystemGarbageCollection(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	err := c.UpdateSystemGarbageCollection(ctx, &modelv2.Schedule{
		Schedule: &modelv2.ScheduleObj{
			Cron: "0 * * * * *",
			Type: "Hourly",
		},
	})

	defer c.ResetSystemGarbageCollection(ctx)

	require.NoError(t, err)
}

func TestResetSystemGarbageCollection(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	err := c.ResetSystemGarbageCollection(ctx)

	require.NoError(t, err)
}

func TestGetSystemGarbageCollectionSchedule(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	sched := &modelv2.ScheduleObj{
		Cron: "0 * * * * *",
		Type: "Hourly",
	}

	err := c.NewSystemGarbageCollection(ctx, &modelv2.Schedule{
		Schedule: sched,
	})

	require.NoError(t, err)

	gc, err := c.GetSystemGarbageCollectionSchedule(ctx)

	require.NoError(t, err)
	require.NotNil(t, gc)

	require.Equal(t, gc.Schedule, sched)

	defer c.ResetSystemGarbageCollection(ctx)
}
