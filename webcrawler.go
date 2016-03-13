// https://tour.golang.org/concurrency/10

package tour

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetch Fetcher) map[string]string {
	r := newResult()
	r.Take(url)
	go crawl(url, depth, fetch, r)
	r.Wait()
	return r.Sites
}

// Synchronized result container for parallel fetching.
type result struct {
	Sites  map[string]string // Body for every fetched url.
	Errors map[string]bool   // Not found urls.
	taken  map[string]bool   // To be fetched urls.
	mux    sync.Mutex        // Synchronize all access.
	done   chan bool         // Gets a message when no more urls are pending.
}

// To be called before fetching a url.
func (r *result) Take(url string) bool {
	r.mux.Lock()
	defer r.mux.Unlock()
	if _, ok := r.taken[url]; ok {
		return false
	}
	if _, ok := r.Sites[url]; ok {
		return false
	}
	if _, ok := r.Errors[url]; ok {
		return false
	}
	r.taken[url] = true
	return true
}

func (r *result) give(url string) {
	if _, ok := r.taken[url]; !ok {
		panic(fmt.Sprintf("Cannot give back url %s: not taken", url))
	}
	delete(r.taken, url)
	if len(r.taken) == 0 {
		r.done <- true
	}
}

// Undo a Take() call.
func (r *result) Give(url string) {
	r.mux.Lock()
	r.give(url)
	r.mux.Unlock()
}

// Add a fetch result.
func (r *result) Add(url, body string) {
	r.mux.Lock()
	r.Sites[url] = body
	r.give(url)
	r.mux.Unlock()
}

// Add a fetch error.
func (r *result) Error(url string) {
	r.mux.Lock()
	r.Errors[url] = true
	r.give(url)
	r.mux.Unlock()
}

// Blocks until there are no pending fetches.
func (r *result) Wait() {
	<-r.done
}

func newResult() *result {
	r := &result{}
	r.Sites = make(map[string]string)
	r.Errors = make(map[string]bool)
	r.taken = make(map[string]bool)
	r.done = make(chan bool)
	return r
}

// crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func crawl(url string, depth int, fetcher Fetcher, r *result) {
	if depth <= 0 {
		r.Give(url)
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		r.Error(url)
		return
	}
	for _, u := range urls {
		if r.Take(u) {
			go crawl(u, depth-1, fetcher, r)
		}
	}
	r.Add(url, body)
}
