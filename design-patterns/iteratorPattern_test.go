package designPatterns

import "testing"

func TestCreateIterator(t *testing.T) {
	inputSlice := []any{1, 51, 51, 135, 1}
	iterator := NewListIterator(inputSlice)

	if iterator == nil {
		t.Error("The creation failed")
	}
}

func TestIteratorNext(t *testing.T) {
	inputSlice := []any{1, 51, 51, 135, 1}
	iterator := NewListIterator(inputSlice)

	iterator.Next()

	if iterator.CurrentItem() != 51 {
		t.Error("The current item is not as expected")
	}
}

func TestIteratorNextAtArrayLimit(t *testing.T) {
	inputSlice := []any{1, 51, 51, 135, 111}
	iterator := NewListIterator(inputSlice)

	for {
		if !iterator.Next() {
			break
		}
	}
	if iterator.CurrentItem() != 111 {
		t.Error("Failed")
	}
}

func TestIteratorResetSetsBackAtZero(t *testing.T) {
	inputSlice := []any{1, 51, 51, 135, 111}
	iterator := NewListIterator(inputSlice)

	for {
		if !iterator.Next() {
			break
		}
	}
	if iterator.CurrentItem() != 111 {
		t.Error("The current item should be the last one")
	}

	iterator.Reset()

	if iterator.CurrentItem() != 1 {
		t.Error("The current item should be the first one")
	}
}
