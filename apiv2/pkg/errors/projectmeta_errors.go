package errors

import "github.com/mittwald/goharbor-client/v5/apiv2/pkg/common"

const (
	// ErrProjectMetadataAlreadyExistsMsg is the error message for ErrProjectMetadataAlreadyExists error.
	ErrProjectMetadataAlreadyExistsMsg = "metadata key already exists"
	// ErrProjectMetadataUndefinedMsg is the error message for ErrProjectMetadataUndefined error.
	ErrProjectMetadataUndefinedMsg = "project metadata undefined"
	// ErrProjectMetadataValueUndefinedMsg is the error message used for MetadataKey's being undefined or nil.
	ErrProjectMetadataValueUndefinedMsg = "project metadata value is empty: "

	ErrProjectMetadataInvalidRequestMsg = "request invalid"
	ErrProjectMetadataKeyUndefinedMsg   = "metadata key is undefined"
)

type (
	ErrProjectMetadataUndefined    struct{}
	ErrProjectMetadataKeyUndefined struct{}
	// ErrProjectMetadataValueEnableContentTrustUndefined describes an error regarding a metadata value being undefined or nil.
	ErrProjectMetadataValueEnableContentTrustUndefined struct{}
	// ErrProjectMetadataValueAutoScanUndefined describes an error regarding a metadata value being undefined or nil.
	ErrProjectMetadataValueAutoScanUndefined struct{}
	// ErrProjectMetadataValueSeverityUndefined describes an error regarding a metadata value being undefined or nil.
	ErrProjectMetadataValueSeverityUndefined struct{}
	// ErrProjectMetadataValueReuseSysCveAllowlistUndefined describes an error regarding a metadata value being undefined or nil.
	ErrProjectMetadataValueReuseSysCveAllowlistUndefined struct{}
	// ErrProjectMetadataValuePublicUndefined describes an error regarding a metadata value being undefined or nil.
	ErrProjectMetadataValuePublicUndefined struct{}
	// ErrProjectMetadataValuePreventVulUndefined describes an error regarding a metadata value being undefined or nil.
	ErrProjectMetadataValuePreventVulUndefined  struct{}
	ErrProjectMetadataValueRetentionIDUndefined struct{}
	// ErrProjectMetadataAlreadyExists describes an error, which happens
	// when a metadata key of a project is tried to be created a second time.
	ErrProjectMetadataAlreadyExists struct{}

	ErrProjectMetadataInvalidRequest struct{}
)

// Error returns the error message.
func (e *ErrProjectMetadataValueEnableContentTrustUndefined) Error() string {
	return string(ErrProjectMetadataValueUndefinedMsg + common.ProjectMetadataKeyEnableContentTrust)
}

// Error returns the error message.
func (e *ErrProjectMetadataUndefined) Error() string {
	return ErrProjectMetadataUndefinedMsg
}

func (e *ErrProjectMetadataKeyUndefined) Error() string {
	return ErrProjectMetadataKeyUndefinedMsg
}

// Error returns the error message.
func (e *ErrProjectMetadataAlreadyExists) Error() string {
	return ErrProjectMetadataAlreadyExistsMsg
}

// Error returns the error message.
func (e *ErrProjectMetadataValueAutoScanUndefined) Error() string {
	return string(ErrProjectMetadataValueUndefinedMsg + common.ProjectMetadataKeyAutoScan)
}

// Error returns the error message.
func (e *ErrProjectMetadataValueSeverityUndefined) Error() string {
	return string(ErrProjectMetadataValueUndefinedMsg + common.ProjectMetadataKeySeverity)
}

// Error returns the error message.
func (e *ErrProjectMetadataValueReuseSysCveAllowlistUndefined) Error() string {
	return string(ErrProjectMetadataValueUndefinedMsg + common.ProjectMetadataKeyReuseSysCVEAllowlist)
}

// Error returns the error message.
func (e *ErrProjectMetadataValuePublicUndefined) Error() string {
	return string(ErrProjectMetadataValueUndefinedMsg + common.ProjectMetadataKeyPublic)
}

// Error returns the error message.
func (e *ErrProjectMetadataValuePreventVulUndefined) Error() string {
	return string(ErrProjectMetadataValueUndefinedMsg + common.ProjectMetadataKeyPreventVul)
}

// Error returns the error message.
func (e *ErrProjectMetadataValueRetentionIDUndefined) Error() string {
	return string(ErrProjectMetadataValueUndefinedMsg + common.ProjectMetadataKeyRetentionID)
}

func (e *ErrProjectMetadataInvalidRequest) Error() string {
	return ErrProjectMetadataInvalidRequestMsg
}
