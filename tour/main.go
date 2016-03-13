package main

import (
	"fmt"
	"github.com/martijnvermaat/tour"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/tree"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

var dispatch = map[string]func(){
	"sqrt":   doSqrt,
	"pic":    doPic,
	"wc":     doWordCount,
	"fib":    doFibonacci,
	"ipaddr": doIPAddr,
	"sqrt2":  doSqrt2,
	"read":   doMyReader,
	"rot13":  doRot13Reader,
	"image":  doImage,
	"trees":  doTrees,
	"crawl":  doCrawl,
}

func doHelp() {
	fmt.Println("Available subcommands:")
	for command := range dispatch {
		fmt.Printf("  %s\n", command)
	}
}

func doSqrt() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide a floating point number")
		return
	}

	if n, err := strconv.ParseFloat(os.Args[2], 64); err == nil {
		fmt.Printf("%f\n", tour.Sqrt(n))
	} else {
		fmt.Printf("Could not parse %q as float64\n", os.Args[2])
	}
}

func doPic() {
	pic.Show(tour.Pic)
}

func doWordCount() {
	b, _ := ioutil.ReadAll(os.Stdin)
	fmt.Println(tour.WordCount(string(b)))
}

func doFibonacci() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide the number of Fibonacci numbers to print")
		return
	}

	if n, err := strconv.Atoi(os.Args[2]); err == nil {
		f := tour.Fibonacci()
		for i := 0; i < n; i++ {
			fmt.Println(f())
		}
	} else {
		fmt.Printf("Could not parse %q as int\n", os.Args[2])
	}
}

func doIPAddr() {
	hosts := map[string]tour.IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func doSqrt2() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide a floating point number")
		return
	}

	if n, err := strconv.ParseFloat(os.Args[2], 64); err == nil {
		if sqrt, err := tour.Sqrt2(n); err == nil {
			fmt.Printf("%f\n", sqrt)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("Could not parse %q as float64\n", os.Args[2])
	}
}

func doMyReader() {
	b := make([]byte, 72)
	for {
		tour.MyReader{}.Read(b)
		fmt.Println(string(b))
	}
}

func doRot13Reader() {
	io.Copy(os.Stdout, tour.Rot13Reader{R: os.Stdin})
}

func doImage() {
	pic.ShowImage(tour.Image{X: 100, Y: 100})
}

func doTrees() {
	t1, t2 := tree.New(4), tree.New(4)

	if tour.Same(t1, t2) {
		fmt.Println("Binary trees t1 and t2 are equivalent :)")
	} else {
		fmt.Println("Binary trees t1 and t2 are not equivalent :x")
	}
}

func doCrawl() {
	sites := tour.Crawl("http://golang.org/", 4, fetcher)
	for url, body := range sites {
		fmt.Printf("Site body for %s: %q\n", url, body)
	}
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	fmt.Printf("Fetching %s\n", url)
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	fmt.Printf("-> not found: %s\n", url)
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
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

func main() {
	if len(os.Args) < 2 {
		doHelp()
	} else if fn, ok := dispatch[os.Args[1]]; !ok {
		doHelp()
	} else {
		fn()
	}
}
