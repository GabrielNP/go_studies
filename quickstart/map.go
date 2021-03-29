package main

import "fmt"

func showMap() {
	object := make(map[string]string)

	object["name"] = "Gabriel"
	object["age"] = "24"

	fmt.Println(object)
}
