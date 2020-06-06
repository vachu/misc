package hanoi

import (
	"container/list"
	"fmt"
)

// Tower ...
type Tower struct {
	*list.List

	name string
}

// New ...
func New(name string) *Tower {
	return &Tower{list.New(), name}
}

// Push ...
func (t *Tower) Push(value int) (e error) {
	if value == 0 {
		e = fmt.Errorf("attempted zero-value push")
	} else {
		t.PushBack(value)
	}
	return
}

// Pop ...
func (t *Tower) Pop() (int, error) {
	if lastElement := t.Back(); lastElement != nil {
		return t.Remove(lastElement).(int), nil
	}
	return 0, fmt.Errorf("cannot pop from empty stack")
}

// Peek ...
func (t *Tower) Peek() (int, error) {
	if lastElement := t.Back(); lastElement != nil {
		return lastElement.Value.(int), nil
	}
	return 0, fmt.Errorf("cannot peek into empty stack")
}

func (t *Tower) String() string {
	str := ""
	if len(t.name) > 0 {
		str = t.name + ": "
	}
	for le := t.Front(); le != nil; le = le.Next() {
		str += fmt.Sprintf("%v ", le.Value)
	}
	return str
}

func makeLegalMove(stepCount int, t1, t2 *Tower) {
	value1, _ := t1.Peek()
	value2, _ := t2.Peek()
	if value1 == value2 {
		panic(fmt.Sprintf("discs with size = %d found in %s and %s", value1, t1.name, t2.name))
	}

	switch {
	case value2 == 0:
		fallthrough
	case value1 != 0 && value1 < value2:
		fmt.Printf("%d. Move from %s to %s\n", stepCount, t1.name, t2.name)
		t1.Pop()
		t2.Push(value1)
	case value1 == 0:
		fallthrough
	case value2 != 0 && value2 < value1:
		fmt.Printf("%d. Move from %s to %s\n", stepCount, t2.name, t1.name)
		t2.Pop()
		t1.Push(value2)
	}
}

func recursiveSolution(discCount int, src, dst, aux *Tower, moveFn func(src, dst *Tower), dumpFn func()) {
	if discCount == 1 {
		moveFn(src, dst)
		dumpFn()
		return
	}
	recursiveSolution(discCount-1, src, aux, dst, moveFn, dumpFn)
	recursiveSolution(1, src, dst, aux, moveFn, dumpFn)
	recursiveSolution(discCount-1, aux, dst, src, moveFn, dumpFn)
}

// RecursiveSolution ...
func RecursiveSolution(src, dst, aux *Tower) {
	discCount := src.Len()
	if discCount <= 0 || dst.Len() != 0 || aux.Len() != 0 {
		panic("Illegal initial state of the Towers")
	}

	dumpFn := func() {
		dumpTowers("\t", src, aux, dst)
	}
	stepCount := 0
	moveFn := func(src, dst *Tower) {
		stepCount++
		srcLen := src.Len()
		makeLegalMove(stepCount, src, dst)
		if src.Len() > srcLen {
			panic("Unexpected move made")
		}
	}
	recursiveSolution(discCount, src, dst, aux, moveFn, dumpFn)
}

// IterativeSolution ...
func IterativeSolution(src, dst, aux *Tower) {
	if src == nil || aux == nil || dst == nil {
		return
	}
	discCount := src.Len()
	if discCount == 0 || dst.Len() != 0 || aux.Len() != 0 {
		panic("Illegal initial state of the Towers")
	}

	var t1, t2 *Tower
	if discCount%2 == 0 {
		t1, t2 = aux, dst
	} else {
		t1, t2 = dst, aux
	}
	moves := [][]*Tower{
		{src, t1},
		{src, t2},
		{aux, dst},
	}
	for i := 0; dst.Len() != discCount; i++ {
		m := i % 3
		makeLegalMove(i+1, moves[m][0], moves[m][1])
		dumpTowers("\t", src, aux, dst)
	}
}

func dumpTowers(padding string, towers ...*Tower) {
	for _, t := range towers {
		fmt.Printf("%s%s\n", padding, t)
	}
}

// Run ...
func Run() {
	towerA, towerB, towerC := New("Tower A"), New("Tower B"), New("Tower C")
	discCount := 5
	for i := 0; i < discCount; i++ {
		towerA.Push(discCount - i)
	}

	fmt.Println("Start ...")
	dumpTowers("", towerA, towerB, towerC)
	fmt.Println()

	fmt.Printf("Goal: to move all the discs from '%s' to '%s'\n", towerA, towerC)
	fmt.Println()

	// IterativeSolution(towerA, towerC, towerB)
	RecursiveSolution(towerA, towerC, towerB)

	fmt.Println("\n==== End ====")
}
