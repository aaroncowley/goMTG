package mtgcard

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

type CardsMongo struct {
    Server string
    Database string
}

type Card struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
}

var db *mgo.Database

const (
    COLLECTION = "cards"
)

func (c *CardsMongo) Connect(){
    session, err := mgo.Dial(c.Server)
    if err != nil {
        log.Fatal(err)
    }
    db = session.DB(c.Database)
}

func (c *CardsMongo) FindAll() ([]Card, error) {
	var cards []Card
	err := db.C(COLLECTION).Find(bson.M{}).All(&cards)
	return cards, err
}

func (c *CardsMongo) FindByName(name string) (card, error) {
	var card Card
    err := db.C(COLLECTION).Find.M(bson.M{ "name" : name }).One(&card)
	return card, err
}

func GetAllCards(w http.ResponseWriter, req *http.Request) {
    cards, err := mtgcard.FindAll()
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJson(w, http.StatusOK, cards)
}

func GetSingleCard(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(r)
	card, err := mtgcard.FindByName(params["name"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Card Doesnt Exist Pal")
		return
	}
	respondWithJson(w, http.StatusOK, card)
}

func main() {
    router := mux.NewRouter()
	router.HandleFunc("/cards", GetAllCards).Methods("GET")
	router.HandleFunc("/cards/{name}", GetSingleCard).Methods("GET")

	log.Fatal(http.ListenAndServe(":4321", router))
}
