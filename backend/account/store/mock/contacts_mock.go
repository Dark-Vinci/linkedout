// Code generated by MockGen. DO NOT EDIT.
// Source: contacts.go
//
// Generated by this command:
//
//	mockgen -source contacts.go -destination ./mock/contacts_mock.go -package mock ContactDatabase
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	account "github.com/dark-vinci/wapp/backend/sdk/models/account"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockContactDatabase is a mock of ContactDatabase interface.
type MockContactDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockContactDatabaseMockRecorder
}

// MockContactDatabaseMockRecorder is the mock recorder for MockContactDatabase.
type MockContactDatabaseMockRecorder struct {
	mock *MockContactDatabase
}

// NewMockContactDatabase creates a new mock instance.
func NewMockContactDatabase(ctrl *gomock.Controller) *MockContactDatabase {
	mock := &MockContactDatabase{ctrl: ctrl}
	mock.recorder = &MockContactDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContactDatabase) EXPECT() *MockContactDatabaseMockRecorder {
	return m.recorder
}

// BlockContact mocks base method.
func (m *MockContactDatabase) BlockContact(ctx context.Context, contactID, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockContact", ctx, contactID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// BlockContact indicates an expected call of BlockContact.
func (mr *MockContactDatabaseMockRecorder) BlockContact(ctx, contactID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockContact", reflect.TypeOf((*MockContactDatabase)(nil).BlockContact), ctx, contactID, userID)
}

// Create mocks base method.
func (m *MockContactDatabase) Create(ctx context.Context, contact account.Contacts) (*account.Contacts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, contact)
	ret0, _ := ret[0].(*account.Contacts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockContactDatabaseMockRecorder) Create(ctx, contact any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockContactDatabase)(nil).Create), ctx, contact)
}

// Delete mocks base method.
func (m *MockContactDatabase) Delete(ctx context.Context, deletedAt time.Time, contactID, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, deletedAt, contactID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockContactDatabaseMockRecorder) Delete(ctx, deletedAt, contactID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockContactDatabase)(nil).Delete), ctx, deletedAt, contactID, userID)
}

// DeleteAllUserContacts mocks base method.
func (m *MockContactDatabase) DeleteAllUserContacts(ctx context.Context, userID uuid.UUID, deletedAt time.Time, tx *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllUserContacts", ctx, userID, deletedAt, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllUserContacts indicates an expected call of DeleteAllUserContacts.
func (mr *MockContactDatabaseMockRecorder) DeleteAllUserContacts(ctx, userID, deletedAt, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllUserContacts", reflect.TypeOf((*MockContactDatabase)(nil).DeleteAllUserContacts), ctx, userID, deletedAt, tx)
}

// GetBlocked mocks base method.
func (m *MockContactDatabase) GetBlocked(ctx context.Context, userID uuid.UUID) ([]account.Contacts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlocked", ctx, userID)
	ret0, _ := ret[0].([]account.Contacts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlocked indicates an expected call of GetBlocked.
func (mr *MockContactDatabaseMockRecorder) GetBlocked(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlocked", reflect.TypeOf((*MockContactDatabase)(nil).GetBlocked), ctx, userID)
}

// GetUserContacts mocks base method.
func (m *MockContactDatabase) GetUserContacts(ctx context.Context, userID uuid.UUID) ([]account.Contacts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserContacts", ctx, userID)
	ret0, _ := ret[0].([]account.Contacts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserContacts indicates an expected call of GetUserContacts.
func (mr *MockContactDatabaseMockRecorder) GetUserContacts(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserContacts", reflect.TypeOf((*MockContactDatabase)(nil).GetUserContacts), ctx, userID)
}

// MakeFavourite mocks base method.
func (m *MockContactDatabase) MakeFavourite(ctx context.Context, contactID, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeFavourite", ctx, contactID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeFavourite indicates an expected call of MakeFavourite.
func (mr *MockContactDatabaseMockRecorder) MakeFavourite(ctx, contactID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeFavourite", reflect.TypeOf((*MockContactDatabase)(nil).MakeFavourite), ctx, contactID, userID)
}

// MakeUnFavourite mocks base method.
func (m *MockContactDatabase) MakeUnFavourite(ctx context.Context, contactID, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeUnFavourite", ctx, contactID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeUnFavourite indicates an expected call of MakeUnFavourite.
func (mr *MockContactDatabaseMockRecorder) MakeUnFavourite(ctx, contactID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeUnFavourite", reflect.TypeOf((*MockContactDatabase)(nil).MakeUnFavourite), ctx, contactID, userID)
}

// UnblockContact mocks base method.
func (m *MockContactDatabase) UnblockContact(ctx context.Context, contactID, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnblockContact", ctx, contactID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnblockContact indicates an expected call of UnblockContact.
func (mr *MockContactDatabaseMockRecorder) UnblockContact(ctx, contactID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnblockContact", reflect.TypeOf((*MockContactDatabase)(nil).UnblockContact), ctx, contactID, userID)
}