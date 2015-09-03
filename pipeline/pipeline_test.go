package pipeline

import "testing"

func TestMergeChannels(t *testing.T) {
	testdata, chMerged := prepareMergeChannelsTestBed(t)
	i := 0
	for output := range chMerged {
		if output != testdata[i] {
			t.Errorf("expected '%s' - got '%s'", testdata[i], output)
		}
		i++
	}
}

func prepareMergeChannelsTestBed(t *testing.T) ([]string, chan interface {}) {
	testdata := []string{
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
	}

	chn := []chan interface{}{
		make(chan interface{}, 1),
		make(chan interface{}, 1),
		make(chan interface{}, 1),
	}
	chMerged := mergeChannels(chn...)

	go func() {
		for i, td := range testdata {
			i := i % len(chn)
			//t.Logf("DEBUG: i = %d\n", i)
			chn[i] <- td + " - CORRUPTED"
		}
		for _, ch := range chn {
			close(ch)
		}
	}()
	return testdata, chMerged
}
