package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type result struct {
	Error    error
	Response *http.Response
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
	"https://www.hoina.com",
}

func main() {
	done := PassData(websites)
	success, failure := ShowResult(done)
	fmt.Println("success 👌", success)
	fmt.Println("failure ❌", failure)
}

func ShowResult(done chan result) (s, f int) {
	var wg sync.WaitGroup
	var success, failure int
	for result := range done {
		wg.Add(1)
		if result.Error != nil {
			failure++
			continue
		}
		success++
	}
	return success, failure
}

func PassData(sites []string) (do chan result) {
	done := make(chan result)
	var wg sync.WaitGroup
	for _, site := range sites {
		wg.Add(1)
		go checkStatus(done, site, &wg)
	}

	// Wait for the pending operations and close chancel when done
	go func() {
		wg.Wait()
		close(done)
	}()

	return done
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
}
