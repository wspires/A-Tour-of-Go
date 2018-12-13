// https://tour.golang.org/concurrency/10
// In this exercise you'll use Go's concurrency features to parallelize a web crawler.
//
// Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.
//
// Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use!
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

// Thread-safe map of string -> int.
type SafeCounter struct {
    v   map[string]int
    mux sync.Mutex
}
var fetched_urls SafeCounter = SafeCounter{v: make(map[string]int)}

// Lock and increment map value for key, returning the new value.
func (c *SafeCounter) ValueInc(key string) int {
    c.mux.Lock()
    defer c.mux.Unlock()
    c.v[key]++
    return c.v[key]
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
    if depth <= 0 {
        return
    }

    // Fetch URLs in parallel.
    done := make(chan bool)
	go func() {
		// Defer setting done channel to always flag as done regardless of when we return.
		defer func() {
			done <- true
		}()

		// Don't fetch the same URL twice.
		// ValueInc increments count for url and returns it, so will be 1 the first time it's seen.
		if fetched_urls.ValueInc(url) != 1 {
			//fmt.Println(fetched_urls.Value(url))
			return
		}

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			Crawl(u, depth - 1, fetcher)
		}
	}()

	// Wait for fetch.
	<-done
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

