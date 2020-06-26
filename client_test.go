package goharborclient

import (
	"context"
	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/model"
	"github.com/mittwald/goharbor-client/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRESTClient_NewProject(t *testing.T) {
	mockClient := &mocks.Client{}

	ctx := context.Background()

	mockClient.On("NewProject", ctx, "test-project", 0, 0).Return(nil).Once()

	mockClient.NewProject(ctx, "test-project", 0, 0)

	mockClient.AssertExpectations(t)
}

func TestRESTClient_ListProjects(t *testing.T) {
	mockClient := &mocks.Client{}

	ctx := context.Background()

	mockClient.On("ListProjects", ctx, "test-project").Return([]*model.Project{}, nil).Once()

	p, _ := mockClient.ListProjects(ctx, "test-project")

	assert.Len(t, p, 0)

	mockClient.AssertExpectations(t)
}
