package main

import "fmt"

func main() {
	in, out := make(chan string, 1), make(chan string, 1)
	go concurrent(in, out)

	ctr := 5
	for str := range out {
		switch {
		default:
			fmt.Println(str)
		case ctr == 0:
			close(in)
		}
		ctr--
	}
	fmt.Println("quitting...")
}

func concurrent(in, out chan string) {
	defer close(out)

	for i := 0; ; i++ {
		select {
		case _, isOpen := <-in:
			if !isOpen {
				return
			}
		default:
			out <- fmt.Sprintf("%d: Hi There!", i+1)
		}
	}
}
