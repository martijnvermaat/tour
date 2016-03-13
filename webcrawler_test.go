package tour

import (
	"fmt"
	"sync"
	"testing"
)

func TestCrawl(t *testing.T) {
	expected := map[string]string{
		"http://golang.org/pkg/os/":  "Package os",
		"http://golang.org/":         "The Go Programming Language",
		"http://golang.org/pkg/":     "Packages",
		"http://golang.org/pkg/fmt/": "Package fmt",
	}
	sites := Crawl("http://golang.org/", 4, newTestFetcher(t, testResults))

	for url, body := range expected {
		if v, ok := sites[url]; !ok {
			t.Errorf("Expected url %s was not crawled", url)
		} else if v != body {
			t.Errorf("Crawled body for url %s was %q, expected %q", url, v, body)
		}
	}

	for url := range sites {
		if _, ok := expected[url]; !ok {
			t.Errorf("Crawled url %s was not expected", url)
		}
	}
}

// testFetcher is Fetcher that returns canned results and fails when a url is
// fetched more than once.
type testFetcher struct {
	results map[string]*testResult
	fetched map[string]bool
	mux     sync.Mutex
	t       *testing.T
}

type testResult struct {
	body string
	urls []string
}

func (f *testFetcher) Fetch(url string) (string, []string, error) {
	f.mux.Lock()
	if _, ok := f.fetched[url]; ok {
		f.t.Errorf("Fetching already fetched url: %s", url)
	}
	f.fetched[url] = true
	f.mux.Unlock()

	if res, ok := f.results[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// create a populated testFetcher.
func newTestFetcher(t *testing.T, results map[string]*testResult) *testFetcher {
	return &testFetcher{
		results: results,
		fetched: map[string]bool{},
		t:       t,
	}
}

var testResults = map[string]*testResult{
	"http://golang.org/": &testResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &testResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &testResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &testResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
