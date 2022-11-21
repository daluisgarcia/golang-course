package main

import "fmt"

type Topic interface {
	register(observer Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(value string)
}

type Item struct {
	observers []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is now available\n", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

type EmailClient struct {
	id string
}

func (e *EmailClient) getId() string {
	return e.id
}

func (e *EmailClient) updateValue(value string) {
	fmt.Printf("Email sent to client '%s' about item '%s' is now available\n", e.id, value)
}

// func main() {
// 	// Having an item not available, when available must notify all the observers
// 	nvidiaItem := NewItem("RTX 3080")
// 	firstObserver := &EmailClient{id: "first"}
// 	secondObserver := &EmailClient{id: "second"}

// 	nvidiaItem.register(firstObserver)
// 	nvidiaItem.register(secondObserver)

// 	nvidiaItem.UpdateAvailable()
// }
