//runner pattern主要是使用channel监视程序的运行，给程序设置一个超时时间，有如下情况时终止程序。
//终止程序的情况有1，系统发打断消息；2，运行超时，3，运行完成，程序退出。
//这种pattern对于运行在后端的任务很有用。

package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner runs a set of tasks within a given timeout and can be
// shut down on an operating system interrupt.
type Runner struct {
	// interrupt channel reports a signal from the
	// operating system.
	interrupt chan os.Signal

	// complete channel reports that processing is done.
	complete chan error

	// timeout reports that time has run out.
	timeout <-chan time.Time //一个单向的接收channel。

	// tasks holds a set of functions that are executed
	// synchronously in index order.
	tasks []func(int)
}

// ErrTimeout is returned when a value is received on the timeout channel.
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt is returned when an event from the OS is received.
var ErrInterrupt = errors.New("received interrupt")

// New returns a new ready-to-use Runner.
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// 向runner添加任务组，一个任务就是一个参数为int的function。
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start runs all tasks and monitors channel events.
func (r *Runner) Start() error {
	//Notify函数执行后，如果发生了os.Interrupt，就会将他发送给channl r.interrupt
	signal.Notify(r.interrupt, os.Interrupt)

	// Run the different tasks on a different goroutine.
	go func() {
		r.complete <- r.run()
	}()

	select {
	// Signaled when processing is done.
	case err := <-r.complete:
		return err

	// Signaled when we run out of time.
	case <-r.timeout:
		return ErrTimeout
	}
}

// run executes each registered task.
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// Check for an interrupt signal from the OS.
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// Execute the registered task.
		task(id)
	}

	return nil
}

// gotInterrupt verifies if the interrupt signal has been issued.
func (r *Runner) gotInterrupt() bool {
	select {
	// Signaled when an interrupt event is sent.
	case <-r.interrupt:
		// Stop receiving any further signals.和前面的signal.Notify对应
		signal.Stop(r.interrupt)
		return true

	// Continue running as normal.
	default:
		return false
	}
}
