package main

import "fmt"
import "log"

func Node(fn func(interface{}) interface{}, in, out chan interface{}) {
	defer close(out)

	for {
		select {
		case inp, isOpen := <-in:
			if !isOpen {
				return
			} else if inp != nil {
				out <- fn(inp)
			}
		default:
		}
	}
}

func feedInput(in chan interface{}) {
	defer close(in)

	in <- "First"
	in <- "Second"
	in <- "Third"
	in <- 1234
}

func main() {
	in, out := buildChain(
		Process1,
		//Process2,
		Process3,
	)
	if in == nil || out == nil {
		log.Fatalln("One or more channels not created as expected")
	}

	feedInput(in)

	// Print output from the end of the chain
	for output := range out {
		fmt.Println(output)
	}
}

func buildChain(fn ...func(interface{}) interface{}) (in, out chan interface{}) {
	if in, out = nil, nil; len(fn) > 0 {
		in = make(chan interface{}, 1)
		outs := make([]chan interface{}, len(fn))
		for i := 0; i < len(fn); i++ {
			outs[i] = make(chan interface{}, 1)
			if i == 0 {
				go Node(fn[i], in, outs[i])
			} else {
				go Node(fn[i], outs[i-1], outs[i])
			}
		}
		out = outs[len(fn)-1]
	}
	return
}

func Process1(data interface{}) interface{} {
	switch data.(type) {
	case string:
		return fmt.Sprintf("( %v )", data)

	default:
		log.Printf("ERROR: Process1(): Illegal Data Type - %T", data)
		return nil
	}
}

func Process2(data interface{}) interface{} {
	switch data.(type) {
	case string:
		return fmt.Sprintf("{ %v }", data)

	default:
		log.Printf("ERROR: Process2(): Illegal Data Type - %T", data)
		return nil
	}
}

func Process3(data interface{}) interface{} {
	switch data.(type) {
	case string:
		return fmt.Sprintf("< %v >", data)

	default:
		log.Printf("ERROR: Process3(): Illegal Data Type - %T", data)
		return nil
	}
}
