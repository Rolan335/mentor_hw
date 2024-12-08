package wp

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded error = errors.New("too many errors")

type Task interface {
	Do() (any, error)
}

type Results []any

type WorkerPool struct {
	wg           sync.WaitGroup
	workersMax   int
	errMax       int
	errCountChan chan struct{}
	task         chan Task
	result       chan any
}

func NewWorkerPool(workersMax int, errMax int) *WorkerPool {
	return &WorkerPool{
		wg:           sync.WaitGroup{},
		workersMax:   workersMax,
		errMax:       errMax,
		errCountChan: make(chan struct{}, errMax),
		task:         make(chan Task),
		result:       make(chan any),
	}
}

// Add Tasks to queue, proceed and return res.
func (w *WorkerPool) ProceedTasks(task ...Task) (Results, error) {
	if w.workersMax > len(task) {
		w.workersMax = len(task)
	}
	w.wg.Add(len(task))
	go w.kill()
	for range w.workersMax {
		go w.work()
	}
	go func() {
		for _, v := range task {
			w.task <- v
		}
	}()
	res := w.getRes()
	return res, nil
}

func (w *WorkerPool) work() {
	for v := range w.task {
		res, err := v.Do()
		if err != nil {
			//todo
		}
		w.result <- res
		w.wg.Done()
	}
}

func (w *WorkerPool) getRes() Results {
	res := make([]any, 0, w.workersMax)
	for v := range w.result {
		res = append(res, v)
	}
	return res
}

func (w *WorkerPool) kill() {
	w.wg.Wait()
	close(w.task)
	close(w.result)
}
