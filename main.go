package main


import (
	"log"
	"net/http"

	"./api"
	"./utils"

	"github.com/gorilla/mux"
)

func main()  {
	functions.InitNSQ()

	router := mux.NewRouter()
	router.HandleFunc("/", api.HomePage).Methods("GET")
	go router.HandleFunc("/post", api.PostPage).Methods("POST")
	go router.HandleFunc("/emails", api.ArrayTest).Methods("POST")

	log.Fatal(http.ListenAndServe(":9999", router))
	functions.KillNSQ()
}