// +build !integration

package retention

import (
	"context"
	"net/http"
	"strconv"
	"testing"

	"github.com/go-openapi/runtime"
	runtimeclient "github.com/go-openapi/runtime/client"
	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"
	projectapi "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client/project"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client/retention"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/mocks"
	modelv2 "github.com/mittwald/goharbor-client/v3/apiv2/model"
	projectsv2 "github.com/mittwald/goharbor-client/v3/apiv2/project"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

var authInfo = runtimeclient.BasicAuth("foo", "bar")

func BuildLegacyClientWithMock(service *mocks.MockProductsClientService) *client.Harbor {
	return &client.Harbor{
		Products: service,
	}
}

func BuildV2ClientWithMocks(p *mocks.MockProjectClientService, r *mocks.MockRetentionClientService) *v2client.Harbor {
	return &v2client.Harbor{
		Artifact:   &mocks.MockArtifactClientService{},
		Auditlog:   &mocks.MockAuditlogClientService{},
		Icon:       &mocks.MockIconClientService{},
		Preheat:    &mocks.MockPreheatClientService{},
		Project:    p,
		Repository: &mocks.MockRepositoryClientService{},
		Retention:  r,
		Scan:       &mocks.MockScanClientService{},
	}
}

func BuildProjectClientWithMocks(project *mocks.MockProjectClientService) *v2client.Harbor {
	return &v2client.Harbor{
		Project: project,
	}
}

func TestNewClient(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	pr := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(pr, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	require.NotNil(t, cl)
	assert.NotNil(t, cl.AuthInfo)
	assert.NotNil(t, cl.V2Client)
	assert.NotNil(t, cl.LegacyClient)
}

func TestEvaluateRetentionRuleParams(t *testing.T) {
	t.Run("WithParams", func(t *testing.T) {
		params := map[PolicyTemplate]interface{}{
			PolicyTemplateLatestPushedArtifacts: 1,
			PolicyTemplateLatestPulledArtifacts: 2,
			PolicyTemplateDaysSinceLastPush:     3,
			PolicyTemplateDaysSinceLastPull:     4,
		}
		e, err := evaluateRetentionRuleParams(params)
		assert.NoError(t, err)
		assert.NotNil(t, e)
	})

	t.Run("WithoutParams", func(t *testing.T) {
		params := map[PolicyTemplate]interface{}{}

		e, err := evaluateRetentionRuleParams(params)

		if assert.Error(t, err) {
			assert.Nil(t, e)
		}
	})

	t.Run("InvalidParams", func(t *testing.T) {
		params := map[PolicyTemplate]interface{}{
			"foo": "bar",
		}

		e, err := evaluateRetentionRuleParams(params)

		if assert.Error(t, err) {
			assert.Nil(t, e)
		}
	})
}

func TestRESTClient_NewRetentionPolicy(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	r := &mocks.MockRetentionClientService{}
	pr := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(pr, r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postRetentionParams := &retention.CreateRetentionParams{
		Policy: &modelv2.RetentionPolicy{
			Algorithm: AlgorithmOr,
			Rules: []*modelv2.RetentionRule{{
				Action:   "retain",
				Disabled: false,
				Params: map[string]interface{}{
					PolicyTemplateDaysSinceLastPush.String(): 1,
				},
				ScopeSelectors: map[string][]modelv2.RetentionSelector{
					"repository": {{
						Decoration: ScopeSelectorRepoMatches.String(),
						Kind:       SelectorTypeDefault,
						Pattern:    "**",
					}},
				},
				TagSelectors: []*modelv2.RetentionSelector{{
					Decoration: TagSelectorMatches.String(),
					Kind:       SelectorTypeDefault,
					Pattern:    "**",
				}},
				Template: PolicyTemplateDaysSinceLastPush.String(),
			}},
			Scope: &modelv2.RetentionPolicyScope{
				Level: "project",
				Ref:   0,
			},
			Trigger: &modelv2.RetentionRuleTrigger{
				Kind:     "Schedule", // Trigger kind is _always_ 'Schedule'.
				Settings: map[string]interface{}{"cron": "0 * * * *"},
			},
		},
		Context: ctx,
	}

	r.On("CreateRetention", postRetentionParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&retention.CreateRetentionCreated{}, &runtime.APIError{Code: http.StatusCreated})

	err := cl.NewRetentionPolicy(ctx, postRetentionParams.Policy)

	assert.NoError(t, err)
	r.AssertExpectations(t)
}

func TestRESTClient_GetRetentionPolicy(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	pc := &mocks.MockProjectClientService{}
	r := &mocks.MockRetentionClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(pc, r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	retentionIDPtr := "1"

	project := &modelv2.Project{
		Deleted: false,
		Metadata: &modelv2.ProjectMetadata{
			RetentionID: &retentionIDPtr,
		},
		Name:      "test-project",
		ProjectID: 1,
	}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectNameOrID: strconv.Itoa(1),
		Context:         ctx,
	}
	getRetentionParams := &retention.GetRetentionParams{
		ID:      1,
		Context: ctx,
	}

	pc.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: project}, nil)

	r.On("GetRetention", getRetentionParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&retention.GetRetentionOK{}, nil)

	_, err := cl.GetRetentionPolicyByProject(ctx, project)

	assert.NoError(t, err)
}

func TestRESTClient_GetRetentionPolicy_ErrProjectNotFound(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	pc := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildProjectClientWithMocks(pc)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	retentionIDPtr := "1"

	project := &modelv2.Project{
		Deleted: false,
		Metadata: &modelv2.ProjectMetadata{
			RetentionID: &retentionIDPtr,
		},
		Name:      "test-project",
		ProjectID: 1,
	}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectNameOrID: strconv.Itoa(1),
		Context:         ctx,
	}

	pc.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &projectsv2.ErrProjectNotFound{})

	_, err := cl.GetRetentionPolicyByProject(ctx, project)

	if assert.Error(t, err) {
		assert.IsType(t, &projectsv2.ErrProjectNotFound{}, err)
	}
}

