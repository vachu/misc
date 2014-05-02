package main

import "fmt"

func main() {
	in := make(chan interface{}, 1)
	out := make(chan interface{}, 1)

	go feedInput(in)
	go processInput(in, out)
	for output := range out {
		fmt.Println(output)
	}
}

func feedInput(in chan interface{}) {
	inputBulk := []string{
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
	}
	for _, input := range inputBulk {
		in <- input
	}
	in <- 12345
	in <- 1.23
	in <- uint32(23456789)
	close(in)
}

func processInput(in, out chan interface{}) {
	defer close(out)
	var ctr uint
	for input := range in {
		ctr++
		out <- fmt.Sprintf("%8d: %v (%T)", ctr, input, input)
	}
}
