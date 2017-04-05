package main

import "fmt"

func main() {

	const GoldenRatio float64 = 1.618034

	fmt.Println(GoldenRatio)
	const (
		First = iota
		Second
		Third
	)

	fmt.Println(First, Second, Third)

}
