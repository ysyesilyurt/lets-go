package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type UrlCache struct {
	cache   map[string]error // a URL cache that maps URL:FetchStatus (FetchStatus is nil if no err while fetching, error o/w)
	mut sync.Mutex // to sync accesses to cache
}

func (c *UrlCache) PutUrl(url string, fetchStatus error) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.cache[url] = fetchStatus
	return
}

func (c *UrlCache) GetUrl(url string) (error, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()
	val, ok := c.cache[url]
	return val, ok
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	cache := UrlCache{cache: make(map[string]error)}
	CrawlHelper(url, depth, fetcher, &cache)

	fmt.Println("Fetching stats\n--------------")
	for url, err := range cache.cache {
		if err != nil {
			fmt.Printf("%v failed: %v\n", url, err)
		} else {
			fmt.Printf("%v was fetched\n", url)
		}
	}
}

func CrawlHelper(url string, depth int, fetcher Fetcher, cache *UrlCache) {
	if depth <= 0 {
		fmt.Printf("<- Done with %v, depth 0.\n", url)
		return
	} else if _, ok := cache.GetUrl(url); ok {
		fmt.Printf("<- Done with %v, already fetched.\n", url)
		return
	}
	
	body, urls, err := fetcher.Fetch(url)
	cache.PutUrl(url, err)

	if err != nil {
		fmt.Printf("<- Error on %v: %v\n", url, err)
		return
	}
	
	fmt.Printf("Found: %s %q\n", url, body)
	
	done_chan := make(chan bool)
	for i, u := range urls {
		fmt.Printf("-> Crawling child %v/%v of %v : %v.\n", i, len(urls), url, u)
		go func(url string) {
			CrawlHelper(url, depth-1, fetcher, cache)
			done_chan <- true
		}(u)
	}
	
	for i, u := range urls {
		fmt.Printf("<- [%v] %v/%v Waiting for child %v.\n", url, i, len(urls), u)
		<-done_chan
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
