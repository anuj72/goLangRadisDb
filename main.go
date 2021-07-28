package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/locations", getLocatiopns).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))

}

func getLocatiopns(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	query := r.URL.Query()
	filters := query.Get("filters")
	fmt.Println(filters)
	var myquery = filters
	var LocalDb = getRedishData(myquery)
	if len(LocalDb) > 0 {
		json.NewEncoder(w).Encode(LocalDb)
	} else {
		var clientData = getClientData(myquery)
		setRedishData(clientData, myquery)
		fmt.Println("yyyy", clientData)
		json.NewEncoder(w).Encode(clientData)
	}
}
