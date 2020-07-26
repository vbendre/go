package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type marathiBot struct{}

func (eb englishBot) getGreeting() string {
	return "Hello"
}

//alternate syntax omit values if not needed
func (marathiBot) getGreeting() string {
	return "namaskar"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func (eb englishBot) printGreeting() {
// 	fmt.Println(eb.getGreeting())
// }

// func (mb marathiBot) printGreeting() {
// 	fmt.Println(mb.getGreeting())
// }

func main() {

	mb := marathiBot{}
	// mb.printGreeting()

	eb := englishBot{}
	// eb.printGreeting()

	printGreeting(mb)
	printGreeting(eb)
}
