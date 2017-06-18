package main

import (
	"strconv"
	"fmt"
)

func main() {
	var data [10]string
	for i := 0; i<len(data); i++ {
		data[i] = strconv.Itoa(i+1)
	}

	for i, value := range data {
		fmt.Println(i, value)
	}

	x := [4]int {
		1,
		2,
		3,
		//4,
		5,
	}

	fmt.Println(len(x))

}




