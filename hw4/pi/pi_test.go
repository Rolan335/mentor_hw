package pi

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPi(t *testing.T) {
	assert := assert.New(t)
	workers, iters := 12, 5_000_000
	piCalc := NewPiCalculator(workers, iters)
	go piCalc.Calc()
	time.Sleep(time.Second * 3)
	res := piCalc.End()
	assert.Equal(float32(res), float32(math.Pi))
}

func BenchmarkPi_12(b *testing.B) {
	workers, iters := 12, 1_000_000/12
	for range b.N {
		PiCalc := NewPiCalculator(workers, iters)
		go PiCalc.Calc()
		<-PiCalc.DoneCh
		_ = PiCalc.End()
	}
}

func BenchmarkPi_1(b *testing.B) {
	workers, iters := 1, 1_000_000/1
	for range b.N {
		PiCalc := NewPiCalculator(workers, iters)
		go PiCalc.Calc()
		<-PiCalc.DoneCh
		_ = PiCalc.End()
	}
}
