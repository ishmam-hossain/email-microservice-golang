package main


import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"./api"
	
)

func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/", api.HomePage).Methods("GET")
	go router.HandleFunc("/post", api.PostPage).Methods("POST")

	log.Fatal(http.ListenAndServe(":9999", router))
}