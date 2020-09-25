// +build !integration

package retention

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-openapi/runtime"
	runtimeclient "github.com/go-openapi/runtime/client"
	v2client "github.com/mittwald/goharbor-client/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/apiv2/internal/legacyapi/client"
	"github.com/mittwald/goharbor-client/apiv2/internal/legacyapi/client/products"
	"github.com/mittwald/goharbor-client/apiv2/mocks"
	model "github.com/mittwald/goharbor-client/apiv2/model/legacy"
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
		Policy: &model.RetentionPolicy{
			Algorithm: AlgorithmOr,
			Rules: []*model.RetentionRule{{
				Action:   "retain",
				Disabled: false,
				Params: map[string]interface{}{
					PolicyTemplateDaysSinceLastPush.String(): 1,
				},
				ScopeSelectors: map[string][]model.RetentionSelector{
					"repository": {{
						Decoration: ScopeSelectorRepoMatches.String(),
						Kind:       SelectorTypeDefault,
						Pattern:    "**",
						Extras:     "", // The "Extras" field is unused for scope selectors.
					}},
				},
				TagSelectors: []*model.RetentionSelector{{
					Decoration: TagSelectorMatches.String(),
					Extras:     ToTagSelectorExtras(true),
					Kind:       SelectorTypeDefault,
					Pattern:    "**",
				}},
				Template: PolicyTemplateDaysSinceLastPush.String(),
			}},
			Scope: &model.RetentionPolicyScope{
				Level: "project",
				Ref:   0,
			},
			Trigger: &model.RetentionRuleTrigger{
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
}
