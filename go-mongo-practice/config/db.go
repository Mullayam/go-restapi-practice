package config

import "gopkg.in/mgo.v2"

type DB struct {
	Session *mgo.Session
}

func GetSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	return s
}
