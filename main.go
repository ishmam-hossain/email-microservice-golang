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
	router.HandleFunc("/post", api.PostPage).Methods("POST")
	
	// test routes
	router.HandleFunc("/emails", api.ArrayTest).Methods("POST")
	router.HandleFunc("/test1", api.Test1).Methods("GET")
	router.HandleFunc("/test2", api.Test2).Methods("GET")

	log.Fatal(http.ListenAndServe(":9999", router))
	functions.KillNSQ()
}