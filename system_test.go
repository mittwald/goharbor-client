package goharborclient

import (
	"context"
	"testing"

	"github.com/mittwald/goharbor-client/api/v1.10.0/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAPISystemGcScheduleNew tests the creation of a new GC schedule
// and resets it to the default schedule afterwards
func TestAPISystemGcScheduleNew(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	cron := "0 * * * *"
	scheduleType := "Hourly"

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	_, err := c.System().GetSystemGarbageCollection(ctx)
	require.IsType(t, &ErrSystemGcUndefined{}, err)

	gcSchedule, err := c.System().NewSystemGarbageCollection(ctx, cron, scheduleType)
	require.NoError(t, err)

	defer c.System().ResetSystemGarbageCollection(ctx)

	assert.Equal(t, gcSchedule.Schedule.Cron, cron)
	assert.Equal(t, gcSchedule.Schedule.Type, scheduleType)
}

// TestAPISystemGcScheduleUpdate tests the update of an existing GC schedule,
// asserting the updated schedule cron matches the given values
func TestAPISystemGcScheduleUpdate(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	cron := "0 * * * *"
	scheduleType := "Hourly"

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	_, err := c.System().GetSystemGarbageCollection(ctx)
	require.IsType(t, &ErrSystemGcUndefined{}, err)

	_, err = c.System().NewSystemGarbageCollection(ctx, cron, scheduleType)
	require.NoError(t, err)

	cron2 := "* * */1 * *"
	scheduleType2 := "Daily"
	err = c.System().UpdateSystemGarbageCollection(ctx, &model.AdminJobScheduleObj{
		Cron: cron2,
		Type: scheduleType2,
	})

	gcSchedule, err := c.System().GetSystemGarbageCollection(ctx)
	require.NoError(t, err)

	assert.Equal(t, gcSchedule.Schedule.Cron, cron2)
	assert.Equal(t, gcSchedule.Schedule.Type, scheduleType2)

	err = c.System().ResetSystemGarbageCollection(ctx)
	require.NoError(t, err)
}

// TestAPISystemGcScheduleReset tests the reset of an existing GC schedule
func TestAPISystemGcScheduleReset(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	cron := "0 * * * *"
	scheduleType := "Hourly"

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	_, err := c.System().GetSystemGarbageCollection(ctx)
	require.IsType(t, &ErrSystemGcUndefined{}, err)

	_, err = c.System().NewSystemGarbageCollection(ctx, cron, scheduleType)
	require.NoError(t, err)

	err = c.System().ResetSystemGarbageCollection(ctx)
	require.NoError(t, err)

	_, err = c.System().GetSystemGarbageCollection(ctx)
	require.IsType(t, &ErrSystemGcUndefined{}, err)
}
