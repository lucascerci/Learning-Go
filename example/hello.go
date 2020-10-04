package main

import ("fmt")

func main() {
	//first variable
	var x int
	//if dont provide any value it will return 0 for every type of var
	fmt.Println("var with no value = ", x)

	var h int = 5
	var y int = 7
	var sum int = h + y

	fmt.Println("sum", sum)

	q := 5
	p := 7
	soma := q + p

	fmt.Println("sum without defining var", soma)
}