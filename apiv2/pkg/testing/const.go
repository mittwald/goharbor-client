package testing

import (
	"net/url"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	v2client "github.com/testwill/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/testwill/goharbor-client/v5/apiv2/mocks"
	"github.com/testwill/goharbor-client/v5/apiv2/pkg/config"
)

const (
	Host     = "http://core.harbor.domain:80/api/v2.0"
	User     = "admin"
	Password = "Harbor12345"
)

var (
	u, _            = url.Parse(Host)
	V2SwaggerClient = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	AuthInfo        = runtimeclient.BasicAuth(User, Password)
	DefaultOpts     = config.Defaults()
)

type MockClients struct {
	Artifact           mocks.MockArtifactClientService
	User               mocks.MockUserClientService
	Auditlog           mocks.MockAuditlogClientService
	Configure          mocks.MockConfigureClientService
	GC                 mocks.MockGcClientService
	Health             mocks.MockHealthClientService
	Icon               mocks.MockIconClientService
	Immutable          mocks.MockImmutableClientService
	Label              mocks.MockLabelClientService
	Ldap               mocks.MockLdapClientService
	Member             mocks.MockMemberClientService
	OIDC               mocks.MockOidcClientService
	Ping               mocks.MockPingClientService
	Preheat            mocks.MockPreheatClientService
	Project            mocks.MockProjectClientService
	ProjectMetadata    mocks.MockProject_metadataClientService
	Quota              mocks.MockQuotaClientService
	Registry           mocks.MockRegistryClientService
	Replication        mocks.MockReplicationClientService
	Repository         mocks.MockRepositoryClientService
	Retention          mocks.MockRetentionClientService
	Robot              mocks.MockRobotClientService
	Robotv1            mocks.MockRobotv1ClientService
	Scan               mocks.MockScanClientService
	ScanAll            mocks.MockScan_allClientService
	Scanner            mocks.MockScannerClientService
	Search             mocks.MockSearchClientService
	Statistic          mocks.MockStatisticClientService
	SystemCVEAllowlist mocks.MockSystem_cve_allowlistClientService
	Systeminfo         mocks.MockSysteminfoClientService
	Usergroup          mocks.MockUsergroupClientService
	Webhook            mocks.MockWebhookClientService
	Webhookjob         mocks.MockWebhookjobClientService
}

func BuildV2ClientWithMocks(m *MockClients) *v2client.Harbor {
	return &v2client.Harbor{
		User:               &m.User,
		Artifact:           &m.Artifact,
		Auditlog:           &m.Auditlog,
		Configure:          &m.Configure,
		GC:                 &m.GC,
		Health:             &m.Health,
		Icon:               &m.Icon,
		Immutable:          &m.Immutable,
		Label:              &m.Label,
		Ldap:               &m.Ldap,
		Member:             &m.Member,
		OIDC:               &m.OIDC,
		Ping:               &m.Ping,
		Preheat:            &m.Preheat,
		Project:            &m.Project,
		ProjectMetadata:    &m.ProjectMetadata,
		Quota:              &m.Quota,
		Registry:           &m.Registry,
		Replication:        &m.Replication,
		Repository:         &m.Repository,
		Retention:          &m.Retention,
		Robot:              &m.Robot,
		Robotv1:            &m.Robotv1,
		Scan:               &m.Scan,
		ScanAll:            &m.ScanAll,
		Scanner:            &m.Scanner,
		Search:             &m.Search,
		Statistic:          &m.Statistic,
		SystemCVEAllowlist: &m.SystemCVEAllowlist,
		Systeminfo:         &m.Systeminfo,
		Usergroup:          &m.Usergroup,
		Webhook:            &m.Webhook,
		Webhookjob:         &m.Webhookjob,
	}
}
