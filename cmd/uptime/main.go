package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	urlString := flag.String("urls", "https://ma.rkus.io, https://www.uberswe.com", "URLs you would like to check the status code of")
	interval := flag.Int("interval", 60, "The number of seconds between each request")
	flag.Parse()
	if urlString == nil {
		panic("please specify one or more URLs to query")
	}
	if interval == nil || *interval < 1 {
		panic("please specify an interval greater or equal to 1")
	}
	urls := strings.Split(*urlString, ",")
	for _, url := range urls {
		go checkStatus(url)
	}
	for range time.Tick(time.Duration(*interval) * time.Second) {
		for _, url := range urls {
			go checkStatus(url)
		}
	}
}

// checkStatus performs a get request to the specified url and logs the status code or error
func checkStatus(url string) {
	url = strings.TrimSpace(url)
	resp, err := http.Get(url)
	if err != nil {
		if !(resp != nil && resp.StatusCode > 0) {
			log.Printf("%s - %v\n", url, err)
			return
		}
	}
	// check the status code
	log.Printf(" %s - %d\n", url, resp.StatusCode)
}
