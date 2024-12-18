package pipe

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Stage1A(in In) Out {
	out := make(Bi)
	go func() {
		defer func() {
			defer close(out)
			if r := recover(); r != nil {
				for val := range in {
					time.Sleep(time.Millisecond * 100)
					out <- val.(string) + "B"
				}
			}
		}()
		for val := range in {
			time.Sleep(time.Millisecond * 100)
			out <- val.(string) + "A"
		}
	}()

	return out
}

func Stage2B(in In) Out {
	out := make(Bi)
	go func() {
		defer func() {
			defer close(out)
			if r := recover(); r != nil {
				for val := range in {
					time.Sleep(time.Millisecond * 100)
					out <- val.(string) + "B"
				}
			}
		}()
		for val := range in {
			time.Sleep(time.Millisecond * 100)
			if val == "A" {
				panic("AAAAAAAAAAAA")
			}
			out <- val.(string) + "B"
		}
	}()

	return out
}

func Stage3C(in In) Out {
	out := make(Bi)
	go func() {
		defer func() {
			defer close(out)
			if r := recover(); r != nil {
				for val := range in {
					time.Sleep(time.Millisecond * 100)
					out <- val.(string) + "B"
				}
			}
		}()
		for val := range in {
			time.Sleep(time.Millisecond * 100)
			out <- val.(string) + "C"
		}
	}()

	return out
}

func TestPipe(t *testing.T) {
	a := assert.New(t)
	done := make(Bi)
	defer close(done)
	input := NewPipeInput("_", "_", "_", "_", "_")
	stages := []Stage{Stage1A, Stage2B, Stage3C, Stage3C}
	timer := time.Now()
	output := ExecutePipeline(input, done, stages...)
	for result := range output {
		a.Equal("_ABCC", result, "All stages completed")
	}
	a.GreaterOrEqual(850, int(time.Since(timer).Milliseconds()), "all input completed at stage time")
}

func TestPipeSigKill(t *testing.T) {
	a := assert.New(t)
	done := make(Bi)
	go func() {
		time.Sleep(time.Millisecond * 800)
		close(done)
	}()
	input := NewPipeInput("_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_")
	stages := []Stage{Stage1A, Stage2B, Stage3C, Stage3C, Stage1A, Stage2B, Stage3C, Stage3C, Stage1A, Stage2B, Stage3C, Stage3C, Stage1A, Stage2B, Stage3C, Stage3C}
	output := ExecutePipeline(input, done, stages...)
	for result := range output {
		a.NotEqual("_ABCCABCCABCCABCC", result, "not all stages completed")
	}
}
