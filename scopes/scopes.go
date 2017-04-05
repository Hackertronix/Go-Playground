package main

import "fmt"

var iron = 24

func main() {

	gold := 48

	fmt.Println("Inside main")
	fmt.Println("Iron: ", iron)
	fmt.Println("Gold: ", gold)

	printSomething()
}

func printSomething() {
	fmt.Println("Inside printer function")
	fmt.Println("Iron: ", iron)
	//fmt.Println("Gold: ", gold)

}
