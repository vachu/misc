package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	chain "github.com/vachu/misc/pipeline"
)

func feedInput(in chan interface{}) {
	defer close(in)

	in <- 1234
	in <- "First"
	in <- "Second"
	in <- "Third"
}

func main() {
	in, out, diag := chain.BuildPipeline(
		Process1,
		Process2,
		Process3,
		toUpper,
	)
	if in == nil || out == nil {
		log.Fatalln("One or more channels not created as expected")
	}

	feedInput(in)

	// Print output from the end of the chain
	for o := range out {
		fmt.Println(o)
	}
	for d := range diag {
		fmt.Fprintln(os.Stderr, "DIAG:", d)
	}
}

func Process1(data interface{}) (interface{}, error) {
	switch data.(type) {
	case string:
		return fmt.Sprintf("( %v )", data), nil

	default:
		return nil, fmt.Errorf("ERROR: Process1(): Illegal Data Type - %T (data=%v)", data, data)
	}
}

func Process2(data interface{}) (interface{}, error) {
	switch data.(type) {
	case string:
		return fmt.Sprintf("{ %v }", data), nil

	default:
		return nil, fmt.Errorf("ERROR: Process2(): Illegal Data Type - %T (data=%v)", data, data)
	}
}

func Process3(data interface{}) (interface{}, error) {
	switch data.(type) {
	case string:
		return fmt.Sprintf("< %v >", data), nil

	default:
		return nil, fmt.Errorf("ERROR: Process3(): Illegal Data Type - %T (data=%v)", data, data)
	}
}

func toUpper(data interface{}) (interface{}, error) {
	switch data.(type) {
	case string:
		return strings.ToUpper(data.(string)), nil

	default:
		return nil, fmt.Errorf("ERROR: toUpper(): Illegal Data Type - %T (data=%v)", data, data)
	}
}
