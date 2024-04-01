/*
	[31.03.2024]
	This seems to be working fine, I guess a more robust implementation would require some error checking.
	I will revise by implementing an Event Aggregator using channels

	[01.04.2024]
	I added error handling in some parts, while also creating some unit testing
	The RemoveAt func should be moved elsewhere
	The next step would be to introduce generics so that the subject becomes the object passed by the event (which will be wrapped in a channel)
*/

package designPatterns

import (
	"errors"
)

type Subject interface {
	Attach(Observer)
	Dettach(Observer) error
	Notify() []error
}

type ConcreteSubject struct {
	subjectState any
	observers    []Observer
}

func NewSubject() ConcreteSubject {
	return ConcreteSubject{}
}

func (c *ConcreteSubject) Notify() []error {
	errors := []error{}
	for i := 0; i < len(c.observers); i++ {
		var observer Observer = c.observers[i]
		err := observer.Update(c.subjectState)
		if err == nil {
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		return errors
	} else {
		return nil
	}

}

func (c *ConcreteSubject) Attach(observer Observer) {
	c.observers = append(c.observers, observer)
}

func (c *ConcreteSubject) Dettach(observer Observer) error {
	index := -1
	for i := 0; i < len(c.observers); i++ {
		if c.observers[i] == observer {
			index = i
		}
	}
	if index == -1 {
		return errors.New("the observer was not found in the list")
	}
	c.observers = RemoveAt(c.observers, index)

	return nil
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
	Update(any) error
}

type ConcreteObserver struct {
	observerState any
}

func NewObserver() ConcreteObserver {
	return ConcreteObserver{}
}

func (obs *ConcreteObserver) Update(subjectState any) error {
	if obs == nil {
		return errors.New("the observer is null")
	}
	obs.observerState = subjectState
	return nil
}
