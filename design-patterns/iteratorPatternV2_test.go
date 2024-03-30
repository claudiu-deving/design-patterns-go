package designPatterns

import "testing"

func TestCreateIterator_V2(t *testing.T) {
	inputSlice := []any{1, 51, 51, 135, 1}
	iterator := GetEnumerator(inputSlice)

	if iterator == nil {
		t.Error("The creation failed")
	}
}

func TestIteratorNext_V2(t *testing.T) {
	inputSlice := []any{1, 51, 51, 135, 1}
	iterator := GetEnumerator(inputSlice)

	iterator.Next()

	if iterator.CurrentItem() != 51 {
		t.Error("The current item is not as expected")
	}
}

func TestIteratorNextAtArrayLimit_V2(t *testing.T) {
	inputSlice := []any{1, 51, 51, 135, 111}
	iterator := GetEnumerator(inputSlice)

	for {
		if !iterator.Next() {
			break
		}
	}
	if iterator.CurrentItem() != 111 {
		t.Error("Failed")
	}
}

func TestIteratorResetSetsBackAtZero_V2(t *testing.T) {
	inputSlice := []any{1, 51, 51, 135, 111}
	iterator := GetEnumerator(inputSlice)

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
