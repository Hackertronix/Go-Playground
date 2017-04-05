package main

import (
	"fmt"
)

func main() {

	phrase1 := "Lorem Ipsum"
	phrase2 := "Sit Amir"

	//Concatenate Strings
	fmt.Println(phrase1 + phrase2 + "Something else" + "Something totally else")

	//Slicing Strings

	aSliceOftext := phrase1[1:8]
	anotherSlice := phrase2[1:7]

	fmt.Println(aSliceOftext)
	fmt.Println(anotherSlice)
}
