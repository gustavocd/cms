package main

import (
	"github.com/gobuffalo/pop"
	"github.com/gorilla/mux"
	"github.com/gustavocd/cms/pkg/server"
	"log"
	"net/http"
)

func main() {
	conn, err := pop.Connect("development")
	if err != nil {
		log.Fatalf("could not connect to the database %v", err)
		return
	}

	r := mux.NewRouter()
	svr := server.NewServer(conn, r)
	svr.Router.HandleFunc("/api/v1/pages", svr.HandlePagesGetAll()).Methods("GET")
	svr.Router.HandleFunc("/api/v1/pages/{id}", svr.HandlePagesGet()).Methods("GET")
	svr.Router.HandleFunc("/api/v1/pages", svr.HandlePagesCreate()).Methods("POST")
	svr.Router.HandleFunc("/api/v1/pages/{id}", svr.HandlePagesDelete()).Methods("DELETE")
	svr.Router.HandleFunc("/api/v1/pages/{id}", svr.HandlePagesUpdate()).Methods("PUT")

	log.Printf("Server is running on http://localhost:7777 🐹")
	err = http.ListenAndServe(":7777", svr.Router)
	if err != nil {
		log.Fatalf("could not run the server %v", err)
		return
	}
}
