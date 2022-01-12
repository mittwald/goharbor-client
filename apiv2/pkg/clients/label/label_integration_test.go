//go:build integration

package label

import (
	"context"
	"testing"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/stretchr/testify/require"
)

func TestAPICreateLabel(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.CreateLabel(ctx, &model.Label{
		Color:       "#0072bc",
		Description: "test",
		Name:        "label",
		Scope:       ScopeGlobal.String(),
	})

	require.NoError(t, err)

	labels, err := c.ListLabels(ctx, "label", nil, ScopeGlobal)
	require.NoError(t, err)
	require.Equal(t, 1, len(labels))

	defer c.DeleteLabel(ctx, labels[0].ID)

	require.NoError(t, err)
}

func TestAPIUpdateLabel(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.CreateLabel(ctx, &model.Label{
		Color:       "#0072bc",
		Description: "foo",
		Name:        "label",
		Scope:       ScopeGlobal.String(),
	})

	require.NoError(t, err)

	labels, err := c.ListLabels(ctx, "label", nil, ScopeGlobal)
	require.NoError(t, err)
	require.Equal(t, 1, len(labels))

	update := &model.Label{
		Color:       "#72bf44",
		Description: "bar",
		ID:          labels[0].ID,
		Name:        labels[0].Name,
	}

	err = c.UpdateLabel(ctx, labels[0].ID, update)
	require.NoError(t, err)

	defer c.DeleteLabel(ctx, labels[0].ID)

	require.NoError(t, err)
}
