package main

import "fmt"

func main() {
	var test = "my string"
	fmt.Println(test)

	test2 := "test 2"

	test2 = "test 3"

	fmt.Println(newFunc() + " " + test2)

}

func newFunc() string {
	return "test data"
}
