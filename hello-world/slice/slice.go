package main

import "fmt"

func main() {

	test := []string{"test", str()}

	test = append(test, "appendStr")

	for i, str := range test {
		fmt.Println(i, str)
	}

	fmt.Println(test)

}

func str() string {
	return "funcStr"
}
