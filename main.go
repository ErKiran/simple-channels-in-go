package main

import (
	"fmt"
	"net/http"
)

type result struct {
	Error    error
	Response *http.Response
}

var urls = []string{
	"https://www.google.com",
	"https://www.facebook.com",
	"https://www.xainaw.com",
	"https://apple.com/",
	"https://youtube.com/",
	"https://www.onlinekhabar.com",
	"https://www.ekantipur.com",
	"https://www.merojobs.com",
	"https://www.twitter.com",
	"https://www.tiktok.com",
	"https://www.hoina.com",
}

func main() {
	checkStatus := func(done <-chan result, urls ...string) <-chan result {
		results := make(chan result)

		go func() {
			defer close(results)

			for _, url := range urls {
				var res result
				resp, err := http.Get(url)
				res = result{
					Error:    err,
					Response: resp,
				}
				select {
				case <-done:
					return
				case results <- res:
				}
			}
		}()
		return results
	}

	done := make(chan result)

	defer close(done)

	var success, failure int

	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			failure++
			continue
		}
		success++
	}

	fmt.Printf("Number of Successful Call Made to Website %d\n", success)
	fmt.Printf("Number of UnSuccessful Call Made to Website %d\n", failure)
}
