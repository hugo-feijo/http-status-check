package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main()  {

	links := os.Args[1:]
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(5* time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink (link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	c <- link
	fmt.Println(link, "is up!")
}
