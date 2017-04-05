package main

import "fmt"

func main() {

	theString, result := getAwesomeResult()

	fmt.Println("String is: ", theString)
	fmt.Println("Int is: ", result)

}

func getAwesomeResult() (string, int) {

	e := 64*98/76 + 9
	str := "Hey There"

	return str, e
}
