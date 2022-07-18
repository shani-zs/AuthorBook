// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package store is a generated GoMock package.
package store

import (
	context "context"
	entities "projects/GoLang-Interns-2022/authorbook/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthorStorer is a mock of AuthorStorer interface.
type MockAuthorStorer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorStorerMockRecorder
}

// MockAuthorStorerMockRecorder is the mock recorder for MockAuthorStorer.
type MockAuthorStorerMockRecorder struct {
	mock *MockAuthorStorer
}

// NewMockAuthorStorer creates a new mock instance.
func NewMockAuthorStorer(ctrl *gomock.Controller) *MockAuthorStorer {
	mock := &MockAuthorStorer{ctrl: ctrl}
	mock.recorder = &MockAuthorStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorStorer) EXPECT() *MockAuthorStorerMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockAuthorStorer) Delete(ctx context.Context, id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthorStorerMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthorStorer)(nil).Delete), ctx, id)
}

// IncludeAuthor mocks base method.
func (m *MockAuthorStorer) IncludeAuthor(ctx context.Context, id int) (entities.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncludeAuthor", ctx, id)
	ret0, _ := ret[0].(entities.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncludeAuthor indicates an expected call of IncludeAuthor.
func (mr *MockAuthorStorerMockRecorder) IncludeAuthor(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncludeAuthor", reflect.TypeOf((*MockAuthorStorer)(nil).IncludeAuthor), ctx, id)
}

// Post mocks base method.
func (m *MockAuthorStorer) Post(ctx context.Context, author entities.Author) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", ctx, author)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockAuthorStorerMockRecorder) Post(ctx, author interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockAuthorStorer)(nil).Post), ctx, author)
}

// Put mocks base method.
func (m *MockAuthorStorer) Put(ctx context.Context, author entities.Author, id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", ctx, author, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockAuthorStorerMockRecorder) Put(ctx, author, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockAuthorStorer)(nil).Put), ctx, author, id)
}

// MockBookStorer is a mock of BookStorer interface.
type MockBookStorer struct {
	ctrl     *gomock.Controller
	recorder *MockBookStorerMockRecorder
}

// MockBookStorerMockRecorder is the mock recorder for MockBookStorer.
type MockBookStorerMockRecorder struct {
	mock *MockBookStorer
}

// NewMockBookStorer creates a new mock instance.
func NewMockBookStorer(ctrl *gomock.Controller) *MockBookStorer {
	mock := &MockBookStorer{ctrl: ctrl}
	mock.recorder = &MockBookStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookStorer) EXPECT() *MockBookStorerMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockBookStorer) Delete(ctx context.Context, id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockBookStorerMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBookStorer)(nil).Delete), ctx, id)
}

// GetAllBook mocks base method.
func (m *MockBookStorer) GetAllBook(ctx context.Context) ([]entities.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBook", ctx)
	ret0, _ := ret[0].([]entities.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBook indicates an expected call of GetAllBook.
func (mr *MockBookStorerMockRecorder) GetAllBook(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBook", reflect.TypeOf((*MockBookStorer)(nil).GetAllBook), ctx)
}

// GetBookByID mocks base method.
func (m *MockBookStorer) GetBookByID(ctx context.Context, id int) (entities.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookByID", ctx, id)
	ret0, _ := ret[0].(entities.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByID indicates an expected call of GetBookByID.
func (mr *MockBookStorerMockRecorder) GetBookByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByID", reflect.TypeOf((*MockBookStorer)(nil).GetBookByID), ctx, id)
}

// GetBooksByTitle mocks base method.
func (m *MockBookStorer) GetBooksByTitle(ctx context.Context, title string) ([]entities.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooksByTitle", ctx, title)
	ret0, _ := ret[0].([]entities.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooksByTitle indicates an expected call of GetBooksByTitle.
func (mr *MockBookStorerMockRecorder) GetBooksByTitle(ctx, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooksByTitle", reflect.TypeOf((*MockBookStorer)(nil).GetBooksByTitle), ctx, title)
}

// Post mocks base method.
func (m *MockBookStorer) Post(ctx context.Context, book *entities.Book) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", ctx, book)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockBookStorerMockRecorder) Post(ctx, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockBookStorer)(nil).Post), ctx, book)
}

// Put mocks base method.
func (m *MockBookStorer) Put(ctx context.Context, book *entities.Book, id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", ctx, book, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockBookStorerMockRecorder) Put(ctx, book, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockBookStorer)(nil).Put), ctx, book, id)
}