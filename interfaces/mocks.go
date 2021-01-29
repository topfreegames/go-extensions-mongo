package interfaces

import (
	context "context"
	mgo "github.com/globalsign/mgo"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMongoDB is a mock of MongoDB interface
type MockMongoDB struct {
	ctrl     *gomock.Controller
	recorder *MockMongoDBMockRecorder
}

// MockMongoDBMockRecorder is the mock recorder for MockMongoDB
type MockMongoDBMockRecorder struct {
	mock *MockMongoDB
}

// NewMockMongoDB creates a new mock instance
func NewMockMongoDB(ctrl *gomock.Controller) *MockMongoDB {
	mock := &MockMongoDB{ctrl: ctrl}
	mock.recorder = &MockMongoDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMongoDB) EXPECT() *MockMongoDBMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockMongoDB) Run(cmd, result interface{}) error {
	ret := m.ctrl.Call(m, "Run", cmd, result)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockMongoDBMockRecorder) Run(cmd, result interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockMongoDB)(nil).Run), cmd, result)
}

// C mocks base method
func (m *MockMongoDB) C(name string) (Collection, Session) {
	ret := m.ctrl.Call(m, "C", name)
	ret0, _ := ret[0].(Collection)
	ret1, _ := ret[1].(Session)
	return ret0, ret1
}

// C indicates an expected call of C
func (mr *MockMongoDBMockRecorder) C(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "C", reflect.TypeOf((*MockMongoDB)(nil).C), name)
}

// Close mocks base method
func (m *MockMongoDB) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockMongoDBMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMongoDB)(nil).Close))
}

// WithContext mocks base method
func (m *MockMongoDB) WithContext(ctx context.Context) MongoDB {
	ret := m.ctrl.Call(m, "WithContext", ctx)
	ret0, _ := ret[0].(MongoDB)
	return ret0
}

// WithContext indicates an expected call of WithContext
func (mr *MockMongoDBMockRecorder) WithContext(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockMongoDB)(nil).WithContext), ctx)
}

// MockCollection is a mock of Collection interface
type MockCollection struct {
	ctrl     *gomock.Controller
	recorder *MockCollectionMockRecorder
}

// MockCollectionMockRecorder is the mock recorder for MockCollection
type MockCollectionMockRecorder struct {
	mock *MockCollection
}

// NewMockCollection creates a new mock instance
func NewMockCollection(ctrl *gomock.Controller) *MockCollection {
	mock := &MockCollection{ctrl: ctrl}
	mock.recorder = &MockCollectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCollection) EXPECT() *MockCollectionMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockCollection) Find(query interface{}) Query {
	ret := m.ctrl.Call(m, "Find", query)
	ret0, _ := ret[0].(Query)
	return ret0
}

// Find indicates an expected call of Find
func (mr *MockCollectionMockRecorder) Find(query interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCollection)(nil).Find), query)
}

// FindId mocks base method
func (m *MockCollection) FindId(id interface{}) Query {
	ret := m.ctrl.Call(m, "FindId", id)
	ret0, _ := ret[0].(Query)
	return ret0
}

// FindId indicates an expected call of FindId
func (mr *MockCollectionMockRecorder) FindId(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindId", reflect.TypeOf((*MockCollection)(nil).FindId), id)
}

// Pipe mocks base method
func (m *MockCollection) Pipe(pipeline interface{}) Pipe {
	ret := m.ctrl.Call(m, "Pipe", pipeline)
	ret0, _ := ret[0].(Pipe)
	return ret0
}

// Pipe indicates an expected call of Pipe
func (mr *MockCollectionMockRecorder) Pipe(pipeline interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipe", reflect.TypeOf((*MockCollection)(nil).Pipe), pipeline)
}

// Insert mocks base method
func (m *MockCollection) Insert(docs ...interface{}) error {
	varargs := []interface{}{}
	for _, a := range docs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Insert", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockCollectionMockRecorder) Insert(docs ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockCollection)(nil).Insert), docs...)
}

