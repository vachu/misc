package pipeline

import "sync"

//import "log"

func node(fn func(interface{}) (interface{}, error), in, out, diag chan interface{}, wg *sync.WaitGroup) {
	defer close(out)
	defer wg.Done()

	wg.Add(1)
	for {
		select {
		case inp, isOpen := <-in:
			if !isOpen {
				return
			} else {
				o, e := fn(inp)
				out <- o
				if e != nil && diag != nil {
					diag <- e
				}
			}
		default:
		}
	}
}

func BuildPipeline2(wantDiag bool, fns ...func(interface{}) (interface{}, error)) (chIn chan<- interface{}, chOut, chDiag <-chan interface{}) {
	chIn, chOut, chDiag = nil, nil, nil
	argCount := len(fns)
	if argCount == 0 {
		return
	}

	chInput := make(chan interface{})
	var chDiagOutput chan interface{}
	if wantDiag {
		chDiagOutput = make(chan interface{})
	}

	var wg sync.WaitGroup
	chOutputs := make([]chan interface{}, argCount)
	for i, fn := range fns {
		chOutputs[i] = make(chan interface{})
		if i == 0 {
			go node(fn, chInput, chOutputs[i], chDiagOutput, &wg)
		} else {
			go node(fn, chOutputs[i-1], chOutputs[i], chDiagOutput, &wg)
		}
	}
	if chDiagOutput != nil {
		go func() {
			wg.Wait()
			close(chDiagOutput)
		}()
	}

	chIn = chInput
	chDiag = chDiagOutput
	chOut = chOutputs[argCount-1]
	return
}
