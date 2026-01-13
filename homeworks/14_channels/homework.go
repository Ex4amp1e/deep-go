package homework14

import (
	"errors"
	"sync"
)

// go test -v homework_test.go

var ErrPoolIsFull error = errors.New("pool is full")
var ErrPoolIsClosed error = errors.New("pool is closed")

type WorkerPool struct {
	tasks        chan func()
	wg           sync.WaitGroup
	shutdownOnce sync.Once
}

func NewWorkerPool(workersNumber int) *WorkerPool {
	wp := WorkerPool{
		tasks:        make(chan func(), workersNumber),
		wg:           sync.WaitGroup{},
		shutdownOnce: sync.Once{},
	}

	wp.wg.Add(workersNumber)
	for range workersNumber {
		go func() {
			defer wp.wg.Done()
			for task := range wp.tasks {
				task()
			}
		}()
	}

	return &wp
}

// Return an error if the pool is full
func (wp *WorkerPool) AddTask(task func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = ErrPoolIsClosed
		}

	}()

	select {
	case wp.tasks <- task:
		return nil
	default:
		return ErrPoolIsFull
	}
}

// Shutdown all workers and wait for all
// tasks in the pool to complete
func (wp *WorkerPool) Shutdown() {
	wp.shutdownOnce.Do(func() {
		close(wp.tasks)
		wp.wg.Wait()
	})

}
