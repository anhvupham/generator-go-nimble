package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// json.Marshal body of request into interface
func parsebody(r *http.Request, v interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}

// searches collection for 'u' prior to upserting v
func addunique(u bson.M, v interface{}) (id string, err error) {
	var s interface{}
	err = mdb.Find(u).One(&s)
	if err == nil {
		return "", errors.New("record already exists")
	}

	var newId *mgo.ChangeInfo
	newId, err = mdb.Upsert(u, v)
	if err != nil {
		return "", err
	}
	bid, ok := newId.UpsertedId.(bson.ObjectId)
	if !ok {
		return "", errors.New("error getting object id")
	}
	id = bid.Hex()
	return id, nil
}

// Marshal "id" key into a bson.ObjectId
func parseid(r *http.Request) (id bson.ObjectId, err error) {
	v := mux.Vars(r)
	s := v["id"]
	if bson.IsObjectIdHex(s) {
		return bson.ObjectIdHex(s), nil
	} else {
		return bson.ObjectId(""), errors.New("invalid object id")
	}
}
