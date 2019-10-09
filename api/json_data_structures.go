package api

// Message struct
type Message struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
}

// Email struct
type Email struct {
	Recipient []string `json:"recipient"`
	Email string `json:"email"`
	Age int `json:"age"`
}