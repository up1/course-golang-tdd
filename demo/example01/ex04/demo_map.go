package main

import "fmt"

func main() {
	x := make(map[string]string)
	x["firstname"] = "Somkiat"
	x["lastname"] = "Puisungoen"
	x["gender"] = "Male"

	fmt.Println(x["firstname"])
	fmt.Println(x["lastname"])
	fmt.Println(x["gender"])
	fmt.Println(x)

	for key, value := range x {
		fmt.Println(key, value)
	}

	data, ok := x["not found"]
	fmt.Println(data, ok)

	if x["firstname"] == "" {
		fmt.Println(x["firstname"])
	}

	if data, ok := x["firstname"]; ok {
		fmt.Println(data, ok)
	}

	delete()
}
