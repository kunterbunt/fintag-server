package model

import (
    "time"
    // "reflect"
)

type Model interface {
    AddCoordinates(*Coordinates) error
    UpdateCoordinates(*Coordinates) error
    GetCoordinates(id string) (*Coordinates, error)
    GetAllCoordinates() ([]*Coordinates, error)
    RemoveCoordinates(id string) error
}

type Coordinates struct {
  Id string `json:"id" bson:"_id"`
  Longitude int `json:"long" bson:"long"`
  Latitude int `json:"lat" bson:"lat"`
  Timestamp time.Time `json:"time" bson:"time"`
}
