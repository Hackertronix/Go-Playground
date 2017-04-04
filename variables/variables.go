package main

import "fmt"

func main() {

	var myAge int = 21

	//default values
	var something int
	var truth bool

	// brief assignment

	myVariable := 24

	//multiple assignment

	myNewAge, myNewVariable := 35, 86

	fmt.Println(myAge)
	fmt.Println(myVariable)

	//modified

	fmt.Println(myNewAge)
	fmt.Println(myNewVariable)
	fmt.Println(truth)
	fmt.Println(something)
}
