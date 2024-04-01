package designPatterns

import (
	"fmt"
	"testing"
)

func TestRemoveAt(t *testing.T) {
	inputSlice := []any{'a', 'b', 'c', 'd', 'e', 'f'}
	inputSlice = RemoveAt(inputSlice, 1)
	if len(inputSlice) != 5 && inputSlice[1] != 'c' {
		t.Error("The removal failed")
	}
}

func TestObserversAreUpdatedWhenSubsribed(t *testing.T) {
	subject := NewSubject()
	observer1 := NewObserver()
	observer2 := NewObserver()

	subject.subjectState = "Some initial state"

	fmt.Printf("Obs1 state: %v\n", observer1.observerState)
	fmt.Printf("Obs1 state: %v\n", observer2.observerState)
	subject.Attach(&observer1)
	subject.Attach(&observer2)
	changedState := "Some changed state"
	subject.subjectState = changedState
	subject.Notify()

	if observer1.observerState != changedState &&
		observer2.observerState != changedState {
		t.Error("observers are not updated as expected")
	}
}
func TestObserversAreDettached(t *testing.T) {
	subject := NewSubject()
	observer1 := NewObserver()
	observer2 := NewObserver()

	subject.subjectState = "Some initial state"
	fmt.Printf("Obs1 state: %v\n", observer1.observerState)
	subject.Attach(&observer1)
	subject.Attach(&observer2)
	changedState := "Some changed state"
	subject.subjectState = changedState
	subject.Notify()
	if observer1.observerState != changedState &&
		observer2.observerState != changedState {
		t.Error("observers are not updated as expected")
	}
	err := subject.Dettach(&observer1)

	if err != nil {
		t.Error(err.Error())
	}

	if len(subject.observers) != 1 {
		t.Error("the observers are not dettaching as expected")
	}
	secondChangedState := "Another changed state"
	subject.subjectState = secondChangedState
	subject.Notify()

	if observer1.observerState != changedState &&
		observer2.observerState != secondChangedState {
		t.Error("the observers are still changed after detaching")
	}
}

func TestObserverIsNotUpdatedAfterUnsubcribing(t *testing.T) {
	subject := NewSubject()
	observer1 := NewObserver()

	subject.subjectState = "Some initial state"
	fmt.Printf("Obs1 state: %v\n", observer1.observerState)
	subject.Attach(&observer1)
	changedState := "Some changed state"
	subject.subjectState = changedState
	subject.Notify()
	if observer1.observerState != changedState {
		t.Error("observers are not updated as expected")
	}
	err := subject.Dettach(&observer1)

	if err != nil {
		t.Error(err.Error())
	}

	if len(subject.observers) != 0 {
		t.Error("the observers are not dettaching as expected")
	}
	secondChangedState := "Another changed state"
	subject.subjectState = secondChangedState
	subject.Notify()

	if observer1.observerState != changedState {
		t.Error("the observers are still changed after detaching")
	}
}
