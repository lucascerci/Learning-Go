package main

import (
	"fmt"
	"errors"
	"math"
)

func main() {
	//sum function
	//result := sum(2, 3)
	//fmt.Println(result);

	//sqrt function
	//so when we call sqrt function we got 2 returns, result and err
	result, err := sqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

//if u want to return something u need to pass the type after brackets
//this example is gonna take two ints and return one int
func sum(x int, y int) int {
	return x + y
}

//go functions can have multiples returns values
//in this example it can return a float64 or an error
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}