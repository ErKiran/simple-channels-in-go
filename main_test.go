package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ExpectedOutCome struct {
	Successfull int
	Failure     int
}

func TestWebsite(t *testing.T) {
	websites := [][]string{
		{
			"https://www.google.com",
			"https://www.facebook.com",
			"https://www.xainaw.com",
			"https://apple.com/",
		},
		{
			"https://youtube.com/",
			"https://www.onlinekhabar.com",
			"https://www.ekantipur.com",
			"https://www.merojobs.com",
			"https://www.twitter.com",
		},
		{
			"https://www.merojobs.com",
			"https://www.twitter.com",
			"https://www.tiktok.com",
		},
		{
			"https://www.twitter.com",
			"https://www.tiktok.com",
			"https://www.hoina.com",
		},
	}

	results := []ExpectedOutCome{
		{Successfull: 3, Failure: 1},
		{Successfull: 5, Failure: 0},
		{Successfull: 3, Failure: 0},
		{Successfull: 2, Failure: 1},
	}

	for index, website := range websites {
		done := PassData(website)

		s, f := ShowResult(done)

		assert := assert.New(t)

		assert.Equal(s, results[index].Successfull, "ExpectedOutCome Successfull count should match the Result")
		assert.Equal(f, results[index].Failure, "ExpectedOutCome Failure count should match the Result")
	}
}
