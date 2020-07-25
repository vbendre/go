package main

import (
	"fmt"
	"strings"
)

type sliceString []string

// receiver function
func (sl sliceString) print() {
	for _, str := range sl {
		fmt.Println(str)
	}
}

func (sl sliceString) toString() string {

	return strings.Join(sl, ",")
}

func main() {
	sl := sliceString{"hello", "there", test()}
	sl = append(sl, "!")

	fmt.Println(sl[1:])

	sl.print()

	sl1, sl2 := spiltSlicedString(sl, 1)

	sl1.print()
	sl2.print()

	fmt.Println(sl.toString())

}

//function with multiple return values
func spiltSlicedString(sl sliceString, index int) (sliceString, sliceString) {
	return sl[:index], sl[index:]
}

func test() string {
	return "testString"
}
