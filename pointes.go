package main

import ("fmt")

func main() {
	i := 7
	inc(i)
	//if we pass just an i and inc doest return anything so the original value is not modified
	fmt.Println(i);

	//if however we pass a point to the variable then the function is gonna be able to look at
	//the memory reference and modify the original version
	j := 7
	incP(&j)
	fmt.Println(j);
	
}

func inc(x int) {
	x++
}

//with point
func incP(x *int) {
	*x++
}
