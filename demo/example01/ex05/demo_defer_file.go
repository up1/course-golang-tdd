package main

import (
	"os"
	"fmt"
)

func main() {
	f, err := os.Open("demo_defer_file.go")
	defer f.Close()

	if err != nil {
		fmt.Println("Start...")
	}
}
