package pipe

import (
	"sync"
	"time"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) Out

var wg = &sync.WaitGroup{}

func NewPipeInput(data ...interface{}) Bi{
	pipeInput := make(Bi, len(data))
	for _, v := range data{
		pipeInput <- v
	}
	close(pipeInput)
	return pipeInput
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	pipeOut := make(Bi)
	for _, stage := range stages {
		wg.Add(1)
		in = stage(in)
		go func() {
			// Без time.sleep, на данных отрабатывают не все стейджи.
			time.Sleep(time.Millisecond)
			defer wg.Done()
			for v := range in {
				select {
				case <-done:
					return
				default:
					pipeOut <- v
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(pipeOut)
	}()
	return pipeOut
}
