package designPatterns

import "fmt"

/*
	The first version, I don't know yet how to make this actually generic, now, it only accepts []any which is dumb
*/
func IteratorDemoInitialVersion() {

	sliceOfInts := []any{1, 45, 6, 1, 4, 15, 141, 1, 151, 315, 1, 31, 1}
	iterator := NewListIterator(sliceOfInts)

	first := iterator.Reset()

	if first != 1 {
		fmt.Print("The first element should be 1")
	} else {
		fmt.Print(first)
	}

	iterator.Next()

	item := iterator.CurrentItem()

	if item != 45 {
		fmt.Print("The current element should be 45")
	} else {
		fmt.Print(item)
	}
}

type ListIterator interface {
	Reset() any
	Next() bool
	CurrentItem() any
}

type ConcreteIterator struct {
	inputSlice []any
	index      int64
}

func (concIt *ConcreteIterator) Reset() any {
	concIt.index = 0
	return concIt.CurrentItem()
}

func (concIt *ConcreteIterator) Next() bool {
	if len(concIt.inputSlice)-1 > int(concIt.index) {
		concIt.index++
		return true
	} else {
		return false
	}
}

func (concIt *ConcreteIterator) CurrentItem() any {
	return concIt.inputSlice[concIt.index]
}

func NewListIterator(inputSlice []any) ListIterator {
	return &ConcreteIterator{inputSlice: inputSlice}
}
