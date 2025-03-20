package models

// Model
type Album struct { //Template to store albums
	ID     string  `json:"id"` //Creating attribute ID to be public, string and set in json that this is id
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
