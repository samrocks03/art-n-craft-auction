// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	dto "github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
	mock "github.com/stretchr/testify/mock"

	repository "github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"

	uuid "github.com/google/uuid"
)

// UserStorer is an autogenerated mock type for the UserStorer type
type UserStorer struct {
	mock.Mock
}

// CheckEmailExists provides a mock function with given fields: user
func (_m *UserStorer) CheckEmailExists(user repository.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CheckEmailExists")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(repository.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: user
func (_m *UserStorer) CreateUser(user repository.User) (dto.UserSignupResponse, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 dto.UserSignupResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.User) (dto.UserSignupResponse, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(repository.User) dto.UserSignupResponse); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(dto.UserSignupResponse)
	}

	if rf, ok := ret.Get(1).(func(repository.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUsers provides a mock function with given fields: start, count
func (_m *UserStorer) GetAllUsers(start int, count int) ([]dto.GetAllUserResponse, error) {
	ret := _m.Called(start, count)

	if len(ret) == 0 {
		panic("no return value specified for GetAllUsers")
	}

	var r0 []dto.GetAllUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]dto.GetAllUserResponse, error)); ok {
		return rf(start, count)
	}
	if rf, ok := ret.Get(0).(func(int, int) []dto.GetAllUserResponse); ok {
		r0 = rf(start, count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.GetAllUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(start, count)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUsersByRole provides a mock function with given fields: start, count, role
func (_m *UserStorer) GetAllUsersByRole(start int, count int, role string) ([]dto.GetAllUserResponse, error) {
	ret := _m.Called(start, count, role)

	if len(ret) == 0 {
		panic("no return value specified for GetAllUsersByRole")
	}

	var r0 []dto.GetAllUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]dto.GetAllUserResponse, error)); ok {
		return rf(start, count, role)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []dto.GetAllUserResponse); ok {
		r0 = rf(start, count, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.GetAllUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(start, count, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoleID provides a mock function with given fields: role
func (_m *UserStorer) GetRoleID(role string) (uuid.UUID, error) {
	ret := _m.Called(role)

	if len(ret) == 0 {
		panic("no return value specified for GetRoleID")
	}

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (uuid.UUID, error)); ok {
		return rf(role)
	}
	if rf, ok := ret.Get(0).(func(string) uuid.UUID); ok {
		r0 = rf(role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: reqEmail
func (_m *UserStorer) GetUserByEmail(reqEmail string) (dto.User, error) {
	ret := _m.Called(reqEmail)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 dto.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (dto.User, error)); ok {
		return rf(reqEmail)
	}
	if rf, ok := ret.Get(0).(func(string) dto.User); ok {
		r0 = rf(reqEmail)
	} else {
		r0 = ret.Get(0).(dto.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(reqEmail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserStorer creates a new instance of UserStorer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserStorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserStorer {
	mock := &UserStorer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
