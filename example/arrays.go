package main

import (
	"fmt"
)

func main() {
	//hold five ints
	var a [5]int
	//all of the elements will be initialized with 0 values
	fmt.Println(a)

	//setting value of an index
	a[2] = 7
	fmt.Println(a)

	//inicializing the content of a array more easily
	//drop the var
	//int the end of the line open the brackets with the values
	b := [5]int{5,4,3,2,1}
	fmt.Println(b)

	//so arrays with fixed lenght sound incovenient, well, they are, if i wanted to put 6 element
	//out of the b array i cant because the lenght of the array is part of the array type
	//to overcome this its possible to use slices, slices are a abstraction of the top of arrays to make
	//them easily to work with, they dont have a fix number of lenght
	
	//slice example
	//just remove the fixed lenght
	c := []int{5,4,3,2,1}
	//now we can use append function
	//append doesnt modify the old array, create a new one 
	c = append(c, 13)
	fmt.Println(c)
}