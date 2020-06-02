package observer

import "fmt"

/*
	观察者模式：当对象间存在一对多关系时，则使用观察者模式（Observer Pattern）。比如，当一个对象被修改时，则会自动通知依赖它的对象。
	ConcreteSubject 是实例对象，能够注册观察者，当ConcreteSubject改变状态时，能够逐一的通知观察者。
*/
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
