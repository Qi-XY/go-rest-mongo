package config

import (
	"gopkg.in/mgo.v2"
)

func GetMongoDB() (*mgo.Database, error) {

	session, err := mgo.Dial("mongodb://localhost:27018")
	if err != nil {
		return nil, err
	}

	db := session.DB("blockchain")
	return db, nil
}
