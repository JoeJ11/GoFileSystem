package main

import (
	// "log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// var greeting_handler GreetingHandler
	// http.Handle("/greeting", greeting_handler)
	r := mux.NewRouter()
	r.HandleFunc("/greeting", GreetingHandler)
	r.HandleFunc("/kv/insert", HandleInsert).Methods("POST")
	r.HandleFunc("/kv/delete", HandleDelete).Methods("POST")
	r.HandleFunc("/kv/get", HandleGet).Methods("GET")
	r.HandleFunc("/kv/update", HandleUpdate).Methods("POST")
	r.HandleFunc("/kvman/countkey", HandleCountKey).Methods("GET")
	r.HandleFunc("/kvman/dump", HandleDump).Methods("GET")
	r.HandleFunc("/kvman/shutdown", HandleShutdown).Methods("GET")

	http.ListenAndServe("localhost:4000", r)
}
