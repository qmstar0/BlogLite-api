package shutdown

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	tasks = make([]func() error, 0)
	mx    sync.RWMutex
)

func init() {
	go func() {
		downChan := make(chan os.Signal, 1)
		signal.Notify(downChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-downChan

		mx.RLock()
		defer mx.RUnlock()

		var errs = make(map[int]error)
		for i := range tasks {
			err := tasks[i]()
			if err != nil {
				errs[i] = err
			}
		}

		if len(errs) != 0 {
			fmt.Println("\033[1mAn error occurred while exiting and executing the exit task. (Add order index by task)\033[m")
			for i, err := range errs {
				fmt.Printf("\033[31m%d:%s\033[m\n", i, err)
			}
		} else {
			fmt.Println("\033[32mSafe Exit\033[m")
		}
		os.Exit(0)
	}()
}
func OnShutdown(fn func() error) {
	mx.Lock()
	defer mx.Unlock()
	tasks = append(tasks, fn)
}
