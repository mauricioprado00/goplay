package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan int, furl *fetchedUrls) {
	defer func() { ch <- (-1) }()
	if depth <= 0 || furl.canFetch(url) == false {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		ch <- 1
		go Crawl(u, depth-1, fetcher, ch, furl)
	}
	return
}

func main() {
	var waitc, ch = 1, make(chan int)
	var furl = &fetchedUrls{}

	furl.init()
	rand.Seed(time.Now().UTC().UnixNano())

	go Crawl("http://golang.org/", 4, fetcher, ch, furl)

	for waitc > 0 {
		val := <-ch
		waitc += val
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fetchedUrls struct {
	urls map[string]bool
	ch   *chan int
}

func (this *fetchedUrls) waitMutex() {
	<-*this.ch
}
func (this *fetchedUrls) releaseMutex() {
	*this.ch <- 1
}
func (this *fetchedUrls) canFetch(url string) bool {
	//wait the mutex
	this.waitMutex()
	defer this.releaseMutex()
	_, ok := this.urls[url]
	if ok {
		return false
	} else {
		this.urls[url] = true
	}
	return true
}

func (this *fetchedUrls) setCh(ch *chan int) {
	this.ch = ch
}

func (this *fetchedUrls) init() {
	c := make(chan int, 1)
	//this.ch = &c
	this.setCh(&c)
	this.releaseMutex()
	this.urls = make(map[string]bool)
}

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(50)))
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
