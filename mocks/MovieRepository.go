// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	model "github.com/misyuari/go-omdb/domain/movie/model"
	mock "github.com/stretchr/testify/mock"
)

// MovieRepository is an autogenerated mock type for the MovieRepository type
type MovieRepository struct {
	mock.Mock
}

// Detail provides a mock function with given fields: id
func (_m *MovieRepository) Detail(id string) (*model.Movie, error) {
	ret := _m.Called(id)

	var r0 *model.Movie
	if rf, ok := ret.Get(0).(func(string) *model.Movie); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: page, keyword
func (_m *MovieRepository) List(page uint32, keyword string) (*[]model.Movie, error) {
	ret := _m.Called(page, keyword)

	var r0 *[]model.Movie
	if rf, ok := ret.Get(0).(func(uint32, string) *[]model.Movie); ok {
		r0 = rf(page, keyword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]model.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint32, string) error); ok {
		r1 = rf(page, keyword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}