// +build !integration
// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	products "github.com/mittwald/goharbor-client/internal/api/v1.10.0/client/products"
	mock "github.com/stretchr/testify/mock"
)

// MockClientService is an autogenerated mock type for the ClientService type
type MockClientService struct {
	mock.Mock
}

// DeleteProjectsProjectID provides a mock function with given fields: params, authInfo
func (_m *MockClientService) DeleteProjectsProjectID(params *products.DeleteProjectsProjectIDParams, authInfo runtime.ClientAuthInfoWriter) (*products.DeleteProjectsProjectIDOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.DeleteProjectsProjectIDOK
	if rf, ok := ret.Get(0).(func(*products.DeleteProjectsProjectIDParams, runtime.ClientAuthInfoWriter) *products.DeleteProjectsProjectIDOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.DeleteProjectsProjectIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.DeleteProjectsProjectIDParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProjectsProjectIDMembersMid provides a mock function with given fields: params, authInfo
func (_m *MockClientService) DeleteProjectsProjectIDMembersMid(params *products.DeleteProjectsProjectIDMembersMidParams, authInfo runtime.ClientAuthInfoWriter) (*products.DeleteProjectsProjectIDMembersMidOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.DeleteProjectsProjectIDMembersMidOK
	if rf, ok := ret.Get(0).(func(*products.DeleteProjectsProjectIDMembersMidParams, runtime.ClientAuthInfoWriter) *products.DeleteProjectsProjectIDMembersMidOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.DeleteProjectsProjectIDMembersMidOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.DeleteProjectsProjectIDMembersMidParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProjectsProjectIDMetadatasMetaName provides a mock function with given fields: params, authInfo
func (_m *MockClientService) DeleteProjectsProjectIDMetadatasMetaName(params *products.DeleteProjectsProjectIDMetadatasMetaNameParams, authInfo runtime.ClientAuthInfoWriter) (*products.DeleteProjectsProjectIDMetadatasMetaNameOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.DeleteProjectsProjectIDMetadatasMetaNameOK
	if rf, ok := ret.Get(0).(func(*products.DeleteProjectsProjectIDMetadatasMetaNameParams, runtime.ClientAuthInfoWriter) *products.DeleteProjectsProjectIDMetadatasMetaNameOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.DeleteProjectsProjectIDMetadatasMetaNameOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.DeleteProjectsProjectIDMetadatasMetaNameParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRegistriesID provides a mock function with given fields: params, authInfo
func (_m *MockClientService) DeleteRegistriesID(params *products.DeleteRegistriesIDParams, authInfo runtime.ClientAuthInfoWriter) (*products.DeleteRegistriesIDOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.DeleteRegistriesIDOK
	if rf, ok := ret.Get(0).(func(*products.DeleteRegistriesIDParams, runtime.ClientAuthInfoWriter) *products.DeleteRegistriesIDOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.DeleteRegistriesIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.DeleteRegistriesIDParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteReplicationPoliciesID provides a mock function with given fields: params, authInfo
func (_m *MockClientService) DeleteReplicationPoliciesID(params *products.DeleteReplicationPoliciesIDParams, authInfo runtime.ClientAuthInfoWriter) (*products.DeleteReplicationPoliciesIDOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.DeleteReplicationPoliciesIDOK
	if rf, ok := ret.Get(0).(func(*products.DeleteReplicationPoliciesIDParams, runtime.ClientAuthInfoWriter) *products.DeleteReplicationPoliciesIDOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.DeleteReplicationPoliciesIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.DeleteReplicationPoliciesIDParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUsersUserID provides a mock function with given fields: params, authInfo
func (_m *MockClientService) DeleteUsersUserID(params *products.DeleteUsersUserIDParams, authInfo runtime.ClientAuthInfoWriter) (*products.DeleteUsersUserIDOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.DeleteUsersUserIDOK
	if rf, ok := ret.Get(0).(func(*products.DeleteUsersUserIDParams, runtime.ClientAuthInfoWriter) *products.DeleteUsersUserIDOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.DeleteUsersUserIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.DeleteUsersUserIDParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHealth provides a mock function with given fields: params, authInfo
func (_m *MockClientService) GetHealth(params *products.GetHealthParams, authInfo runtime.ClientAuthInfoWriter) (*products.GetHealthOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.GetHealthOK
	if rf, ok := ret.Get(0).(func(*products.GetHealthParams, runtime.ClientAuthInfoWriter) *products.GetHealthOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.GetHealthOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.GetHealthParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjects provides a mock function with given fields: params, authInfo
func (_m *MockClientService) GetProjects(params *products.GetProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*products.GetProjectsOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.GetProjectsOK
	if rf, ok := ret.Get(0).(func(*products.GetProjectsParams, runtime.ClientAuthInfoWriter) *products.GetProjectsOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.GetProjectsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.GetProjectsParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectsProjectID provides a mock function with given fields: params, authInfo
func (_m *MockClientService) GetProjectsProjectID(params *products.GetProjectsProjectIDParams, authInfo runtime.ClientAuthInfoWriter) (*products.GetProjectsProjectIDOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.GetProjectsProjectIDOK
	if rf, ok := ret.Get(0).(func(*products.GetProjectsProjectIDParams, runtime.ClientAuthInfoWriter) *products.GetProjectsProjectIDOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.GetProjectsProjectIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.GetProjectsProjectIDParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectsProjectIDMembers provides a mock function with given fields: params, authInfo
func (_m *MockClientService) GetProjectsProjectIDMembers(params *products.GetProjectsProjectIDMembersParams, authInfo runtime.ClientAuthInfoWriter) (*products.GetProjectsProjectIDMembersOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.GetProjectsProjectIDMembersOK
	if rf, ok := ret.Get(0).(func(*products.GetProjectsProjectIDMembersParams, runtime.ClientAuthInfoWriter) *products.GetProjectsProjectIDMembersOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.GetProjectsProjectIDMembersOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.GetProjectsProjectIDMembersParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectsProjectIDMetadatas provides a mock function with given fields: params, authInfo
func (_m *MockClientService) GetProjectsProjectIDMetadatas(params *products.GetProjectsProjectIDMetadatasParams, authInfo runtime.ClientAuthInfoWriter) (*products.GetProjectsProjectIDMetadatasOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.GetProjectsProjectIDMetadatasOK
	if rf, ok := ret.Get(0).(func(*products.GetProjectsProjectIDMetadatasParams, runtime.ClientAuthInfoWriter) *products.GetProjectsProjectIDMetadatasOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.GetProjectsProjectIDMetadatasOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.GetProjectsProjectIDMetadatasParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectsProjectIDMetadatasMetaName provides a mock function with given fields: params, authInfo
func (_m *MockClientService) GetProjectsProjectIDMetadatasMetaName(params *products.GetProjectsProjectIDMetadatasMetaNameParams, authInfo runtime.ClientAuthInfoWriter) (*products.GetProjectsProjectIDMetadatasMetaNameOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.GetProjectsProjectIDMetadatasMetaNameOK
	if rf, ok := ret.Get(0).(func(*products.GetProjectsProjectIDMetadatasMetaNameParams, runtime.ClientAuthInfoWriter) *products.GetProjectsProjectIDMetadatasMetaNameOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.GetProjectsProjectIDMetadatasMetaNameOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.GetProjectsProjectIDMetadatasMetaNameParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRegistries provides a mock function with given fields: params, authInfo
func (_m *MockClientService) GetRegistries(params *products.GetRegistriesParams, authInfo runtime.ClientAuthInfoWriter) (*products.GetRegistriesOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.GetRegistriesOK
	if rf, ok := ret.Get(0).(func(*products.GetRegistriesParams, runtime.ClientAuthInfoWriter) *products.GetRegistriesOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.GetRegistriesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.GetRegistriesParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReplicationPolicies provides a mock function with given fields: params, authInfo
func (_m *MockClientService) GetReplicationPolicies(params *products.GetReplicationPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*products.GetReplicationPoliciesOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.GetReplicationPoliciesOK
	if rf, ok := ret.Get(0).(func(*products.GetReplicationPoliciesParams, runtime.ClientAuthInfoWriter) *products.GetReplicationPoliciesOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.GetReplicationPoliciesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.GetReplicationPoliciesParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReplicationPoliciesID provides a mock function with given fields: params, authInfo
func (_m *MockClientService) GetReplicationPoliciesID(params *products.GetReplicationPoliciesIDParams, authInfo runtime.ClientAuthInfoWriter) (*products.GetReplicationPoliciesIDOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.GetReplicationPoliciesIDOK
	if rf, ok := ret.Get(0).(func(*products.GetReplicationPoliciesIDParams, runtime.ClientAuthInfoWriter) *products.GetReplicationPoliciesIDOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.GetReplicationPoliciesIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.GetReplicationPoliciesIDParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSystemGcSchedule provides a mock function with given fields: params, authInfo
func (_m *MockClientService) GetSystemGcSchedule(params *products.GetSystemGcScheduleParams, authInfo runtime.ClientAuthInfoWriter) (*products.GetSystemGcScheduleOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.GetSystemGcScheduleOK
	if rf, ok := ret.Get(0).(func(*products.GetSystemGcScheduleParams, runtime.ClientAuthInfoWriter) *products.GetSystemGcScheduleOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.GetSystemGcScheduleOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.GetSystemGcScheduleParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: params, authInfo
func (_m *MockClientService) GetUsers(params *products.GetUsersParams, authInfo runtime.ClientAuthInfoWriter) (*products.GetUsersOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.GetUsersOK
	if rf, ok := ret.Get(0).(func(*products.GetUsersParams, runtime.ClientAuthInfoWriter) *products.GetUsersOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.GetUsersOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.GetUsersParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostProjects provides a mock function with given fields: params, authInfo
func (_m *MockClientService) PostProjects(params *products.PostProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*products.PostProjectsCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.PostProjectsCreated
	if rf, ok := ret.Get(0).(func(*products.PostProjectsParams, runtime.ClientAuthInfoWriter) *products.PostProjectsCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.PostProjectsCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.PostProjectsParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostProjectsProjectIDMembers provides a mock function with given fields: params, authInfo
func (_m *MockClientService) PostProjectsProjectIDMembers(params *products.PostProjectsProjectIDMembersParams, authInfo runtime.ClientAuthInfoWriter) (*products.PostProjectsProjectIDMembersCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.PostProjectsProjectIDMembersCreated
	if rf, ok := ret.Get(0).(func(*products.PostProjectsProjectIDMembersParams, runtime.ClientAuthInfoWriter) *products.PostProjectsProjectIDMembersCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.PostProjectsProjectIDMembersCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.PostProjectsProjectIDMembersParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostProjectsProjectIDMetadatas provides a mock function with given fields: params, authInfo
func (_m *MockClientService) PostProjectsProjectIDMetadatas(params *products.PostProjectsProjectIDMetadatasParams, authInfo runtime.ClientAuthInfoWriter) (*products.PostProjectsProjectIDMetadatasOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.PostProjectsProjectIDMetadatasOK
	if rf, ok := ret.Get(0).(func(*products.PostProjectsProjectIDMetadatasParams, runtime.ClientAuthInfoWriter) *products.PostProjectsProjectIDMetadatasOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.PostProjectsProjectIDMetadatasOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.PostProjectsProjectIDMetadatasParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostRegistries provides a mock function with given fields: params, authInfo
func (_m *MockClientService) PostRegistries(params *products.PostRegistriesParams, authInfo runtime.ClientAuthInfoWriter) (*products.PostRegistriesCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.PostRegistriesCreated
	if rf, ok := ret.Get(0).(func(*products.PostRegistriesParams, runtime.ClientAuthInfoWriter) *products.PostRegistriesCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.PostRegistriesCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.PostRegistriesParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostReplicationPolicies provides a mock function with given fields: params, authInfo
func (_m *MockClientService) PostReplicationPolicies(params *products.PostReplicationPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*products.PostReplicationPoliciesCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.PostReplicationPoliciesCreated
	if rf, ok := ret.Get(0).(func(*products.PostReplicationPoliciesParams, runtime.ClientAuthInfoWriter) *products.PostReplicationPoliciesCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.PostReplicationPoliciesCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.PostReplicationPoliciesParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostSystemGcSchedule provides a mock function with given fields: params, authInfo
func (_m *MockClientService) PostSystemGcSchedule(params *products.PostSystemGcScheduleParams, authInfo runtime.ClientAuthInfoWriter) (*products.PostSystemGcScheduleOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.PostSystemGcScheduleOK
	if rf, ok := ret.Get(0).(func(*products.PostSystemGcScheduleParams, runtime.ClientAuthInfoWriter) *products.PostSystemGcScheduleOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.PostSystemGcScheduleOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.PostSystemGcScheduleParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostUsers provides a mock function with given fields: params, authInfo
func (_m *MockClientService) PostUsers(params *products.PostUsersParams, authInfo runtime.ClientAuthInfoWriter) (*products.PostUsersCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.PostUsersCreated
	if rf, ok := ret.Get(0).(func(*products.PostUsersParams, runtime.ClientAuthInfoWriter) *products.PostUsersCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.PostUsersCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.PostUsersParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutProjectsProjectID provides a mock function with given fields: params, authInfo
func (_m *MockClientService) PutProjectsProjectID(params *products.PutProjectsProjectIDParams, authInfo runtime.ClientAuthInfoWriter) (*products.PutProjectsProjectIDOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.PutProjectsProjectIDOK
	if rf, ok := ret.Get(0).(func(*products.PutProjectsProjectIDParams, runtime.ClientAuthInfoWriter) *products.PutProjectsProjectIDOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.PutProjectsProjectIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.PutProjectsProjectIDParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutProjectsProjectIDMembersMid provides a mock function with given fields: params, authInfo
func (_m *MockClientService) PutProjectsProjectIDMembersMid(params *products.PutProjectsProjectIDMembersMidParams, authInfo runtime.ClientAuthInfoWriter) (*products.PutProjectsProjectIDMembersMidOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.PutProjectsProjectIDMembersMidOK
	if rf, ok := ret.Get(0).(func(*products.PutProjectsProjectIDMembersMidParams, runtime.ClientAuthInfoWriter) *products.PutProjectsProjectIDMembersMidOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.PutProjectsProjectIDMembersMidOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.PutProjectsProjectIDMembersMidParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutProjectsProjectIDMetadatasMetaName provides a mock function with given fields: params, authInfo
func (_m *MockClientService) PutProjectsProjectIDMetadatasMetaName(params *products.PutProjectsProjectIDMetadatasMetaNameParams, authInfo runtime.ClientAuthInfoWriter) (*products.PutProjectsProjectIDMetadatasMetaNameOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.PutProjectsProjectIDMetadatasMetaNameOK
	if rf, ok := ret.Get(0).(func(*products.PutProjectsProjectIDMetadatasMetaNameParams, runtime.ClientAuthInfoWriter) *products.PutProjectsProjectIDMetadatasMetaNameOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.PutProjectsProjectIDMetadatasMetaNameOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.PutProjectsProjectIDMetadatasMetaNameParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutRegistriesID provides a mock function with given fields: params, authInfo
func (_m *MockClientService) PutRegistriesID(params *products.PutRegistriesIDParams, authInfo runtime.ClientAuthInfoWriter) (*products.PutRegistriesIDOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.PutRegistriesIDOK
	if rf, ok := ret.Get(0).(func(*products.PutRegistriesIDParams, runtime.ClientAuthInfoWriter) *products.PutRegistriesIDOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.PutRegistriesIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.PutRegistriesIDParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutReplicationPoliciesID provides a mock function with given fields: params, authInfo
func (_m *MockClientService) PutReplicationPoliciesID(params *products.PutReplicationPoliciesIDParams, authInfo runtime.ClientAuthInfoWriter) (*products.PutReplicationPoliciesIDOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.PutReplicationPoliciesIDOK
	if rf, ok := ret.Get(0).(func(*products.PutReplicationPoliciesIDParams, runtime.ClientAuthInfoWriter) *products.PutReplicationPoliciesIDOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.PutReplicationPoliciesIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.PutReplicationPoliciesIDParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutSystemGcSchedule provides a mock function with given fields: params, authInfo
func (_m *MockClientService) PutSystemGcSchedule(params *products.PutSystemGcScheduleParams, authInfo runtime.ClientAuthInfoWriter) (*products.PutSystemGcScheduleOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.PutSystemGcScheduleOK
	if rf, ok := ret.Get(0).(func(*products.PutSystemGcScheduleParams, runtime.ClientAuthInfoWriter) *products.PutSystemGcScheduleOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.PutSystemGcScheduleOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.PutSystemGcScheduleParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutUsersUserID provides a mock function with given fields: params, authInfo
func (_m *MockClientService) PutUsersUserID(params *products.PutUsersUserIDParams, authInfo runtime.ClientAuthInfoWriter) (*products.PutUsersUserIDOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *products.PutUsersUserIDOK
	if rf, ok := ret.Get(0).(func(*products.PutUsersUserIDParams, runtime.ClientAuthInfoWriter) *products.PutUsersUserIDOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.PutUsersUserIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.PutUsersUserIDParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}
