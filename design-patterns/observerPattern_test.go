package designPatterns

import (
	"testing"
)

func TestRemoveAt(t *testing.T) {
	inputSlice := []any{'a', 'b', 'c', 'd', 'e', 'f'}
	inputSlice = RemoveAt(inputSlice, 1)
	if len(inputSlice) != 5 && inputSlice[1] != 'c' {
		t.Error("The removal failed")
	}
}

func TestObserverPatternDemo(t *testing.T) {
	ObserverPatternDemo()
}
