// +build !integration

package retention

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-openapi/runtime"
	runtimeclient "github.com/go-openapi/runtime/client"
	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"
	projectapi "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client/project"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client/products"
	"github.com/mittwald/goharbor-client/v3/apiv2/mocks"
	modelv2 "github.com/mittwald/goharbor-client/v3/apiv2/model"
	legacymodel "github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"
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

func BuildV2ClientWithMocks() *v2client.Harbor {
	return &v2client.Harbor{
		Artifact:   &mocks.MockArtifactClientService{},
		Auditlog:   &mocks.MockAuditlogClientService{},
		Icon:       &mocks.MockIconClientService{},
		Preheat:    &mocks.MockPreheatClientService{},
		Project:    &mocks.MockProjectClientService{},
		Repository: &mocks.MockRepositoryClientService{},
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

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks()

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

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks()

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postRetentionParams := &products.PostRetentionsParams{
		Policy: &legacymodel.RetentionPolicy{
			Algorithm: AlgorithmOr,
			Rules: []*legacymodel.RetentionRule{{
				Action:   "retain",
				Disabled: false,
				Params: map[string]interface{}{
					PolicyTemplateDaysSinceLastPush.String(): 1,
				},
				ScopeSelectors: map[string][]legacymodel.RetentionSelector{
					"repository": {{
						Decoration: ScopeSelectorRepoMatches.String(),
						Kind:       SelectorTypeDefault,
						Pattern:    "**",
					}},
				},
				TagSelectors: []*legacymodel.RetentionSelector{{
					Decoration: TagSelectorMatches.String(),
					Kind:       SelectorTypeDefault,
					Pattern:    "**",
				}},
				Template: PolicyTemplateDaysSinceLastPush.String(),
			}},
			Scope: &legacymodel.RetentionPolicyScope{
				Level: "project",
				Ref:   0,
			},
			Trigger: &legacymodel.RetentionRuleTrigger{
				Kind:     "Schedule", // Trigger kind is _always_ 'Schedule'.
				Settings: map[string]interface{}{"cron": "0 * * * *"},
			},
		},
		Context: ctx,
	}

	p.On("PostRetentions", postRetentionParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostRetentionsCreated{}, &runtime.APIError{Code: http.StatusCreated})

	err := cl.NewRetentionPolicy(ctx, postRetentionParams.Policy)

	assert.NoError(t, err)
	p.AssertExpectations(t)
}

func TestRESTClient_GetRetentionPolicy(t *testing.T) {
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
		ProjectID: 1,
		Context:   ctx,
	}
	getRetentionParams := &products.GetRetentionsIDParams{
		ID:      1,
		Context: ctx,
	}

	pc.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: project}, nil)

	p.On("GetRetentionsID", getRetentionParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetRetentionsIDOK{}, nil)

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
		ProjectID: 1,
		Context:   ctx,
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

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildProjectClientWithMocks(pc)

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
		ProjectID: 1,
		Context:   ctx,
	}

	getRetentionParams := &products.GetRetentionsIDParams{
		ID:      10,
		Context: ctx,
	}

	pc.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: project}, nil)

	p.On("GetRetentionsID", getRetentionParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusUnauthorized})

	_, err := cl.GetRetentionPolicyByProject(ctx, project)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrRetentionUnauthorized{}, err)
	}
}

func TestRESTClient_UpdateRetentionPolicy(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks()

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	policy := &legacymodel.RetentionPolicy{
		Algorithm: "",
		ID:        1,
		Rules:     nil,
		Scope:     nil,
		Trigger:   nil,
	}

	putRetentionParams := &products.PutRetentionsIDParams{
		ID:      1,
		Policy:  policy,
		Context: ctx,
	}

	p.On("PutRetentionsID", putRetentionParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PutRetentionsIDOK{}, &runtime.APIError{Code: http.StatusOK})

	err := cl.UpdateRetentionPolicy(ctx, putRetentionParams.Policy)

	assert.NoError(t, err)
	p.AssertExpectations(t)
}

func TestRESTClient_UpdateRetentionPolicy_ErrRetentionNotProvided(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks()

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()
	putRetentionParams := &products.PutRetentionsIDParams{}

	err := cl.UpdateRetentionPolicy(ctx, putRetentionParams.Policy)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrRetentionNotProvided{}, err)
	}
}

func TestRESTClient_UpdateRetentionPolicy_ErrRetentionDoesNotExist(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks()

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	policy := &legacymodel.RetentionPolicy{
		Algorithm: "",
		ID:        1,
	}

	putRetentionParams := &products.PutRetentionsIDParams{
		ID:      1,
		Policy:  policy,
		Context: ctx,
	}

	p.On("PutRetentionsID", putRetentionParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusOK})

	err := cl.UpdateRetentionPolicy(ctx, putRetentionParams.Policy)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrRetentionDoesNotExist{}, err)
	}
}

func TestRESTClient_DisableRetentionPolicy(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks()

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	policy := &legacymodel.RetentionPolicy{
		Algorithm: "",
		ID:        1,
		Rules:     []*legacymodel.RetentionRule{},
	}

	putRetentionParams := &products.PutRetentionsIDParams{
		ID:      1,
		Policy:  policy,
		Context: ctx,
	}

	p.On("PutRetentionsID", putRetentionParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PutRetentionsIDOK{}, &runtime.APIError{Code: http.StatusOK})

	err := cl.DisableRetentionPolicy(ctx, policy)

	assert.NoError(t, err)
	p.AssertExpectations(t)
}

func TestRESTClient_DisableRetentionPolicy_ErrRetentionDoesNotExist(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks()

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	policy := &legacymodel.RetentionPolicy{
		Algorithm: "",
		ID:        1,
		Rules:     []*legacymodel.RetentionRule{},
	}

	putRetentionParams := &products.PutRetentionsIDParams{
		ID:      1,
		Policy:  policy,
		Context: ctx,
	}

	p.On("PutRetentionsID", putRetentionParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, nil)

	err := cl.DisableRetentionPolicy(ctx, policy)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrRetentionDoesNotExist{}, err)
	}

	p.AssertExpectations(t)
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
