package concurrent

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

var ErrTimeout = errors.New("received timeout")

var ErrInterrupt = errors.New("received interrupt")

// Runner 在给定的一段时间内执行
// 操作系统发生中断的时候中断任务的执行
type Runner struct {
	// os signal
	interrupt chan os.Signal

	// task complete (nil or task error or interrupt error)
	complete chan error

	// do task timeout
	timeout <-chan time.Time

	// tasks
	tasks []func(int)
}

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d)}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	// notify os interrupt
	signal.Notify(r.interrupt, os.Interrupt)

	// run task
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		fmt.Println("ok")
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
