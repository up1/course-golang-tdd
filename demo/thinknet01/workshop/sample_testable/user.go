package main

import (
  "fmt"
)

// Model
type User struct {
  id int
  fname string
  lname string
}

type Creator interface {
  CreateUser() User
}

func (u User) CreateUser() User {
  fmt.Println("Call mongodb")
  return User{}
}

// Service
func register(c Creator) {
  c.CreateUser()
}


func main() {
  register(User{})
}
