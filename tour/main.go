package main

import (
	"fmt"
	"github.com/martijnvermaat/tour"
	"golang.org/x/tour/pic"
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

func main() {
	if len(os.Args) < 2 {
		doHelp()
	} else if fn, ok := dispatch[os.Args[1]]; !ok {
		doHelp()
	} else {
		fn()
	}
}
