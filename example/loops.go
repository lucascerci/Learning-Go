package main

import (
	"fmt"
)

func main() {
	//the only type of loop on go is the for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//while loop (delete the variable declaration from the beginning and the increments in the end) 
	//then u got the while loop
	ix := 0
	for ix < 5 {
		fmt.Println(ix)
		ix++
	}

	//loop with arrays
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	//loop with map
	m := make(map[string]string);
	m["a"] = "alpha"
	m["b"] = "beta"

	//but you use the key instead of index
	for key, value := range m {
		fmt.Println("index:", key, "value:", value)
	}
}