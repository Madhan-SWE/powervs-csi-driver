// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ppc64le-cloud/powervs-csi-driver/pkg/cloud (interfaces: Cloud)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	cloud "github.com/ppc64le-cloud/powervs-csi-driver/pkg/cloud"
)

// MockCloud is a mock of Cloud interface.
type MockCloud struct {
	ctrl     *gomock.Controller
	recorder *MockCloudMockRecorder
}

// MockCloudMockRecorder is the mock recorder for MockCloud.
type MockCloudMockRecorder struct {
	mock *MockCloud
}

// NewMockCloud creates a new mock instance.
func NewMockCloud(ctrl *gomock.Controller) *MockCloud {
	mock := &MockCloud{ctrl: ctrl}
	mock.recorder = &MockCloudMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloud) EXPECT() *MockCloudMockRecorder {
	return m.recorder
}

// AttachDisk mocks base method.
func (m *MockCloud) AttachDisk(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachDisk", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AttachDisk indicates an expected call of AttachDisk.
func (mr *MockCloudMockRecorder) AttachDisk(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachDisk", reflect.TypeOf((*MockCloud)(nil).AttachDisk), arg0, arg1)
}

// CreateDisk mocks base method.
func (m *MockCloud) CreateDisk(arg0 string, arg1 *cloud.DiskOptions) (*cloud.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDisk", arg0, arg1)
	ret0, _ := ret[0].(*cloud.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDisk indicates an expected call of CreateDisk.
func (mr *MockCloudMockRecorder) CreateDisk(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDisk", reflect.TypeOf((*MockCloud)(nil).CreateDisk), arg0, arg1)
}

// DeleteDisk mocks base method.
func (m *MockCloud) DeleteDisk(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDisk", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDisk indicates an expected call of DeleteDisk.
func (mr *MockCloudMockRecorder) DeleteDisk(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDisk", reflect.TypeOf((*MockCloud)(nil).DeleteDisk), arg0)
}

// DetachDisk mocks base method.
func (m *MockCloud) DetachDisk(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachDisk", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DetachDisk indicates an expected call of DetachDisk.
func (mr *MockCloudMockRecorder) DetachDisk(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachDisk", reflect.TypeOf((*MockCloud)(nil).DetachDisk), arg0, arg1)
}

// GetDiskByID mocks base method.
func (m *MockCloud) GetDiskByID(arg0 string) (*cloud.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskByID", arg0)
	ret0, _ := ret[0].(*cloud.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiskByID indicates an expected call of GetDiskByID.
func (mr *MockCloudMockRecorder) GetDiskByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskByID", reflect.TypeOf((*MockCloud)(nil).GetDiskByID), arg0)
}

// GetDiskByName mocks base method.
func (m *MockCloud) GetDiskByName(arg0 string) (*cloud.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskByName", arg0)
	ret0, _ := ret[0].(*cloud.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiskByName indicates an expected call of GetDiskByName.
func (mr *MockCloudMockRecorder) GetDiskByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskByName", reflect.TypeOf((*MockCloud)(nil).GetDiskByName), arg0)
}

// GetImageByID mocks base method.
func (m *MockCloud) GetImageByID(arg0 string) (*cloud.PVMImage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageByID", arg0)
	ret0, _ := ret[0].(*cloud.PVMImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageByID indicates an expected call of GetImageByID.
func (mr *MockCloudMockRecorder) GetImageByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageByID", reflect.TypeOf((*MockCloud)(nil).GetImageByID), arg0)
}

// GetPVMInstanceByID mocks base method.
func (m *MockCloud) GetPVMInstanceByID(arg0 string) (*cloud.PVMInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPVMInstanceByID", arg0)
	ret0, _ := ret[0].(*cloud.PVMInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPVMInstanceByID indicates an expected call of GetPVMInstanceByID.
func (mr *MockCloudMockRecorder) GetPVMInstanceByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPVMInstanceByID", reflect.TypeOf((*MockCloud)(nil).GetPVMInstanceByID), arg0)
}

// GetPVMInstanceByName mocks base method.
func (m *MockCloud) GetPVMInstanceByName(arg0 string) (*cloud.PVMInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPVMInstanceByName", arg0)
	ret0, _ := ret[0].(*cloud.PVMInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPVMInstanceByName indicates an expected call of GetPVMInstanceByName.
func (mr *MockCloudMockRecorder) GetPVMInstanceByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPVMInstanceByName", reflect.TypeOf((*MockCloud)(nil).GetPVMInstanceByName), arg0)
}

// IsAttached mocks base method.
func (m *MockCloud) IsAttached(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAttached", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAttached indicates an expected call of IsAttached.
func (mr *MockCloudMockRecorder) IsAttached(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAttached", reflect.TypeOf((*MockCloud)(nil).IsAttached), arg0, arg1)
}

// ResizeDisk mocks base method.
func (m *MockCloud) ResizeDisk(arg0 string, arg1 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResizeDisk", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResizeDisk indicates an expected call of ResizeDisk.
func (mr *MockCloudMockRecorder) ResizeDisk(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResizeDisk", reflect.TypeOf((*MockCloud)(nil).ResizeDisk), arg0, arg1)
}

// WaitForAttachmentState mocks base method.
func (m *MockCloud) WaitForAttachmentState(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForAttachmentState", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForAttachmentState indicates an expected call of WaitForAttachmentState.
func (mr *MockCloudMockRecorder) WaitForAttachmentState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForAttachmentState", reflect.TypeOf((*MockCloud)(nil).WaitForAttachmentState), arg0, arg1)
}
