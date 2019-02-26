package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/payload", ListenToGit).Methods("POST")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([[]byte("Hello World"))
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func ListenToGit(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(body)
}
