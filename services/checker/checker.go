package checker

import (
	"net/http"
	"time"
)

// Check checks if source alive
func Check(url string) bool {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode == 502 {
		return false
	}
	return true
}
