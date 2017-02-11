package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	chain "../pipeline"
)

func feedInput(in chan<- interface{}) {
	defer close(in)

	in <- 123
	in <- "First"
	in <- "Second"
	in <- "Third"
}

func main() {
	inChannel, outChannel, diagChannel := chain.BuildPipeline2(true,
		Process1,
		Process2,
		Process3,
		toUpper,
	)
	if inChannel == nil || outChannel == nil {
		log.Fatalln("One or more channels not created as expected")
	}

	go feedInput(inChannel)
	if diagChannel != nil {
		go func() {
			for d := range diagChannel {
				fmt.Fprintln(os.Stderr, d)
			}
		}()
	}
	for o := range outChannel {
		fmt.Println(o)
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
