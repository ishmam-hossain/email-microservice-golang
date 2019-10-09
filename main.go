package main


import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"

	"github.com/gorilla/mux"
	"github.com/bitly/go-nsq"
	// "email-microservice/Utils/utils"
)


func nsqDataWriter() (interface{}) {
	nsqConfig := nsq.NewConfig()
	nsqWriter, err := nsq.NewProducer("127.0.0.1:4150", nsqConfig)
	if err != nil {
		return err
	}
	return nsqWriter
}

// NsqPublish publishes data to queue
func NsqPublish(topicName string, data []byte) error {
	nsqConfig := nsq.NewConfig()
	nsqWriter, err := nsq.NewProducer("127.0.0.1:4150", nsqConfig)
	if err != nil {
		return err
	}

	return nsqWriter.Publish(topicName, data)
}

// Message struct
type Message struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
}

// HomePage is the index route
func HomePage(w http.ResponseWriter, r *http.Request ) {
	fmt.Fprint(w, "Nothing here..!")
}

// PostPage takes some json and dumps to NSQ
func PostPage(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var msg Message
	err = json.Unmarshal(body, &msg)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	nsqErr := NsqPublish("api_to_nsq", []byte(response))
	if nsqErr != nil {
		http.Error(w, nsqErr.Error(), 500)
	}

	w.Header().Set("content-type", "application/json")
	w.Write(response)
}

func main()  {
	// functions.Hello()
	router := mux.NewRouter()
	router.HandleFunc("/", HomePage).Methods("GET")
	go router.HandleFunc("/post", PostPage).Methods("POST")

	log.Fatal(http.ListenAndServe(":9999", router))
}