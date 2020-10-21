package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

type result struct {
	Error    error          `json:"error"`
	Response *http.Response `json:"response"`
}

var websites = []string{
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
	// "https://www.hoinawwww.com",
}

func main() {
	done := make(chan result)
	var wg sync.WaitGroup

	for _, site := range websites {
		wg.Add(1)
		go checkStatus(done, site, &wg)
	}

	var success, failure, count int

	go func() {
		wg.Wait()
		close(done)
	}()

	for result := range done {
		count++
		wg.Add(1)
		if result.Error != nil {
			failure++
			continue
		}
		success++

		if count == len(websites) {
			fmt.Printf("Number of Successful Call Made to Website %d\n", success)
			fmt.Printf("Number of UnSuccessful Call Made to Website %d\n", failure)
			os.Exit(0)
		}
	}
}

func checkStatus(done chan result, url string, wg *sync.WaitGroup) {
	defer (*wg).Done()

	resp, err := http.Get(url)

	if err != nil {
		done <- result{Error: err, Response: resp}
	} else {
		done <- result{Error: nil, Response: resp}
	}
}
