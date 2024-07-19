package immutable

import (
	"context"
	immutableapi "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/immutable"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"

	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"

	"github.com/go-openapi/runtime"
)

type RESTClient struct {
	// Options contains optional configuration when making API calls.
	Options *config.Options

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(v2Client *v2client.Harbor, opts *config.Options, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Options:  opts,
		V2Client: v2Client,
		AuthInfo: authInfo,
	}
}

type Client interface {
	CreateImmuRule(ctx context.Context, projectNameOrID string, immutableRule *model.ImmutableRule) error
	UpdateImmuRule(ctx context.Context, projectNameOrID string, immutableRule *model.ImmutableRule) error
	DeleteImmuRule(ctx context.Context, projectNameOrID string, immutableRuleID int64) error
	ListImmuRules(ctx context.Context, projectNameOrID string) ([]*model.ImmutableRule, error)
}

func (c *RESTClient) CreateImmuRule(ctx context.Context, projectNameOrID string, immutableRule *model.ImmutableRule) error {
	params := &immutableapi.CreateImmuRuleParams{
		ProjectNameOrID: projectNameOrID,
		ImmutableRule:   immutableRule,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Immutable.CreateImmuRule(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerImmutableRuleErrors(err)
	}

	return nil
}

func (c *RESTClient) UpdateImmuRule(ctx context.Context, projectNameOrID string, immutableRule *model.ImmutableRule, immutableRuleID int64) error {
	params := &immutableapi.UpdateImmuRuleParams{
		ProjectNameOrID: projectNameOrID,
		ImmutableRule:   immutableRule,
		ImmutableRuleID: immutableRuleID,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	params.ImmutableRule.ID = immutableRuleID

	_, err := c.V2Client.Immutable.UpdateImmuRule(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerImmutableRuleErrors(err)
	}

	return nil
}

func (c *RESTClient) DeleteImmuRule(ctx context.Context, projectNameOrID string, immutableRuleID int64) error {
	params := &immutableapi.DeleteImmuRuleParams{
		ProjectNameOrID: projectNameOrID,
		ImmutableRuleID: immutableRuleID,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Immutable.DeleteImmuRule(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerImmutableRuleErrors(err)
	}

	return nil
}

func (c *RESTClient) ListImmuRules(ctx context.Context, projectNameOrID string) ([]*model.ImmutableRule, error) {
	var immutableRules []*model.ImmutableRule
	page := c.Options.Page

	params := &immutableapi.ListImmuRulesParams{
		Page:            &page,
		PageSize:        &c.Options.PageSize,
		ProjectNameOrID: projectNameOrID,
		Q:               &c.Options.Query,
		Sort:            &c.Options.Sort,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	for {
		resp, err := c.V2Client.Immutable.ListImmuRules(params, c.AuthInfo)
		if err != nil {
			return nil, handleSwaggerImmutableRuleErrors(err)
		}

		if len(resp.Payload) == 0 {
			break
		}

		totalCount := resp.XTotalCount

		immutableRules = append(immutableRules, resp.Payload...)

		if int64(len(immutableRules)) >= totalCount {
			break
		}

		page++
	}

	return immutableRules, nil
}