// UpsertId mocks base method
func (m *MockCollection) UpsertId(id, update interface{}) (*mgo.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "UpsertId", id, update)
	ret0, _ := ret[0].(*mgo.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertId indicates an expected call of UpsertId
func (mr *MockCollectionMockRecorder) UpsertId(id, update interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertId", reflect.TypeOf((*MockCollection)(nil).UpsertId), id, update)
}

// Upsert mocks base method
func (m *MockCollection) Upsert(selector, update interface{}) (*mgo.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "Upsert", selector, update)
	ret0, _ := ret[0].(*mgo.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert
func (mr *MockCollectionMockRecorder) Upsert(selector, update interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockCollection)(nil).Upsert), selector, update)
}

// RemoveId mocks base method
func (m *MockCollection) RemoveId(id interface{}) error {
	ret := m.ctrl.Call(m, "RemoveId", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveId indicates an expected call of RemoveId
func (mr *MockCollectionMockRecorder) RemoveId(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveId", reflect.TypeOf((*MockCollection)(nil).RemoveId), id)
}

// Remove mocks base method
func (m *MockCollection) Remove(selector interface{}) error {
	ret := m.ctrl.Call(m, "Remove", selector)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockCollectionMockRecorder) Remove(selector interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockCollection)(nil).Remove), selector)
}

// RemoveAll mocks base method
func (m *MockCollection) RemoveAll(selector interface{}) (*mgo.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "RemoveAll", selector)
	ret0, _ := ret[0].(*mgo.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveAll indicates an expected call of RemoveAll
func (mr *MockCollectionMockRecorder) RemoveAll(selector interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockCollection)(nil).RemoveAll), selector)
}

// Bulk mocks base method
func (m *MockCollection) Bulk() Bulk {
	ret := m.ctrl.Call(m, "Bulk")
	ret0, _ := ret[0].(Bulk)
	return ret0
}

// Bulk indicates an expected call of Bulk
func (mr *MockCollectionMockRecorder) Bulk() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bulk", reflect.TypeOf((*MockCollection)(nil).Bulk))
}

// MockSession is a mock of Session interface
type MockSession struct {
	ctrl     *gomock.Controller
	recorder *MockSessionMockRecorder
}

// MockSessionMockRecorder is the mock recorder for MockSession
type MockSessionMockRecorder struct {
	mock *MockSession
}

// NewMockSession creates a new mock instance
func NewMockSession(ctrl *gomock.Controller) *MockSession {
	mock := &MockSession{ctrl: ctrl}
	mock.recorder = &MockSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSession) EXPECT() *MockSessionMockRecorder {
	return m.recorder
}

// Copy mocks base method
func (m *MockSession) Copy() *mgo.Session {
	ret := m.ctrl.Call(m, "Copy")
	ret0, _ := ret[0].(*mgo.Session)
	return ret0
}

// Copy indicates an expected call of Copy
func (mr *MockSessionMockRecorder) Copy() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockSession)(nil).Copy))
}

// Close mocks base method
func (m *MockSession) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockSessionMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSession)(nil).Close))
}

// MockQuery is a mock of Query interface
type MockQuery struct {
	ctrl     *gomock.Controller
	recorder *MockQueryMockRecorder
}

// MockQueryMockRecorder is the mock recorder for MockQuery
type MockQueryMockRecorder struct {
	mock *MockQuery
}

// NewMockQuery creates a new mock instance
func NewMockQuery(ctrl *gomock.Controller) *MockQuery {
	mock := &MockQuery{ctrl: ctrl}
	mock.recorder = &MockQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuery) EXPECT() *MockQueryMockRecorder {
	return m.recorder
}

// Iter mocks base method
func (m *MockQuery) Iter() Iter {
	ret := m.ctrl.Call(m, "Iter")
	ret0, _ := ret[0].(Iter)
	return ret0
}

// Iter indicates an expected call of Iter
func (mr *MockQueryMockRecorder) Iter() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iter", reflect.TypeOf((*MockQuery)(nil).Iter))
}

// All mocks base method
func (m *MockQuery) All(result interface{}) error {
	ret := m.ctrl.Call(m, "All", result)
	ret0, _ := ret[0].(error)
	return ret0
}

