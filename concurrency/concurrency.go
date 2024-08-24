package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

type WebsiteChecker func(string) bool
type result struct {
	url    string
	exists bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		tempUrl := url
		go func(u string) {
			resultChannel <- result{url: u, exists: wc(u)}
		}(tempUrl)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.url] = r.exists
	}

	return results
}

const tenSecondsTimeout = 10

func NewRacer(a, b string) (winner string, error error) {
	return configurableRacer(a, b, tenSecondsTimeout)
}

func configurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
