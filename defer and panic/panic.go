package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	names, err := ioutil.ReadFile("names.txt")

	if err != nil {
		panic(err)
	}

	for _, e := range names {

		fmt.Println(string(e))
	}

}
