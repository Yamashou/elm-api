package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	var db *gorm.DB
	db, err := gorm.Open("mysql", "user:pass@tcp(mysql:3306)/db?parseTime=true")
	if err != nil {
		panic(err)
	}
	q := Query{db}
	defer q.db.Close()

	r := mux.NewRouter()

	r.Handle("/", UseContextGet(http.HandlerFunc(q.Index))).
		Methods(http.MethodGet)
	r.HandleFunc("/list", q.List).
		Methods(http.MethodGet)
	r.Handle("/create", UseContextPost(http.HandlerFunc(q.CreatExperiment))).
		Methods(http.MethodPost)
	r.Handle("/add", UseContextPost(http.HandlerFunc(q.AddFeatureData))).
		Methods(http.MethodPost)
	r.Handle("/learn", UseContextPost(http.HandlerFunc(q.Learning))).
		Methods(http.MethodPost)
	r.Handle("/fetch", UseContextPost(http.HandlerFunc(q.Fetch))).
		Methods(http.MethodPost)

	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

	return
}
