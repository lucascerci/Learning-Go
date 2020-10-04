package main

import "fmt"

func main() {
	fmt.Println(div(3, 0))
	fmt.Println(div(3, 5))
	demPanic()
}

func div(num1, num2 int) int {
	//u used recover with defer cause u want to recover from the situation that u created
	defer func() {
		fmt.Println(recover())
	}()
	//this function will run after all the errors have been created
	solution := num1 / num2

	return solution
}

func demPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("Panic")
}
