package main

import (
	"fmt"
	"github.com/martijnvermaat/tour"
	"golang.org/x/tour/pic"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	switch os.Args[1] {
	case "sqrt":
		if n, err := strconv.ParseFloat(os.Args[2], 64); err == nil {
			fmt.Printf("%f\n", tour.Sqrt(n))
		} else {
			fmt.Printf("Could not parse %q as float64\n", os.Args[2])
		}
	case "pic":
		pic.Show(tour.Pic)
	case "wc":
		b, _ := ioutil.ReadAll(os.Stdin)
		fmt.Println(tour.WordCount(string(b)))
	case "fib":
		if n, err := strconv.Atoi(os.Args[2]); err == nil {
			f := tour.Fibonacci()
			for i := 0; i < n; i++ {
				fmt.Println(f())
			}
		} else {
			fmt.Printf("Could not parse %q as int\n", os.Args[2])
		}
	default:
		fmt.Println("Available subcommands: sqrt, pic, wc, fib")
	}
}
