// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	proto "github.com/misyuari/go-omdb/domain/movie/delivery/grpc/proto"
)

// MoviesClient is an autogenerated mock type for the MoviesClient type
type MoviesClient struct {
	mock.Mock
}

// Detail provides a mock function with given fields: ctx, in, opts
func (_m *MoviesClient) Detail(ctx context.Context, in *proto.MovieRequestDetail, opts ...grpc.CallOption) (*proto.MovieDetail, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.MovieDetail
	if rf, ok := ret.Get(0).(func(context.Context, *proto.MovieRequestDetail, ...grpc.CallOption) *proto.MovieDetail); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.MovieDetail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.MovieRequestDetail, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, in, opts
func (_m *MoviesClient) List(ctx context.Context, in *proto.MovieRequestList, opts ...grpc.CallOption) (*proto.MovieList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.MovieList
	if rf, ok := ret.Get(0).(func(context.Context, *proto.MovieRequestList, ...grpc.CallOption) *proto.MovieList); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.MovieList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.MovieRequestList, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}