package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	Id       int    `json:"Id"`
	ParentId int    `json:"ParentId"`
	KeyWord  string `json:"KeyWord"`
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/isUnique/{keyword}", isUnique)
	myRouter.HandleFunc("/registerNew/{keyword}/{params}", registerNew)
	myRouter.HandleFunc("/edit/{keyword}/{params}", isUnique)
	myRouter.HandleFunc("/delete/{keyword}", AddNewKeyWord)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func isUnique(w http.ResponseWriter, r *http.Request) {

}

func registerNew(w http.ResponseWriter, r *http.Request) {

}

func edit(w http.ResponseWriter, r *http.Request) {

}

func delete(w http.ResponseWriter, r *http.Request) {

}

func AddNewKeyWord(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(mux.Vars(r)["Id"])
	ParentId, _ := strconv.Atoi(mux.Vars(r)["ParentId"])
	registrationDb(Id, ParentId)
}

func main() {
	handleRequests()
}
