package main

import (
	"fmt"
)

func main() {
	//type definicion map, then keys, then value
	//to create a map u use the make function
	vertices := make(map[string]int)

	//setting key value pairs
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	fmt.Println(vertices)

	//gettin the value from a particular key
	fmt.Println(vertices["square"])

	//delete function to remove something
	delete(vertices, "triangle")
}