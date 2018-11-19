package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/aaroncowley/rest-api/config"
	. "github.com/aaroncowley/rest-api/dao"
)

var config = Config{}
var dao = CardsDAO{}

func GetAllCards(w http.ResponseWriter, r *http.Request) {
    cards, err := dao.FindAll()
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJson(w, http.StatusOK, cards)
}

func GetCardByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	card, err := dao.FindByName(params["name"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Card Doesnt Exist Pal")
		return
	}
	respondWithJson(w, http.StatusOK, card)
}

func ListOfCardNames(w http.ResponseWriter, r *http.Request) {
    cards, err := dao.ListAll()
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJson(w, http.StatusOK, cards)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
    respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}
func respondWithList(w http.ResponseWriter, code int, payload interface{}) {
    response := payload
    w.Header().Set("Content-Type", "application/json")
    w.WriterHeader(code)
    w.Write(response)
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
