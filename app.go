package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	conf "github.com/aaroncowley/rest-api/config"
	cardsDao "github.com/aaroncowley/rest-api/dao"
)

var config = conf.Config{}
var dao = cardsDao.CardsDAO{}

// GetAllCards - Endpoint for getting all Cards from Database
func GetAllCards(w http.ResponseWriter, r *http.Request) {
	cards, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, cards)
}

// GetCardByName - Endpoint for Finding a single card by name
func GetCardByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	card, err := dao.FindByName(params["name"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Card Doesnt Exist Pal")
		return
	}
	respondWithJSON(w, http.StatusOK, card)
}

// ListOfCardNames - returns exactly that
func ListOfCardNames(w http.ResponseWriter, r *http.Request) {
	cards, err := dao.ListAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, cards)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithString(w http.ResponseWriter, code int, payload []string) {
	w.Write([]byte(strings.Join(payload, "\n")))
}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/cards", GetAllCards).Methods("GET")
	router.HandleFunc("/cards/{name}", GetCardByName).Methods("GET")
	router.HandleFunc("/cardlist", ListOfCardNames).Methods("GET")

	log.Fatal(http.ListenAndServe("192.168.217.52:4321", router))
}
