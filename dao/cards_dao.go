package dao

import (
    "log"

    . "github.com/aaroncowley/rest-api/models"
    mgo "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type CardsDAO struct {
    Server string
    Database string
}

var db *mgo.Database

const (
    COLLECTION = "cards"
)

func (c *CardsDAO) Connect() {
    session, err := mgo.Dial(c.Server)
    if err != nil {
        log.Fatal(err)
    }
    db = session.DB(c.Database)
}

func (c *CardsDAO) FindAll() ([]Card, error) {
    var cards []Card
    err := db.C(COLLECTION).Find(bson.M{}).All(&cards)
    return cards, err
}

func (c *CardsDAO) FindById(id string) (Card, error) {
    var card Card
    err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&card)
    return card, err
}

func (c *CardsDAO) FindByName(name string) (Card, error) {
    var card Card
    err := db.C(COLLECTION).Find(bson.M{"name" : name}).One(&card)
    return card, err
}

func (c *CardsDAO) ListAll() ([]string, error) {
    var list []string
    err := db.C(COLLECTION).Find(bson.M{}).Select("name").Sort("name").All(&list)
    return list, err
}

//TODO: Insert
//TODO: Update
//TODO: Delete


