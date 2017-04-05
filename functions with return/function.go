package main

import "fmt"

func main() {

	result := getAwesomeResult()

	fmt.Println("Result is: ", result)

}

func getAwesomeResult() int {

	e := 64*98/76 + 9

	return e
}
