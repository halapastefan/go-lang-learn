package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	for url := range c {

		go func(u string) {
			time.Sleep(time.Second * 2)
			checkUrl(u, c)
		}(url)
	}

}

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		c <- url
		fmt.Println(url, " might be down")
		return
	}

	fmt.Println(url, " is up")
	c <- url

}
