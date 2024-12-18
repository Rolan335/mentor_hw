package wp

import (
	"context"
	"errors"
	"sync/atomic"
)

var ErrErrorsLimitExceeded error = errors.New("errors limit exceeded")

type Task func() error

type WorkerPool struct {
	workersDone int64
	workersMax  int64
	errCount    int64
	errMax      int64
	tasks       chan Task
}

func NewWorkerPool(workersMax int, errMax int) *WorkerPool {
	return &WorkerPool{
		workersDone: 0,
		workersMax:  int64(workersMax),
		errCount:    0,
		errMax:      int64(errMax),
		tasks:       make(chan Task, workersMax),
	}
}

// Add Tasks to queue and proceed.
func (w *WorkerPool) ProceedTasks(tasks []Task) error {
	//we need at least one worker
	if w.workersMax <= 0 {
		w.workersMax = 1
	}
	// if more workers than tasks, we need only workers == tasks number of workers
	if w.workersMax > int64(len(tasks)) {
		w.workersMax = int64(len(tasks))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//sending tasks to workers
	go w.send(ctx, tasks)

	for range w.workersMax {
		go w.work(ctx)
	}
	//checkForErrors waits until all work done or got maxErrors
	err := w.checkForErrors(cancel, len(tasks))
	if err != nil {
		return err
	}
	return nil
}

func (w *WorkerPool) send(ctx context.Context, tasks []Task) {
	defer close(w.tasks)
	for _, v := range tasks {
		select {
		case <-ctx.Done():
			return
		case w.tasks <- v:
		}
	}
}

func (w *WorkerPool) checkForErrors(cancel context.CancelFunc, taskCount int) error {
	defer cancel()
	watchErrors := true
	if w.errMax < 0 {
		watchErrors = false
	}
	//cancel all work and will return error
	if w.errMax == 0 {
		return ErrErrorsLimitExceeded
	}
	for {
		if atomic.LoadInt64(&w.errCount) >= w.errMax && watchErrors {
			return ErrErrorsLimitExceeded
		}
		if atomic.LoadInt64(&w.workersDone) >= int64(taskCount) {
			if atomic.LoadInt64(&w.errCount) >= w.errMax && watchErrors {
				return ErrErrorsLimitExceeded
			}
			return nil
		}
	}
}

func (w *WorkerPool) work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-w.tasks:
			if !ok {
				return
			}
			err := v()
			if err != nil {
				atomic.AddInt64(&w.errCount, 1)
			}
			atomic.AddInt64(&w.workersDone, 1)
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wp := NewWorkerPool(n, m)
	err := wp.ProceedTasks(tasks)
	if err != nil {
		return err
	}
	return nil
}
