package main

import "fmt"

type Adder interface {
	add()
}

type Database struct {
}

func (d Database) add() {
	fmt.Println("Call from Database")
}

type InMemoryData struct {
}

func (d InMemoryData) add() {
	fmt.Println("Call from InMemoryData")
}

func createPersistence(d Adder) {
	d.add()
}

func main(){
	db := Database{}
	createPersistence(db)

}


