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
	doneCh      chan struct{}
	tasks       chan Task
}

func NewWorkerPool(workersMax int, errMax int) *WorkerPool {
	return &WorkerPool{
		workersDone: 0,
		workersMax:  int64(workersMax),
		errCount:    0,
		errMax:      int64(errMax),
		doneCh:      make(chan struct{}),
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
	go w.kill(ctx, len(tasks))

	err := w.checkForErrors(cancel)
	if err != nil {
		return err
	}
	select {
	case <-w.doneCh:
		return nil
	
	}
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

func (w *WorkerPool) kill(ctx context.Context, taskCount int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if w.workersDone >= int64(taskCount) {
				close(w.doneCh)
				return
			}
		}
	}
}

func (w *WorkerPool) checkForErrors(cancel context.CancelFunc) error {
	//stop function so it won't be checking errors
	if w.errMax < 0 {
		return nil
	}
	//cancel all work and will return error
	if w.errMax == 0 {
		cancel()
		return ErrErrorsLimitExceeded
	}
	for {
		select {
		case <-w.doneCh:
			return nil
		default:
			if w.errCount >= w.errMax {
				cancel()
				return ErrErrorsLimitExceeded
			}
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

// package wp

// import (
// 	"context"
// 	"errors"
// 	"sync"
// )

// var ErrErrorsLimitExceeded error = errors.New("too many errors")

// type Task interface {
// 	Do() (any, error)
// }

// type Results []any

// type WorkerPool struct {
// 	wg           sync.WaitGroup
// 	workersMax   int
// 	errMax       int
// 	errCountChan chan struct{}
// 	task         chan Task
// 	result       chan any
// }

// func NewWorkerPool(workersMax int, errMax int) *WorkerPool {
// 	return &WorkerPool{
// 		wg:           sync.WaitGroup{},
// 		workersMax:   workersMax,
// 		errMax:       errMax,
// 		errCountChan: make(chan struct{}),
// 		task:         make(chan Task),
// 		result:       make(chan any),
// 	}
// }

// // Add Tasks to queue, proceed and return res.
// func (w *WorkerPool) ProceedTasks(task ...Task) (Results, error) {
// 	if w.workersMax > len(task) {
// 		w.workersMax = len(task)
// 	}

// 	// ctx, cancel := context.WithCancel(context.Background())
// 	// defer cancel()
// 	// go w.checkForErrors(cancel)

// 	w.wg.Add(len(task))
// 	go w.kill()

// 	for range w.workersMax {
// 		go w.work()
// 	}
// 	go func() {
// 		for _, v := range task {
// 			w.task <- v
// 		}
// 	}()
// 	res := w.getRes()
// 	return res, nil
// }

// func (w *WorkerPool) checkForErrors(cancel context.CancelFunc) {
// 	//stop function so it won't be checking errors
// 	if w.errMax <= 0 {
// 		return
// 	}
// 	counter := 0
// 	for range w.errCountChan {
// 		counter++
// 		if counter >= w.errMax {
// 			cancel()
// 			return
// 		}
// 	}
// }

// func (w *WorkerPool) work() {
// 	for v := range w.task {
// 		res, err := v.Do()
// 		if err != nil {
// 			w.errCountChan <- struct{}{}
// 		}
// 		w.result <- res
// 		w.wg.Done()
// 	}
// }

// func (w *WorkerPool) getRes() Results {
// 	res := make([]any, 0, w.workersMax)
// 	for v := range w.result {
// 		res = append(res, v)
// 	}
// 	return res
// }

// // todo ctx
// func (w *WorkerPool) kill() {
// 	w.wg.Wait()
// 	close(w.task)
// 	close(w.result)
// }
