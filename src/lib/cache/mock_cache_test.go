// Code generated by mockery v2.51.0. DO NOT EDIT.

package cache

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// mockCache is an autogenerated mock type for the Cache type
type mockCache struct {
	mock.Mock
}

// Contains provides a mock function with given fields: ctx, key
func (_m *mockCache) Contains(ctx context.Context, key string) bool {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Contains")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, key
func (_m *mockCache) Delete(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, key, value
func (_m *mockCache) Fetch(ctx context.Context, key string, value interface{}) error {
	ret := _m.Called(ctx, key, value)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ping provides a mock function with given fields: ctx
func (_m *mockCache) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Ping")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: ctx, key, value, expiration
func (_m *mockCache) Save(ctx context.Context, key string, value interface{}, expiration ...time.Duration) error {
	_va := make([]interface{}, len(expiration))
	for _i := range expiration {
		_va[_i] = expiration[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key, value)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, ...time.Duration) error); ok {
		r0 = rf(ctx, key, value, expiration...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Scan provides a mock function with given fields: ctx, match
func (_m *mockCache) Scan(ctx context.Context, match string) (Iterator, error) {
	ret := _m.Called(ctx, match)

	if len(ret) == 0 {
		panic("no return value specified for Scan")
	}

	var r0 Iterator
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (Iterator, error)); ok {
		return rf(ctx, match)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) Iterator); ok {
		r0 = rf(ctx, match)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Iterator)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, match)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockCache creates a new instance of mockCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockCache {
	mock := &mockCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
