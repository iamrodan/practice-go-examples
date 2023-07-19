package waitgroups

import (
	"fmt"
	"net/http"
	"sync"
)

func SitesAvailabilityCheck(urls []string) {
	var wg sync.WaitGroup // better approach than sleep time

	for _, u := range urls {
		// incrementing the addgroup as waitgroup will wait until its value is 0
		wg.Add(1)
		go func(url string) {
			// Decrement the waitgroup value by one so after all completion the wait of wg completes
			defer wg.Done()
			checkUrl(url)
		}(u)
	}

	wg.Wait()
}

func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down !!!")
		return
	}
	fmt.Println(url, "is up and running.")
}
