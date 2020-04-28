package harbor

import "time"

// To ensure type-safe queries to the harbor API,
// the following typings include typings from the upstream sources:
// https://github.com/goharbor/harbor/src/replication/dao/models/

// const definition
const (
	RegistryTypeHarbor           RegistryType = "harbor"
	RegistryTypeDockerHub        RegistryType = "docker-hub"
	RegistryTypeDockerRegistry   RegistryType = "docker-registry"
	RegistryTypeHuawei           RegistryType = "huawei-SWR"
	RegistryTypeGoogleGcr        RegistryType = "google-gcr"
	RegistryTypeAwsEcr           RegistryType = "aws-ecr"
	RegistryTypeAzureAcr         RegistryType = "azure-acr"
	RegistryTypeAliAcr           RegistryType = "ali-acr"
	RegistryTypeJfrogArtifactory RegistryType = "jfrog-artifactory"
	RegistryTypeQuayio           RegistryType = "quay-io"
	RegistryTypeGitLab           RegistryType = "gitlab"

	RegistryTypeHelmHub RegistryType = "helm-hub"

	FilterStyleTypeText  = "input"
	FilterStyleTypeRadio = "radio"
	FilterStyleTypeList  = "list"
)

// RegistryType indicates the type of registry
type RegistryType string

// CredentialType represents the supported credential types
// e.g: u/p, OAuth token
type CredentialType string

// Credential keeps the access key and/or secret for the related registry
type Credential struct {
	// Type of the credential
	Type CredentialType `json:"type"`
	// The key of the access account, for OAuth token, it can be empty
	AccessKey string `json:"access_key"`
	// The secret or password for the key
	AccessSecret string `json:"access_secret"`
}

// Registry keeps the related info of registry
// Data required for the secure access way is not contained here.
// DAO layer is not considered here
type Registry struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Type        RegistryType `json:"type"`
	URL         string       `json:"url"`
	// TokenServiceURL is only used for local harbor instance to
	// avoid the requests passing through the external proxy for now
	TokenServiceURL string      `json:"token_service_url"`
	Credential      *Credential `json:"credential"`
	Insecure        bool        `json:"insecure"`
	Status          string      `json:"status"`
	CreationTime    time.Time   `json:"creation_time"`
	UpdateTime      time.Time   `json:"update_time"`
}

// RegistryQuery defines the query conditions for listing registries
type RegistryQuery struct {
	// Name is name of the registry to query
	Name string
	// Pagination specifies the pagination
	Pagination *Pagination
}

// FilterType represents the type info of the filter.
type FilterType string

// FilterStyle ...
type FilterStyle struct {
	Type   FilterType `json:"type"`
	Style  string     `json:"style"`
	Values []string   `json:"values,omitempty"`
}

// EndpointPattern ...
type EndpointPattern struct {
	EndpointType EndpointType `json:"endpoint_type"`
	Endpoints    []*Endpoint  `json:"endpoints"`
}

// EndpointType ..
type EndpointType string

// Endpoint ...
type Endpoint struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// RegistryInfo provides base info and capability declarations of the registry
type RegistryInfo struct {
	Type                     RegistryType   `json:"type"`
	Description              string         `json:"description"`
	SupportedResourceTypes   []ResourceType `json:"-"`
	SupportedResourceFilters []*FilterStyle `json:"supported_resource_filters"`
	SupportedTriggers        []TriggerType  `json:"supported_triggers"`
}

// TriggerType represents the type of trigger.
type TriggerType string

// ResourceType represents the type of the resource
type ResourceType string

// ResourceMetadata of resource
type ResourceMetadata struct {
	Repository *Repository `json:"repository"`
	Vtags      []string    `json:"v_tags"`
	Labels     []string    `json:"labels"`
}

// Repository info of the resource
type Repository struct {
	Name     string                 `json:"name"`
	Metadata map[string]interface{} `json:"metadata"`
}

// Resource represents the general replicating content
type Resource struct {
	Type         ResourceType           `json:"type"`
	Metadata     *ResourceMetadata      `json:"metadata"`
	Registry     *Registry              `json:"registry"`
	ExtendedInfo map[string]interface{} `json:"extended_info"`
	// Indicate if the resource is a deleted resource
	Deleted bool `json:"deleted"`
	// indicate whether the resource can be overridden
	Override bool `json:"override"`
}
