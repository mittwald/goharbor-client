package retention

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-openapi/runtime"

	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/retention"
	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/projectmeta"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/common"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
)

const (
	// AlgorithmOr is the default algorithm when operating on harbor retention rules
	AlgorithmOr string = "or"

	// Key for defining matching repositories
	ScopeSelectorRepoMatches ScopeSelector = "repoMatches"

	// Key for defining excluded repositories
	ScopeSelectorRepoExcludes ScopeSelector = "repoExcludes"

	// Key for defining matching tag expressions
	TagSelectorMatches TagSelector = "matches"

	// Key for defining excluded tag expressions
	TagSelectorExcludes TagSelector = "excludes"

	// The kind of the retention selector, _always_ defaults to 'doublestar'
	SelectorTypeDefault string = "doublestar"

	// Retain the most recently pushed n artifacts - count
	PolicyTemplateLatestPushedArtifacts PolicyTemplate = "latestPushedK"

	// Retain the most recently pulled n artifacts - count
	PolicyTemplateLatestPulledArtifacts PolicyTemplate = "latestPulledN"

	// Retain the artifacts pushed within the last n days
	PolicyTemplateDaysSinceLastPush PolicyTemplate = "nDaysSinceLastPush"

	// Retain the artifacts pulled within the last n days
	PolicyTemplateDaysSinceLastPull PolicyTemplate = "nDaysSinceLastPull"

	// Retain always
	PolicyTemplateRetainAlways PolicyTemplate = "always"
)

type Client interface {
	NewRetentionPolicy(ctx context.Context, ret *modelv2.RetentionPolicy) error
	GetRetentionPolicyByProject(ctx context.Context, projectNameOrID string) (*modelv2.RetentionPolicy, error)
	GetRetentionPolicyByID(ctx context.Context, id int64) (*modelv2.RetentionPolicy, error)
	DeleteRetentionPolicyByID(ctx context.Context, id int64) error
	UpdateRetentionPolicy(ctx context.Context, ret *modelv2.RetentionPolicy) error
}

// RESTClient is a subclient for handling retention related actions.
type RESTClient struct {
	// Options contains optional configuration when making API calls.
	Options *config.Options

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

// ScopeSelector is the retention selector decoration used for operations on retention objects.
type ScopeSelector string

func (r ScopeSelector) String() string {
	return string(r)
}

// PolicyTemplate defines the possible values used for the policy matching mechanism.
type PolicyTemplate string

func (p PolicyTemplate) String() string {
	return string(p)
}

// TagSelector defines the possible values used for the tag matching mechanism. Valid values are: "matches, excludes".
type TagSelector string

// String returns the string value of a TagSelector.
func (t TagSelector) String() string {
	return string(t)
}

// NewRetentionPolicy creates a new tag retention policy for a project.
func (c *RESTClient) NewRetentionPolicy(ctx context.Context, ret *modelv2.RetentionPolicy) error {
	if ret == nil {
		return &ErrRetentionNotProvided{}
	}

	params := &retention.CreateRetentionParams{
		Policy:  ret,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Retention.CreateRetention(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerRetentionErrors(err)
	}

	return nil
}

// GetRetentionPolicyByProject returns the retention policy associated to a project.
func (c *RESTClient) GetRetentionPolicyByProject(ctx context.Context, projectNameOrID string) (*modelv2.RetentionPolicy, error) {
	pm := projectmeta.NewClient(c.V2Client, c.Options, c.AuthInfo)

	val, err := pm.GetProjectMetadataValue(ctx, projectNameOrID, common.ProjectMetadataKeyRetentionID)
	if err != nil {
		return nil, err
	}

	id, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("could not convert retention id %q to int64, project: %s", val, projectNameOrID)
	}

	return c.GetRetentionPolicyByID(ctx, id)
}

// GetRetentionPolicyByID returns a retention policy identified by it's id.
func (c *RESTClient) GetRetentionPolicyByID(ctx context.Context, id int64) (*modelv2.RetentionPolicy, error) {
	params := &retention.GetRetentionParams{
		ID:      id,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Retention.GetRetention(params, c.AuthInfo)
	if err != nil {
		fmt.Println()
		return nil, handleSwaggerRetentionErrors(err)
	}

	return resp.Payload, nil
}

func (c *RESTClient) DeleteRetentionPolicyByID(ctx context.Context, id int64) error {
	params := &retention.DeleteRetentionParams{
		ID:      id,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Retention.DeleteRetention(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerRetentionErrors(err)
	}

	return nil
}

// UpdateRetentionPolicy updates the specified retention policy ret.
func (c *RESTClient) UpdateRetentionPolicy(ctx context.Context, ret *modelv2.RetentionPolicy) error {
	if ret == nil {
		return &ErrRetentionNotProvided{}
	}

	params := &retention.UpdateRetentionParams{
		ID:      ret.ID,
		Policy:  ret,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Retention.UpdateRetention(params, c.AuthInfo)

	if resp == nil {
		return &ErrRetentionDoesNotExist{}
	}

	if err != nil {
		return handleSwaggerRetentionErrors(err)
	}

	return nil
}

// ToTagSelectorExtras converts a boolean to the representative string value used by Harbor.
// Represents the functionality of the 'untagged artifacts' checkbox when editing tag retention rules in the Harbor UI.
func ToTagSelectorExtras(untagged bool) string {
	if untagged {
		return `{"untagged":"true"}`
	}
	return `{"untagged":"false"}`
}

// evaluateRetentionRuleParams evaluates the provided map of PolicyTemplate by comparing the keys to the pre-defined PolicyTemplates.
// Returns an error if the provided or resulting map is empty.
func evaluateRetentionRuleParams(params map[PolicyTemplate]interface{}) (map[string]interface{}, error) {
	res := make(map[string]interface{})

	if len(params) == 0 {
		return nil, errors.New("no retention rule params provided")
	}

	for k, v := range params {
		if _, ok := params[k]; ok {
			switch k {
			case PolicyTemplateDaysSinceLastPull:
				res[k.String()] = v
			case PolicyTemplateDaysSinceLastPush:
				res[k.String()] = v
			case PolicyTemplateLatestPulledArtifacts:
				res[k.String()] = v
			case PolicyTemplateLatestPushedArtifacts:
				res[k.String()] = v
			default:
				continue
			}
		}
	}

	if len(res) == 0 {
		return nil, errors.New("invalid params provided")
	}

	return res, nil
}
