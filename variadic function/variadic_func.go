package main

import "fmt"

func main() {

	add(8, 87, 90, 89, 11, 21)

}

func add(nums ...int) {

	var total int

	for _, n := range nums {
		total += n
	}

	fmt.Println("Total is: ", total)
}
