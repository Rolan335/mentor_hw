package pi

import (
	"math"
	"sync"
)

type PiCalculator struct {
	maxIters int
	workers  int
	sumCh    chan float64
	sigCh    chan struct{}
	//DoneCh will be closed when iters maxItersCompleted
	DoneCh chan struct{}
	wg     *sync.WaitGroup
}

func NewPiCalculator(workersNum int, iters int) *PiCalculator {
	return &PiCalculator{
		maxIters: iters,
		workers:  workersNum,
		sumCh:    make(chan float64, workersNum),
		sigCh:    make(chan struct{}),
		DoneCh:   make(chan struct{}),
		wg:       &sync.WaitGroup{},
	}
}

func (p *PiCalculator) Calc() {
	go p.isDone()
	for i := range p.workers {
		p.wg.Add(1)
		go p.calcRow(i)
	}
}

func (p *PiCalculator) End() float64 {
	close(p.sigCh)
	p.wg.Wait()
	close(p.sumCh)
	var sum float64
	for v := range p.sumCh {
		sum += v
	}
	return sum * 4
}

func (p *PiCalculator) calcRow(i int) {
	defer p.wg.Done()
	var sum float64
	for range p.maxIters {
		select {
		case <-p.sigCh:
			p.sumCh <- sum
			return
		default:
			sum += math.Pow(float64(-1), float64(i)) / float64(2*i+1)
			i += p.workers
		}
	}
	p.sumCh <- sum
}

func (p *PiCalculator) isDone() {
	for {
		if len(p.sumCh) == p.workers {
			close(p.DoneCh)
			return
		}
	}
}

// func (p *PiCalculator) EndWhenReady() <-chan float64 {
// 	resCh := make(chan float64)
// 	go func() {
// 		p.wg.Wait()
// 		close(p.sumCh)
// 		close(p.sigCh)
// 		var sum float64
// 		for v := range p.sumCh {
// 			sum += v
// 		}
// 		resCh <- sum * 4
// 	}()
// 	return resCh
// }
