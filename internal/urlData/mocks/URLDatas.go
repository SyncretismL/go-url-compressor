// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	urlData "compressor/internal/urlData"

	mock "github.com/stretchr/testify/mock"
)

// URLDatas is an autogenerated mock type for the URLDatas type
type URLDatas struct {
	mock.Mock
}

// GetFullURL provides a mock function with given fields: url
func (_m *URLDatas) GetFullURL(url *urlData.URLData) error {
	ret := _m.Called(url)
	//for test only
	url.URL = "http://hhh.com/data/phonenumber2"
	var r0 error
	if rf, ok := ret.Get(0).(func(*urlData.URLData) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetURL provides a mock function with given fields: url
func (_m *URLDatas) SetURL(url *urlData.URLData) error {
	ret := _m.Called(url)

	var r0 error
	if rf, ok := ret.Get(0).(func(*urlData.URLData) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetURLCompressed provides a mock function with given fields: url
func (_m *URLDatas) SetURLCompressed(url *urlData.URLData) error {
	ret := _m.Called(url)

	var r0 error
	if rf, ok := ret.Get(0).(func(*urlData.URLData) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
