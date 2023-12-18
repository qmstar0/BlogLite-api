package shutdown

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	once sync.Once
	task = new(Tasks)
)

type Tasks struct {
	funcs []func()
	mutex sync.Mutex
}

func (t *Tasks) Add(fn func()) *Tasks {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.funcs = append(t.funcs, fn)
	return t
}

func wait() {
	once.Do(func() {
		downChan := make(chan os.Signal, 1)
		signal.Notify(downChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-downChan
		task.mutex.Lock()
		for _, fn := range task.funcs {
			fn()
		}
		task.mutex.Unlock()
		fmt.Println("Safe Exit")
		os.Exit(0)
	})
}

func WaitForShutdown() *Tasks {
	go wait()
	return task
}