func TestRESTClient_GetRetentionPolicy_ErrRetentionUnauthorized(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	pc := &mocks.MockProjectClientService{}
	r := &mocks.MockRetentionClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(pc, r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	retentionIDPtr := "10"

	project := &modelv2.Project{
		Deleted: false,
		Metadata: &modelv2.ProjectMetadata{
			RetentionID: &retentionIDPtr,
		},
		Name:      "test-project",
		ProjectID: 1,
	}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectNameOrID: strconv.Itoa(1),
		Context:         ctx,
	}

	getRetentionParams := &retention.GetRetentionParams{
		ID:      10,
		Context: ctx,
	}

	pc.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: project}, nil)

	r.On("GetRetention", getRetentionParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusUnauthorized})

	_, err := cl.GetRetentionPolicyByProject(ctx, project)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrRetentionUnauthorized{}, err)
	}

	r.AssertExpectations(t)
}

func TestRESTClient_UpdateRetentionPolicy(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	pr := &mocks.MockProjectClientService{}
	r := &mocks.MockRetentionClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(pr, r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	policy := &modelv2.RetentionPolicy{
		Algorithm: "",
		ID:        1,
		Rules:     nil,
		Scope:     nil,
		Trigger:   nil,
	}

	putRetentionParams := &retention.UpdateRetentionParams{
		ID:      1,
		Policy:  policy,
		Context: ctx,
	}

	r.On("UpdateRetention", putRetentionParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&retention.UpdateRetentionOK{}, &runtime.APIError{Code: http.StatusOK})

	err := cl.UpdateRetentionPolicy(ctx, putRetentionParams.Policy)

	assert.NoError(t, err)
	r.AssertExpectations(t)
}

func TestRESTClient_UpdateRetentionPolicy_ErrRetentionNotProvided(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	pr := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(pr, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()
	putRetentionParams := &retention.UpdateRetentionParams{}

	err := cl.UpdateRetentionPolicy(ctx, putRetentionParams.Policy)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrRetentionNotProvided{}, err)
	}
}

func TestRESTClient_UpdateRetentionPolicy_ErrRetentionDoesNotExist(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	pr := &mocks.MockProjectClientService{}
	r := &mocks.MockRetentionClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(pr, r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	policy := &modelv2.RetentionPolicy{
		Algorithm: "",
		ID:        1,
	}

	putRetentionParams := &retention.UpdateRetentionParams{
		ID:      1,
		Policy:  policy,
		Context: ctx,
	}

	r.On("UpdateRetention", putRetentionParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusOK})

	err := cl.UpdateRetentionPolicy(ctx, putRetentionParams.Policy)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrRetentionDoesNotExist{}, err)
	}

	r.AssertExpectations(t)
}

func TestRESTClient_DisableRetentionPolicy(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	pr := &mocks.MockProjectClientService{}
	r := &mocks.MockRetentionClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(pr, r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	policy := &modelv2.RetentionPolicy{
		Algorithm: "",
		ID:        1,
		Rules:     []*modelv2.RetentionRule{},
	}

	putRetentionParams := &retention.UpdateRetentionParams{
		ID:      1,
		Policy:  policy,
		Context: ctx,
	}

	r.On("UpdateRetention", putRetentionParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&retention.UpdateRetentionOK{}, &runtime.APIError{Code: http.StatusOK})

	err := cl.DisableRetentionPolicy(ctx, policy)

	assert.NoError(t, err)

	r.AssertExpectations(t)
}

func TestRESTClient_DisableRetentionPolicy_ErrRetentionDoesNotExist(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	pr := &mocks.MockProjectClientService{}
	r := &mocks.MockRetentionClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(pr, r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	policy := &modelv2.RetentionPolicy{
		Algorithm: "",
		ID:        1,
		Rules:     []*modelv2.RetentionRule{},
	}

	putRetentionParams := &retention.UpdateRetentionParams{
		ID:      1,
		Policy:  policy,
		Context: ctx,
	}

	r.On("UpdateRetention", putRetentionParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, nil)

	err := cl.DisableRetentionPolicy(ctx, policy)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrRetentionDoesNotExist{}, err)
	}

	r.AssertExpectations(t)
}

func TestErrRetentionUnauthorized_Error(t *testing.T) {
	var e ErrRetentionUnauthorized

	assert.Equal(t, ErrRetentionUnauthorizedMsg, e.Error())
}

func TestErrRetentionNotProvided_Error(t *testing.T) {
	var e ErrRetentionNotProvided

	assert.Equal(t, ErrRetentionNotProvidedMsg, e.Error())
}

func TestErrRetentionNoPermission_Error(t *testing.T) {
	var e ErrRetentionNoPermission

	assert.Equal(t, ErrRetentionNoPermissionMsg, e.Error())
}

func TestErrRetentionDoesNotExist_Error(t *testing.T) {
	var e ErrRetentionDoesNotExist

	assert.Equal(t, ErrRetentionDoesNotExistMsg, e.Error())
}

func TestErrRetentionInternalErrors_Error(t *testing.T) {
	var e ErrRetentionInternalErrors

	assert.Equal(t, ErrRetentionInternalErrorsMsg, e.Error())
}
