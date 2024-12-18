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
	doneCh   chan struct{}
	wg       *sync.WaitGroup
}

func NewPiCalculator(workersNum int, iters int) *PiCalculator {
	return &PiCalculator{
		maxIters: iters,
		workers:  workersNum,
		sumCh:    make(chan float64, workersNum),
		sigCh:    make(chan struct{}),
		doneCh:   make(chan struct{}),
		wg:       &sync.WaitGroup{},
	}
}

// function to start calculations
func (p *PiCalculator) Calc() {
	for i := range p.workers {
		p.wg.Add(1)
		go p.calcRow(i)
	}
}

// function to get results, should be called after sig of termination is passed or when <-p.DoneCh() closed
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
	for range p.maxIters / p.workers {
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

// returns a channel that will be closed that maxIters completed or when signal from outer sent
func (p *PiCalculator) DoneCh() <-chan struct{} {
	go func() {
		defer close(p.doneCh)
		for {
			select {
			case <-p.sigCh:
				return
			default:
				if len(p.sumCh) == p.workers {
					return
				}
			}
		}
	}()
	return p.doneCh
}
