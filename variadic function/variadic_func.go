package main

import "fmt"

func main() {

	r := add(8, 87, 90, 89, 11, 21)

	fmt.Println(r)

}

func add(nums ...int) int {

	var total int

	for _, n := range nums {
		total += n
	}

	fmt.Println("Total is: ", total)

	return total
}
