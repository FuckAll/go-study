package concurrent

import (
	"fmt"
	"os"
	"os/signal"
	"testing"
	"time"
)

func TestRunner_Start(t *testing.T) {
	runner := New(98 * time.Second)
	func1 := func(id int) { fmt.Println(id) }
	func2 := func(id int) { fmt.Println(id); time.Sleep(15 * time.Second) }
	func3 := func(id int) { fmt.Println(id); time.Sleep(10 * time.Second) }
	tasks := []func(int){
		func1, func2, func3,
	}
	runner.Add(tasks...)
	err := runner.Start()
	if err != nil {
		t.Error(err)
	}
}

func TestRunner_Notify(t *testing.T) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	select {
	case <-ch:
		fmt.Println("ok")
	}

}
