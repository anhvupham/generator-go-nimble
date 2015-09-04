package main

import (
	"os"

	"gopkg.in/mgo.v2"
)
import "gopkg.in/boj/redistore.v1"

var (
	mdb       *mgo.Collection
	rdb       *redistore.RediStore
	port      int
	dbname    string
)

func init() {
	//Define port
	port = <%= port %>

	//Define database name
	dbname = "<%= dbname %>"

	// Connect to MongoDB
	session, err := mgo.Dial(<%= mongohost %>)
	if err != nil {
		panic(err)
	}
	session.SetSafe(&mgo.Safe{})
	mdb = session.DB(dbname).C(dbname)

	// Connect to Redis
	rdb, err = redistore.NewRediStore(10, "tcp", "<%= redishost %>:<%= redisport %>", "", []byte("<%= rediskey %>"))
	if err != nil {
		panic(err)
	}
}
