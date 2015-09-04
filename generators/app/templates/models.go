package main

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id          bson.ObjectId `json:"id bson:"_id,omitempty"`
	Email       string        `json:"email"`
	Hash        string        `json:"password"`
	Firstname   string        `json:"firstname"`
	Lastname    string        `json:"lastname"`
	Company     string        `json:"company"`
	Masteradmin bool          `json:"masteradmin"`
	Logincount  int           `json:"logincount"`
}
