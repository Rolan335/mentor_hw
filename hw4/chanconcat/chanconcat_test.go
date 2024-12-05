package chanconcat

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChanconcat(t *testing.T) {
	assert := assert.New(t)
	sig := func(duration time.Duration) <-chan interface{} {
		ch := make(chan interface{})
		go func() {
			defer close(ch)
			time.Sleep(duration)
		}()
		return ch
	}
	start := time.Now()
	_, ok := <-or(
		sig(time.Second*5),
		sig(time.Second*5),
		sig(time.Second*5),
		sig(time.Second*15),
		sig(time.Second*15),
		sig(time.Second*20),
	)
	goroutinesCount := runtime.NumGoroutine()
	sigCount := 6
	since := time.Since(start)
	assert.Equal(ok, false, " \"or\" channel close")
	assert.Equal(int(since.Seconds()), 5, "\"or\" channel closed after first sig")
	assert.Equal(runtime.NumGoroutine(), goroutinesCount-sigCount, "all goroutines made for \"or\" stopped successfuly")
}
