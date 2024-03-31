/*
	[31.03.2024]
	This seems to be working fine, I guess a more robust implementation would require some error checking.
	I will revise by implementing an Event Aggregator using channels
*/

package designPatterns

import (
	"fmt"
	"log"
)

func ObserverPatternDemo() {
	subject := NewSubject()
	observer1 := NewObserver()
	observer2 := NewObserver()

	subject.subjectState = "Some initial state"
	fmt.Printf("Obs1 state: %v\n", observer1.observerState)
	fmt.Printf("Obs1 state: %v\n", observer2.observerState)
	subject.Attach(&observer1)
	subject.Attach(&observer2)
	subject.subjectState = "Some changed state"
	subject.Notify()
	fmt.Printf("Obs1 state: %v\n", observer1.observerState)
	fmt.Printf("Obs1 state: %v\n", observer2.observerState)
}

type Subject interface {
	Attach(Observer)
	Dettach(Observer)
	Notify()
}

type ConcreteSubject struct {
	subjectState any
	observers    []Observer
}

func NewSubject() ConcreteSubject {
	return ConcreteSubject{}
}

func (c *ConcreteSubject) Notify() {
	for i := 0; i < len(c.observers); i++ {
		var observer Observer = c.observers[i]
		observer.Update(c.subjectState)
	}
}

func (c *ConcreteSubject) Attach(observer Observer) {
	c.observers = append(c.observers, observer)
}

func (c *ConcreteSubject) Dettach(observer Observer) {
	index := -1
	for i := 0; i < len(c.observers); i++ {
		if c.observers[i] == observer {
			index = i
		}
	}
	if index == -1 {
		log.Fatal("The observer was not found in the list")
	}
	RemoveAt(c.observers, index)
}

func RemoveAt[T any](slice []T, index int) []T {
	upperBound := len(slice) - 1
	for i := 0; i < upperBound; i++ {
		if i >= index {
			slice[i] = slice[i+1]
		}
	}
	slice = slice[:upperBound]
	return slice
}

type Observer interface {
	Update(any)
}

type ConcreteObserver struct {
	observerState any
}

func NewObserver() ConcreteObserver {
	return ConcreteObserver{}
}

func (obs *ConcreteObserver) Update(subjectState any) {
	obs.observerState = subjectState
}
