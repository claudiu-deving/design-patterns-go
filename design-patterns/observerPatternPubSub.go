package designPatterns

func ObserverPatternPubSubDemo() {

	observer1 := "a"
	observer2 := "b"
	subject := "x"

	pubsubs := make(map[any][]any)
	observers := []any{
		&observer1,
		&observer2,
	}
	pubsubs[&subject] = observers
}
