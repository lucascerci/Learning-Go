package main

import ("fmt")

//struc is a collection of fields so its possible group things together to create more logical type
type person struct {
	//it can be anything, it can be another struct too
	name string
	age int
}

func main() {
	p := person{name: "Lucas", age: 20}

	//print all struct
	fmt.Println(p);

	//print just a specific field
	fmt.Println(p.name)
}