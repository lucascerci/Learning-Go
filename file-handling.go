package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//creating archive
	file, err := os.Create("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	//writing on archive (txt)
	file.WriteString("Hello world, inside a file that was created using GO!!")
	file.Close()

	//reading the txt
	stream, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	//convert stream into string
	s1 := string(stream)

	fmt.Println(s1)
}
