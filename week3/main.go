package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	//g := new(errgroup.Group)
	var urls = []string{
		"https://www.baidu.com/",
		"http://www.bbbbbbbbbbbbbbbb.com/",
		"https://www.bilibili.com/",
	}
	g, ctx := errgroup.WithContext(context.Background())
	respChan := make(chan *http.Response, 10)

	for _, url := range urls {

		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			} else {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case respChan <- resp:
				}
			}
			return err
		})

		// Wait for all HTTP fetches to complete.
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Println("Error: ", err)
	}
}
