// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	model "github.com/misyuari/go-omdb/domain/log/model"
	mock "github.com/stretchr/testify/mock"
)

// LogRepository is an autogenerated mock type for the LogRepository type
type LogRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: log
func (_m *LogRepository) Create(log *model.Log) {
	_m.Called(log)
}
