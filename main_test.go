package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebsite(t *testing.T) {
	websites := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.xainaw.com",
		"https://apple.com/",
	}
	done := PassData(websites)

	s, f := ShowResult(done)

	assert := assert.New(t)

	assert.Equal(s, 3, "Three call should be successfull")
	assert.Equal(f, 1, "Only one http call should be unsuccessfull")
}
