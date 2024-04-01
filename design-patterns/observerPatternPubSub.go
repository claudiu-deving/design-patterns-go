/*
	[01.04.2024]
	I can't figure this one out yet
*/

package designPatterns

func ObserverPatternPubSubDemo() {

}

type PubSub interface {
	Subscribe(func(source any, item any))
	Publish(item any)
}

type PubSubC struct {
	pubsubs map[any][]any
}

func (p *PubSubC) Subscribe(item any) {

}
