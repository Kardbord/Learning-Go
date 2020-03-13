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

type Crawler struct {
	knownURLs map[string]bool
	mu sync.RWMutex
}

func (crawler *Crawler) RLock() {
	crawler.mu.RLock()
}

func (crawler *Crawler) RULock() {
	crawler.mu.RUnlock()
}

func (crawler *Crawler) Lock() {
	crawler.mu.Lock()
}

func (crawler *Crawler) ULock() {
	crawler.mu.Unlock()
}

func (crawler *Crawler) containsURL(s string) bool {
	crawler.RLock()
	defer crawler.RULock()
	if _, ok := crawler.knownURLs[s]; ok {
		return true
	}
	return false
}

func (crawler *Crawler) insertURL(s string) {
	crawler.Lock()
	defer crawler.ULock()
	crawler.knownURLs[s] = true
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (crawler *Crawler) Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	if crawler.containsURL(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	crawler.insertURL(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(nextUrl string, d int, f Fetcher) {
			crawler.Crawl(nextUrl, d, f)
			wg.Done()
		}(u, depth-1, fetcher)
	}
	wg.Wait()
	return
}

func main() {
	crawler := Crawler{make(map[string]bool), sync.RWMutex{}}
	crawler.Crawl("https://golang.org/", 4, fetcher)
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
