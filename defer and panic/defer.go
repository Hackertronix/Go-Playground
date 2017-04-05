package main

import "fmt"

func main() {

	defer doSomething()
	doSomethingElse()

}

func doSomething() {
	fmt.Println("Doing something")
}

func doSomethingElse() {

	fmt.Println("Doing Something Else....")
}
