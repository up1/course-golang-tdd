package model

type User struct {
	Id        int    `json:"user_id"`
	Firstname string `json:"fname"`
	Lastname  string `json:"lname"`
	Age       int    `json:"age"`
}
