package main

import "fmt"

func printMap(m map[string]string) {

	for english, marathi := range m {
		fmt.Printf("\n %v %v", english, marathi)
	}

}

func main() {

	colors := map[string]string{
		"red":   "laal",
		"green": "hirwa",
		"black": "kala",
	}

	fmt.Printf("colors %v \n", colors)

	//initialize empty map
	rang := make(map[string]string)

	rang["white"] = "pandhara"
	fmt.Printf("rang %v", rang)

	delete(colors, "red")

	printMap(colors)

}
