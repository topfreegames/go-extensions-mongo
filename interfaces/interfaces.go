package interfaces

import (
	"context"

	"github.com/globalsign/mgo"
)

//MongoDB represents the contract for a Mongo DB
type MongoDB interface {
	Run(cmd interface{}, result interface{}) error
	C(name string) (Collection, Session)
	Close()
	WithContext(ctx context.Context) MongoDB
}

//Collection represents a mongoDB collection
type Collection interface {
	Find(query interface{}) Query
	FindId(id interface{}) Query
	Pipe(pipeline interface{}) Pipe
	Insert(docs ...interface{}) error
	UpsertId(id interface{}, update interface{}) (*mgo.ChangeInfo, error)
	Upsert(selector interface{}, update interface{}) (*mgo.ChangeInfo, error)
	RemoveId(id interface{}) error
	Remove(selector interface{}) error
	RemoveAll(selector interface{}) (*mgo.ChangeInfo, error)
	Bulk() Bulk
}

//Session is the mongoDB session
type Session interface {
	Copy() *mgo.Session
	Close()
}

//Query wraps mongo Query
type Query interface {
	Iter() Iter
	All(result interface{}) error
	One(result interface{}) error
	Limit(n int) Query
}

//Pipe wraps mongo Pipe
type Pipe interface {
	All(result interface{}) error
	Batch(n int) Pipe
}

//Iter wraps mongo Iter
type Iter interface {
	Next(result interface{}) bool
	Close() error
}

// Bulk contains methods to be executed at one at server
type Bulk interface {
	Upsert(pairs ...interface{})
	Run() (*mgo.BulkResult, error)
}