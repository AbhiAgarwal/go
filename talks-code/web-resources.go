// Fetching web resources concurrently
// http://matt.aimonetti.net/posts/2012/11/27/real-life-concurrency-in-go/

package main

import (
	"fmt"
	"net/http"
	"time"
)

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

var urls = []string{
	"http://www.abhi.co/",
	"http://golang.org/",
	"http://google.com/apple/",
}

func asyncHttpGets(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse)
	responses := []*HttpResponse{}
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			resp, err := http.Get(url)
			ch <- &HttpResponse{url, resp, err}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		default:
			fmt.Printf(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
	return responses
}

func main() {
	results := asyncHttpGets(urls)
	for _, result := range results {
		fmt.Printf("%s status: %s\n", result.url,
			result.response.Status)
	}
}
