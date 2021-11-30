package common

type MetadataKey string

const (
	ProjectMetadataKeyEnableContentTrust   MetadataKey = "enable_content_trust"
	ProjectMetadataKeyAutoScan             MetadataKey = "auto_scan"
	ProjectMetadataKeySeverity             MetadataKey = "severity"
	ProjectMetadataKeyReuseSysCVEAllowlist MetadataKey = "reuse_sys_cve_allowlist"
	ProjectMetadataKeyPublic               MetadataKey = "public"
	ProjectMetadataKeyPreventVul           MetadataKey = "prevent_vul"
	ProjectMetadataKeyRetentionID          MetadataKey = "retention_id"
)

func (m MetadataKey) String() string {
	return string(m)
}
