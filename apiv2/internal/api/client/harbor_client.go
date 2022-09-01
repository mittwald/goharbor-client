// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/artifact"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/auditlog"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/configure"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/gc"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/health"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/icon"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/immutable"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/label"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/ldap"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/member"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/oidc"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/ping"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/preheat"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/project"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/project_metadata"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/quota"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/registry"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/replication"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/repository"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/retention"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/robot"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/robotv1"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/scan"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/scan_all"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/scanner"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/search"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/statistic"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/system_cve_allowlist"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/systeminfo"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/user"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/usergroup"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/webhook"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/webhookjob"
)

// Default harbor HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v2.0"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new harbor HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Harbor {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new harbor HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Harbor {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new harbor client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Harbor {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Harbor)
	cli.Transport = transport
	cli.Artifact = artifact.New(transport, formats)
	cli.Auditlog = auditlog.New(transport, formats)
	cli.Configure = configure.New(transport, formats)
	cli.GC = gc.New(transport, formats)
	cli.Health = health.New(transport, formats)
	cli.Icon = icon.New(transport, formats)
	cli.Immutable = immutable.New(transport, formats)
	cli.Label = label.New(transport, formats)
	cli.Ldap = ldap.New(transport, formats)
	cli.Member = member.New(transport, formats)
	cli.OIDC = oidc.New(transport, formats)
	cli.Ping = ping.New(transport, formats)
	cli.Preheat = preheat.New(transport, formats)
	cli.Project = project.New(transport, formats)
	cli.ProjectMetadata = project_metadata.New(transport, formats)
	cli.Quota = quota.New(transport, formats)
	cli.Registry = registry.New(transport, formats)
	cli.Replication = replication.New(transport, formats)
	cli.Repository = repository.New(transport, formats)
	cli.Retention = retention.New(transport, formats)
	cli.Robot = robot.New(transport, formats)
	cli.Robotv1 = robotv1.New(transport, formats)
	cli.Scan = scan.New(transport, formats)
	cli.ScanAll = scan_all.New(transport, formats)
	cli.Scanner = scanner.New(transport, formats)
	cli.Search = search.New(transport, formats)
	cli.Statistic = statistic.New(transport, formats)
	cli.SystemCVEAllowlist = system_cve_allowlist.New(transport, formats)
	cli.Systeminfo = systeminfo.New(transport, formats)
	cli.User = user.New(transport, formats)
	cli.Usergroup = usergroup.New(transport, formats)
	cli.Webhook = webhook.New(transport, formats)
	cli.Webhookjob = webhookjob.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Harbor is a client for harbor
type Harbor struct {
	Artifact artifact.ClientService

	Auditlog auditlog.ClientService

	Configure configure.ClientService

	GC gc.ClientService

	Health health.ClientService

	Icon icon.ClientService

	Immutable immutable.ClientService

	Label label.ClientService

	Ldap ldap.ClientService

	Member member.ClientService

	OIDC oidc.ClientService

	Ping ping.ClientService

	Preheat preheat.ClientService

	Project project.ClientService

	ProjectMetadata project_metadata.ClientService

	Quota quota.ClientService

	Registry registry.ClientService

	Replication replication.ClientService

	Repository repository.ClientService

	Retention retention.ClientService

	Robot robot.ClientService

	Robotv1 robotv1.ClientService

	Scan scan.ClientService

	ScanAll scan_all.ClientService

	Scanner scanner.ClientService

	Search search.ClientService

	Statistic statistic.ClientService

	SystemCVEAllowlist system_cve_allowlist.ClientService

	Systeminfo systeminfo.ClientService

	User user.ClientService

	Usergroup usergroup.ClientService

	Webhook webhook.ClientService

	Webhookjob webhookjob.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Harbor) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Artifact.SetTransport(transport)
	c.Auditlog.SetTransport(transport)
	c.Configure.SetTransport(transport)
	c.GC.SetTransport(transport)
	c.Health.SetTransport(transport)
	c.Icon.SetTransport(transport)
	c.Immutable.SetTransport(transport)
	c.Label.SetTransport(transport)
	c.Ldap.SetTransport(transport)
	c.Member.SetTransport(transport)
	c.OIDC.SetTransport(transport)
	c.Ping.SetTransport(transport)
	c.Preheat.SetTransport(transport)
	c.Project.SetTransport(transport)
	c.ProjectMetadata.SetTransport(transport)
	c.Quota.SetTransport(transport)
	c.Registry.SetTransport(transport)
	c.Replication.SetTransport(transport)
	c.Repository.SetTransport(transport)
	c.Retention.SetTransport(transport)
	c.Robot.SetTransport(transport)
	c.Robotv1.SetTransport(transport)
	c.Scan.SetTransport(transport)
	c.ScanAll.SetTransport(transport)
	c.Scanner.SetTransport(transport)
	c.Search.SetTransport(transport)
	c.Statistic.SetTransport(transport)
	c.SystemCVEAllowlist.SetTransport(transport)
	c.Systeminfo.SetTransport(transport)
	c.User.SetTransport(transport)
	c.Usergroup.SetTransport(transport)
	c.Webhook.SetTransport(transport)
	c.Webhookjob.SetTransport(transport)
}
