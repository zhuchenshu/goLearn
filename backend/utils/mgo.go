package utils

import (
	"github.com/globalsign/mgo"
	"time"
)

type Person struct {
	Name  string
	Phone string
}

type dbServer struct {
	url     string
	session *mgo.Session
}

var server *dbServer

func init() {
	server = &dbServer{
		url:     "47.106.145.145:27017",
		session: initSession("47.106.145.145:27017"),
	}
}

func initSession(url string) *mgo.Session {
	session, err := mgo.DialWithTimeout(url, 10*time.Second) //timeout
	if err != nil {
		panic(err) //panic
	}

	session.SetSocketTimeout(10 * time.Minute) //timeout
	return session
}

func do(databaseName, collection string, session *mgo.Session, f func(*mgo.Collection) error) error {
	sessionCopy := session.Clone()
	defer func() {
		sessionCopy.Close()
	}()
	c := sessionCopy.DB(databaseName).C(collection)
	return f(c)
}

func (server *dbServer) withCollectionInner(databaseName, collection string, f func(*mgo.Collection) error) error {
	if server.session == nil {
		server.session = initSession(server.url)
	}

	return do(databaseName, collection, server.session, f)
}

func WithCollection(databaseName, collection string, f func(*mgo.Collection) error) error {
	if server != nil {
		return server.withCollectionInner(databaseName, collection, f)
	}

	return NewError(0, "")
}

type Error struct {
	errorCode int
	errorDes  string
}

func NewError(code int, desc string) *Error {
	return &Error{errorCode: code, errorDes: desc}
}

func (errorCode *Error) Error() string {
	return errorCode.errorDes
}
