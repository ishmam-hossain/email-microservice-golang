package api

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../utils"

)

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

	nsqErr := functions.NsqPublish("api_to_nsq", []byte(response))
	if nsqErr != nil {
		http.Error(w, nsqErr.Error(), 500)
	}

	w.Header().Set("content-type", "application/json")
	w.Write(response)
}