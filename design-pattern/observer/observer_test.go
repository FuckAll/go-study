package observer

import (
	"fmt"
	"testing"
)

func TestObserverPattern(t *testing.T) {
	subject := new(ConcreteSubject)
	one := new(ObserverOne)
	err := subject.registerObserver(one)
	if err != nil {
		t.Error(err)
	}
	two := new(ObserverTwo)
	err = subject.registerObserver(two)
	if err != nil {
		t.Error(err)
	}
	message := new(DefaultMessage)
	message.setMessage("hello world")
	err = subject.notifyObserver(message)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("--------------------------------")
	err = subject.removeObserver(one)
	if err != nil {
		t.Error(err)
	}
	err = subject.notifyObserver(message)
	if err != nil {
		t.Error(err)
	}
}
