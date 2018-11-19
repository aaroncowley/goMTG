package models

import "gopkg.in/mgo.v2/bson"

// Full mgo struct replica of MtgJson Card
type Card struct {
    ID              bson.ObjectId   `bson:"_id" json:"id"`
    Layout          string          `bson:"layout" json:"layout"`
    Name            string          `bson:"name" json:"name"`
    ManaCost        string          `bson:"manaCost" json"manaCost"`
    Cmc             string          `bson:"cmc" json:"cmc"`
    Colors          []string        `bson:"colors" json:"colors"`
    Type            string          `bson:"type" json:"type"`
    Types           []string        `bson:"types" json:"types"`
    Text            string          `bson:"text" json:"text"`
    ImageName       string          `bson:"imageName" json:"imageName"`
    ColorIdentity   []string        `bson:"colorIdentity" json:"colorIdentity"`
}

