package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("mongo18.183:9800")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("bucket0001").C("examobj0001")
	err = c.Insert(&Person{"Ale", "+86 10 8000 8000"},
		&Person{"Cla", "+86 10 6000 6000"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone: ", result.Phone)
}
