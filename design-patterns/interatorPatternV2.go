package designPatterns

/*[30.03.2024]*/
//We keep the same interface, I will try here to create a method for a generic slice that creates an Enumerator
//This works but of course, I learned about error handling so that should be implemented next
//Actually I had to create a new interface, to make it generic

type ListIteratorGeneric[T any] interface {
	Reset() T
	Next() bool
	CurrentItem() T
}

func GetEnumerator[T any](s []T) ListIteratorGeneric[T] {
	return &Enumerator[T]{inputSlice: s}
}

//I was missing the generic struct
type Enumerator[T any] struct {
	index      int
	inputSlice []T
}

func (concIt *Enumerator[T]) Reset() T {
	concIt.index = 0
	return concIt.CurrentItem()
}

func (concIt *Enumerator[T]) Next() bool {
	if len(concIt.inputSlice)-1 > int(concIt.index) {
		concIt.index++
		return true
	} else {
		return false
	}
}

func (concIt *Enumerator[T]) CurrentItem() T {
	return concIt.inputSlice[concIt.index]
}
