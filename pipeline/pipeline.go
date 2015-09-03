package pipeline

import "sync"

func node(fn func(interface{}) (interface{}, error), in, out, diag chan interface{}) {
	defer close(out)
	defer close(diag)

	for {
		select {
		case inp, isOpen := <-in:
			if !isOpen {
				return
			} else if inp != nil {
				o, e := fn(inp)
				out <- o
				if e != nil {
					diag <- e
				}
			}
		default:
		}
	}
}

func BuildPipeline(fn ...func(interface{}) (interface{}, error)) (in, out, diag chan interface{}) {
	if in, out = nil, nil; len(fn) > 0 {
		in = make(chan interface{}, 1)
		outs := make([]chan interface{}, len(fn))
		diags := make([]chan interface{}, len(fn))
		for i := 0; i < len(fn); i++ {
			outs[i] = make(chan interface{}, 1)
			diags[i] = make(chan interface{}, 1)
			if i == 0 {
				go node(fn[i], in, outs[i], diags[i])
			} else {
				go node(fn[i], outs[i-1], outs[i], diags[i])
			}
		}
		out = outs[len(fn)-1]
		diag = mergeChannels(diags...)
	}
	return
}

func mergeChannels(outs ...chan interface{}) (merged chan interface{}) {
	merged = nil
	if len(outs) > 0 {
		for _, ch := range outs {
			if ch == nil {
				return
			}
		}

		var wg sync.WaitGroup
		wg.Add(len(outs))
		merged = make(chan interface{}, len(outs)*1)

		go func() {
			toBeDoneCtr := len(outs)
			for {
				for i := 0; i < len(outs); i++ {
					select {
					case data, isOpen := <-outs[i]:
						if isOpen {
							merged <- data
						} else {
							if toBeDoneCtr > 0 {
								wg.Done()
								toBeDoneCtr--
							}
						}
					default:
						if toBeDoneCtr == 0 {
							return
						}
					}
				}
			}
		}()
		go func() {
			wg.Wait()
			close(merged)
		}()
	}
	return
}
