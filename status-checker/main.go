package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://golang.org",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		// fmt.Println("checking link --" + link)
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }
	// alternate syntax
	for {
		fmt.Println(<-c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		// fmt.Println(link + " might be down!!")
		c <- link + " might be down!!"
		return
	}

	// fmt.Println(link + " link is working!!")

	c <- link + " link is working!!"
}