// All indicates an expected call of All
func (mr *MockQueryMockRecorder) All(result interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockQuery)(nil).All), result)
}

// One mocks base method
func (m *MockQuery) One(result interface{}) error {
	ret := m.ctrl.Call(m, "One", result)
	ret0, _ := ret[0].(error)
	return ret0
}

// One indicates an expected call of One
func (mr *MockQueryMockRecorder) One(result interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockQuery)(nil).One), result)
}

// Limit mocks base method
func (m *MockQuery) Limit(n int) Query {
	ret := m.ctrl.Call(m, "Limit", n)
	ret0, _ := ret[0].(Query)
	return ret0
}

// Limit indicates an expected call of Limit
func (mr *MockQueryMockRecorder) Limit(n interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Limit", reflect.TypeOf((*MockQuery)(nil).Limit), n)
}

// MockPipe is a mock of Pipe interface
type MockPipe struct {
	ctrl     *gomock.Controller
	recorder *MockPipeMockRecorder
}

// MockPipeMockRecorder is the mock recorder for MockPipe
type MockPipeMockRecorder struct {
	mock *MockPipe
}

// NewMockPipe creates a new mock instance
func NewMockPipe(ctrl *gomock.Controller) *MockPipe {
	mock := &MockPipe{ctrl: ctrl}
	mock.recorder = &MockPipeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPipe) EXPECT() *MockPipeMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockPipe) All(result interface{}) error {
	ret := m.ctrl.Call(m, "All", result)
	ret0, _ := ret[0].(error)
	return ret0
}

// All indicates an expected call of All
func (mr *MockPipeMockRecorder) All(result interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockPipe)(nil).All), result)
}

// Batch mocks base method
func (m *MockPipe) Batch(n int) Pipe {
	ret := m.ctrl.Call(m, "Batch", n)
	ret0, _ := ret[0].(Pipe)
	return ret0
}

// Batch indicates an expected call of Batch
func (mr *MockPipeMockRecorder) Batch(n interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Batch", reflect.TypeOf((*MockPipe)(nil).Batch), n)
}

// MockIter is a mock of Iter interface
type MockIter struct {
	ctrl     *gomock.Controller
	recorder *MockIterMockRecorder
}

// MockIterMockRecorder is the mock recorder for MockIter
type MockIterMockRecorder struct {
	mock *MockIter
}

// NewMockIter creates a new mock instance
func NewMockIter(ctrl *gomock.Controller) *MockIter {
	mock := &MockIter{ctrl: ctrl}
	mock.recorder = &MockIterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIter) EXPECT() *MockIterMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *MockIter) Next(result interface{}) bool {
	ret := m.ctrl.Call(m, "Next", result)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockIterMockRecorder) Next(result interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockIter)(nil).Next), result)
}

// Close mocks base method
func (m *MockIter) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockIterMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIter)(nil).Close))
}

// MockBulk is a mock of Bulk interface
type MockBulk struct {
	ctrl     *gomock.Controller
	recorder *MockBulkMockRecorder
}

// MockBulkMockRecorder is the mock recorder for MockBulk
type MockBulkMockRecorder struct {
	mock *MockBulk
}

// NewMockBulk creates a new mock instance
func NewMockBulk(ctrl *gomock.Controller) *MockBulk {
	mock := &MockBulk{ctrl: ctrl}
	mock.recorder = &MockBulkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBulk) EXPECT() *MockBulkMockRecorder {
	return m.recorder
}

// Upsert mocks base method
func (m *MockBulk) Upsert(pairs ...interface{}) {
	varargs := []interface{}{}
	for _, a := range pairs {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Upsert", varargs...)
}

// Upsert indicates an expected call of Upsert
func (mr *MockBulkMockRecorder) Upsert(pairs ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockBulk)(nil).Upsert), pairs...)
}

// Run mocks base method
func (m *MockBulk) Run() (*mgo.BulkResult, error) {
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(*mgo.BulkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockBulkMockRecorder) Run() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockBulk)(nil).Run))
}