package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main() {

	var g errgroup.Group
	var urls = []string{
		"http://www.baidu.com/",
		"https://go-zh.org/ref/mem",
	}
	for i := range urls {
		url := urls[i]
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Sprintf("error fetched all URLs.%w", err)
	}

}
