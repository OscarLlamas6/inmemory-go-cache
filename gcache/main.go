package main

import (
	"time"

	"log"
	"net/http"

	"github.com/OscarLlamas6/inmemory-go-cache/gcache/pkg/handler"
	"github.com/OscarLlamas6/inmemory-go-cache/gcache/pkg/store"
	"github.com/bluele/gcache"
	"github.com/gorilla/mux"
)

func main() {

	db := map[string]int{
		"Oscar":    1410,
		"Vanessa":  1420,
		"Alfredo":  1290,
		"Josselyn": 1450,
	}

	store := store.NewStore(db, gcache.New(30).LFU().Build(), time.Hour)
	sh := handler.NewScoreHandler(store)
	r := mux.NewRouter()
	r.HandleFunc("/scores/{student}", sh.HandleGet).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
