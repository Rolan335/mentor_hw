package pi

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPi(t *testing.T) {
	assert := assert.New(t)
	workers, iters := 12, 100_000_000
	piCalc := NewPiCalculator(workers, iters)
	go piCalc.Calc()
	<-piCalc.DoneCh()
	res := piCalc.End()
	assert.Equal(float32(res), float32(math.Pi))
}

func BenchmarkPi_12(b *testing.B) {
	workers, iters := 12, 500_000_000
	for range b.N {
		PiCalc := NewPiCalculator(workers, iters)
		go PiCalc.Calc()
		<-PiCalc.DoneCh()
		_ = PiCalc.End()
	}
}

func BenchmarkPi_1(b *testing.B) {
	workers, iters := 1, 500_000_000
	for range b.N {
		PiCalc := NewPiCalculator(workers, iters)
		go PiCalc.Calc()
		<-PiCalc.DoneCh()
		_ = PiCalc.End()
	}
}
