package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
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
	"https://www.hainaw.com", //This website takes forever to Load
}

var totalSitesToVisit = len(websites)

func main() {
	done := make(chan result)
	var wg sync.WaitGroup

	for _, site := range websites {
		wg.Add(1)
		go checkStatus(done, site, &wg)
	}

	var success, failure int

	// Wait for the pending operations and close chancel when done
	go func() {
		wg.Wait()
		close(done)
	}()

	for result := range done {
		wg.Add(1)
		if result.Error != nil {
			failure++
			continue
		}
		success++
		if totalSitesToVisit == 0 {
			fmt.Printf("Number of Successful Call Made to Website %d\n", success)
			fmt.Printf("Number of UnSuccessful Call Made to Website %d\n", failure)
			os.Exit(0)
		}
	}
}

func checkStatus(done chan result, url string, wg *sync.WaitGroup) {
	defer (*wg).Done()
	// If website takes more then 5 second then throw error

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)

	if err != nil {
		done <- result{Error: err, Response: resp}
	} else {
		done <- result{Error: nil, Response: resp}
	}
	totalSitesToVisit--
}
