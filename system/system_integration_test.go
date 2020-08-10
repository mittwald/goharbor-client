// +build integration

package system

import (
	"context"
	"flag"
	"github.com/go-openapi/strfmt"
	"github.com/mittwald/goharbor-client/internal/api/v1_10_0/client"
	"net/url"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	model "github.com/mittwald/goharbor-client/model/v1_10_0"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	host     = "http://localhost:30002/api"
	user     = "admin"
	password = "Harbor12345"
)

var (
	u, _          = url.Parse(host)
	swaggerClient = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo      = runtimeclient.BasicAuth(user, password)
	harborVersion = flag.String("version", "1.10.2",
		"Harbor version, used in conjunction with -integration, "+
			"defaults to 1.10.2")
	skipSpinUp = flag.Bool("skip-spinup", false,
		"Skip kind cluster creation")
)

// TestAPISystemGcScheduleNew tests the creation of a new GC schedule
// and resets it to the default schedule afterwards
func TestAPISystemGcScheduleNew(t *testing.T) {
	cron := "0 * * * *"
	scheduleType := "Hourly"

	ctx := context.Background()
	c := NewClient(swaggerClient, authInfo)

	_, err := c.GetSystemGarbageCollection(ctx)
	require.IsType(t, &ErrSystemGcUndefined{}, err)

	gcSchedule, err := c.NewSystemGarbageCollection(ctx, cron, scheduleType)
	require.NoError(t, err)

	defer c.ResetSystemGarbageCollection(ctx)

	assert.Equal(t, gcSchedule.Schedule.Cron, cron)
	assert.Equal(t, gcSchedule.Schedule.Type, scheduleType)
}

// TestAPISystemGcScheduleUpdate tests the update of an existing GC schedule,
// asserting the updated schedule cron matches the given values
func TestAPISystemGcScheduleUpdate(t *testing.T) {
	cron := "0 * * * *"
	scheduleType := "Hourly"

	ctx := context.Background()
	c := NewClient(swaggerClient, authInfo)

	_, err := c.GetSystemGarbageCollection(ctx)
	require.IsType(t, &ErrSystemGcUndefined{}, err)

	_, err = c.NewSystemGarbageCollection(ctx, cron, scheduleType)
	require.NoError(t, err)

	cron2 := "* * */1 * *"
	scheduleType2 := "Daily"
	err = c.UpdateSystemGarbageCollection(ctx, &model.AdminJobScheduleObj{
		Cron: cron2,
		Type: scheduleType2,
	})

	gcSchedule, err := c.GetSystemGarbageCollection(ctx)
	require.NoError(t, err)

	assert.Equal(t, gcSchedule.Schedule.Cron, cron2)
	assert.Equal(t, gcSchedule.Schedule.Type, scheduleType2)

	err = c.ResetSystemGarbageCollection(ctx)
	require.NoError(t, err)
}

// TestAPISystemGcScheduleReset tests the reset of an existing GC schedule
func TestAPISystemGcScheduleReset(t *testing.T) {
	cron := "0 * * * *"
	scheduleType := "Hourly"

	ctx := context.Background()
	c := NewClient(swaggerClient, authInfo)

	_, err := c.GetSystemGarbageCollection(ctx)
	require.IsType(t, &ErrSystemGcUndefined{}, err)

	_, err = c.NewSystemGarbageCollection(ctx, cron, scheduleType)
	require.NoError(t, err)

	err = c.ResetSystemGarbageCollection(ctx)
	require.NoError(t, err)

	_, err = c.GetSystemGarbageCollection(ctx)
	require.IsType(t, &ErrSystemGcUndefined{}, err)
}

func TestAPIHealth(t *testing.T) {
	ctx := context.Background()
	c := NewClient(swaggerClient, authInfo)

	_, err := c.Health(ctx)
	require.NoError(t, err)
}
