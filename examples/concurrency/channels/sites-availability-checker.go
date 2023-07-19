package channels

import (
	"fmt"
	"net/http"
)

func SitesAvailabilityCheck(urls []string) {
	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	results := make([]string, len(urls))

	for i, _ := range results {
		response := <-c
		fmt.Println(response)
		results[i] = response
	}

}

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("%v is down", url)
	} else {
		c <- fmt.Sprintf("%v is up and running", url)
	}

}
