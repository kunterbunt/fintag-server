package model

import (
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
    "fmt"
    // "errors"
)

const (
    COLLECTION_COORDINATES = "coords"
)

/**
* Provides data access to a MongoDB.
* Implements Model interface.
*/
type MongoDb struct {
    session *mgo.Session
    databaseName string
}

func NewMongoDb(databaseName string) *MongoDb {
    fmt.Print("Connecting to mongodb... ")
    var driver MongoDb
    driver.databaseName = databaseName
    session, err := mgo.Dial("127.0.0.1")
    if err != nil {
		panic(err)
        session.Close()
    }
    driver.session = session
    driver.session.SetMode(mgo.Monotonic, true)
    fmt.Println("done.")
    return &driver
}

func (this *MongoDb) AddCoordinates(coords *Coordinates) error {
  coords.Id = bson.NewObjectId().Hex()
  return this.session.DB(this.databaseName).C(COLLECTION_COORDINATES).Insert(&coords)
}

func (this *MongoDb) UpdateCoordinates(coords *Coordinates) error {
  query := bson.M{"_id": coords.Id}
  change := bson.M{"$set": &coords}
  return this.session.DB(this.databaseName).C(COLLECTION_COORDINATES).Update(query, change)
}

func (this *MongoDb) GetCoordinates(id string) (coords *Coordinates, err error) {
  err = this.session.DB(this.databaseName).C(COLLECTION_COORDINATES).Find(bson.M{"_id": id}).One(&coords)
  return coords, err
}

func (this *MongoDb) GetAllCoordinates() (coordss []*Coordinates, err error) {
  err = this.session.DB(this.databaseName).C(COLLECTION_COORDINATES).Find(nil).All(&coordss)
  return coordss, err
}

func (this *MongoDb) RemoveCoordinates(id string) error {
  return this.session.DB(this.databaseName).C(COLLECTION_COORDINATES).Remove(bson.M{"_id": id})
}
