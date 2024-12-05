package semaphore

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSemaphoreMu(t *testing.T) {
	assert := assert.New(t)

	semCap := 5
	var sem Semaphore = NewSemaphoreMu(int64(semCap))
	assert.Implements((*Semaphore)(nil), sem)
	ctx := context.TODO()
	testSemaphore(ctx, assert, sem, semCap)
	testSemaphoreTryAcquire(assert, sem, semCap)
	ctx2, cancel := context.WithCancel(context.Background())
	cancel()
	testSemaphoreCtxTimeout(ctx2, assert, sem)
}

func TestSemaphoreChan(t *testing.T) {
	assert := assert.New(t)

	semCap := 5
	var sem Semaphore = NewSemaphoreChan(int64(semCap))
	assert.Implements((*Semaphore)(nil), sem)
	ctx := context.TODO()
	testSemaphore(ctx, assert, sem, semCap)
	testSemaphoreTryAcquire(assert, sem, semCap)
	ctx2, cancel := context.WithCancel(context.Background())
	cancel()
	testSemaphoreCtxTimeout(ctx2, assert, sem)
}

func TestSemaphoreCond(t *testing.T) {
	assert := assert.New(t)

	semCap := 5
	var sem Semaphore = NewSemaphoreCond(int64(semCap))
	assert.Implements((*Semaphore)(nil), sem)
	ctx := context.TODO()
	testSemaphore(ctx, assert, sem, semCap)
	testSemaphoreTryAcquire(assert, sem, semCap)
	ctx2, cancel := context.WithCancel(context.Background())
	cancel()
	testSemaphoreCtxTimeout(ctx2, assert, sem)
}

func testSemaphoreCtxTimeout(ctx context.Context, assert *assert.Assertions, sem Semaphore) {
	var (
		iters          = 10
		semCost  int64 = 1
		workDone       = 0
		timer          = time.Now()
	)
	for range iters {
		err := sem.Acquire(ctx, semCost)
		assert.Error(err, fmt.Sprintf("want - ctx.Err(), got - %v", err))
		if err == nil {
			go func() {
				defer sem.Release(semCost)
				time.Sleep(time.Second)
				workDone++
			}()
		}
	}
	since := time.Since(timer).Seconds()
	assert.Equal(0, int(since), "goroutines didn't run")
	assert.Equal(workDone, 0, "goroutines didn't run")
}

func testSemaphoreTryAcquire(assert *assert.Assertions, sem Semaphore, semCap int) {
	var (
		iters          = 10
		semCost  int64 = 1
		workDone       = 0
		timer          = time.Now()
	)
	for range iters {
		isAcquired := sem.TryAcquire(semCost)
		if isAcquired {
			go func() {
				defer sem.Release(semCost)
				time.Sleep(time.Second)
				workDone++
			}()
		}
	}
	for workDone != semCap {
	}
	since := time.Since(timer).Seconds()
	assert.Equal(int(since), 1, "goroutines are not blocked")
	assert.Equal(workDone, semCap, "semCap jobs done")
}

func testSemaphore(ctx context.Context, assert *assert.Assertions, sem Semaphore, semCap int) {
	var (
		iters          = 10
		semCost  int64 = 1
		workDone       = 0
		// timer          = time.Now()
	)
	for range iters {
		err := sem.Acquire(ctx, semCost)
		assert.NoError(err, fmt.Sprintf("want - nil, got - %v", err))
		if err == nil {
			go func() {
				defer sem.Release(semCost)
				time.Sleep(time.Second)
				workDone++
			}()
		}
	}
	for workDone != iters {
	}
	// since := time.Since(timer).Seconds()
	// assert.Equal(iters/semCap, int(since), "only semCap goroutines at a time")
	assert.Equal(workDone, iters, "all jobs done")
}
