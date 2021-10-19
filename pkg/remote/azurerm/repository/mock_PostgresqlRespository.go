// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package repository

import (
	armpostgresql "github.com/Azure/azure-sdk-for-go/sdk/postgresql/armpostgresql"
	mock "github.com/stretchr/testify/mock"
)

// MockPostgresqlRespository is an autogenerated mock type for the PostgresqlRespository type
type MockPostgresqlRespository struct {
	mock.Mock
}

// ListAllDatabasesByServer provides a mock function with given fields: server
func (_m *MockPostgresqlRespository) ListAllDatabasesByServer(server *armpostgresql.Server) ([]*armpostgresql.Database, error) {
	ret := _m.Called(server)

	var r0 []*armpostgresql.Database
	if rf, ok := ret.Get(0).(func(*armpostgresql.Server) []*armpostgresql.Database); ok {
		r0 = rf(server)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*armpostgresql.Database)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*armpostgresql.Server) error); ok {
		r1 = rf(server)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllServers provides a mock function with given fields:
func (_m *MockPostgresqlRespository) ListAllServers() ([]*armpostgresql.Server, error) {
	ret := _m.Called()

	var r0 []*armpostgresql.Server
	if rf, ok := ret.Get(0).(func() []*armpostgresql.Server); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*armpostgresql.Server)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}