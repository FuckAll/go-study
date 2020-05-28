package observer

import "fmt"

type IMessage interface {
	getMessage() string
}

type DefaultMessage struct {
	message string
}

func (d *DefaultMessage) getMessage() string {
	return d.message
}

func (d *DefaultMessage) setMessage(str string) {
	d.message = str
}

type Subject interface {
	registerObserver(observer Observer) error
	removeObserver(observer Observer) error
	notifyObserver(message IMessage) error
}

type Observer interface {
	updateStatus(message IMessage)
}

type ObserverOne struct{}

func (o *ObserverOne) updateStatus(message IMessage) {
	fmt.Printf("ObserverOne get update status %+v\n", message)
}

type ObserverTwo struct {
}

func (o *ObserverTwo) updateStatus(message IMessage) {
	fmt.Printf("ObserverTwo get update status %+v\n", message)
}

type ConcreteSubject struct {
	observers []Observer
}

func (c *ConcreteSubject) registerObserver(observer Observer) error {
	c.observers = append(c.observers, observer)
	return nil
}

func (c *ConcreteSubject) removeObserver(observer Observer) error {
	for idx, o := range c.observers {
		if o == observer {
			c.observers = append(c.observers[:idx], c.observers[idx+1:]...)
			break
		}
	}
	return nil
}

func (c *ConcreteSubject) notifyObserver(message IMessage) error {
	for _, o := range c.observers {
		o.updateStatus(message)
	}
	return nil
}
