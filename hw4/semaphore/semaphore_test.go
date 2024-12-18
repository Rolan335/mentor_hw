package semaphore

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type semTest struct {
	sem     Semaphore
	semCap  int32
	counter int32
}

func (s *semTest) work(ctx context.Context) {
	defer s.sem.Release(1)
	err := s.sem.Acquire(ctx, 1)
	if err != nil {
		atomic.AddInt32(&s.counter, 1)
	}
	time.Sleep(time.Second * 5)
}

func TestSemaphoreMu(t *testing.T) {
	assert := assert.New(t)

	semCap := 5
	var sem Semaphore = NewSemaphoreMu(int64(semCap))
	semTest := semTest{sem: sem, semCap: int32(semCap), counter: 0}
	assert.Implements((*Semaphore)(nil), sem)
	ctx := context.TODO()
	testSemaphore(ctx, assert, semTest)
	sem = NewSemaphoreMu(int64(semCap))
	testSemaphoreTryAcquire(assert, sem, semCap)
	sem = NewSemaphoreMu(int64(semCap))
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
	semTest := semTest{sem: sem, semCap: int32(semCap), counter: 0}
	testSemaphore(ctx, assert, semTest)
	sem = NewSemaphoreMu(int64(semCap))
	testSemaphoreTryAcquire(assert, sem, semCap)
	sem = NewSemaphoreMu(int64(semCap))
	ctx2, cancel := context.WithCancel(context.Background())
	cancel()
	testSemaphoreCtxTimeout(ctx2, assert, sem)
}

func TestSemaphoreCond(t *testing.T) {
	assert := assert.New(t)
	_ = 0

	semCap := 5
	var sem Semaphore = NewSemaphoreCond(int64(semCap))
	assert.Implements((*Semaphore)(nil), sem)
	semTest := semTest{sem: sem, semCap: int32(semCap), counter: 0}
	ctx := context.TODO()
	testSemaphore(ctx, assert, semTest)
	sem = NewSemaphoreMu(int64(semCap))
	testSemaphoreTryAcquire(assert, sem, semCap)
	sem = NewSemaphoreMu(int64(semCap))
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
		iters         = 10
		semCost int64 = 1
		wg            = &sync.WaitGroup{}
		counter int32 = 0
		timer         = time.Now()
	)
	for range iters {
		isAcquired := sem.TryAcquire(semCost)
		if isAcquired {
			wg.Add(1)
			go func() {
				defer sem.Release(semCost)
				atomic.AddInt32(&counter, 1)
				time.Sleep(time.Second)
				wg.Done()
			}()
		}
	}
	time.Sleep(time.Second)
	wg.Wait()
	since := time.Since(timer).Seconds()
	assert.Equal(int(since), 1, "goroutines are not blocked")
	assert.EqualValues(counter, semCap, "semCap jobs done")
}

func testSemaphore(ctx context.Context, assert *assert.Assertions, sem semTest) {
	var (
		iters         = 10
		semCost int64 = 1
	)
	for range iters {
		err := sem.sem.Acquire(ctx, semCost)
		assert.NoError(err, fmt.Sprintf("want - nil, got - %v", err))
		if err == nil {
			go func() {
				defer sem.sem.Release(semCost)
				atomic.AddInt32(&sem.counter, int32(semCost))
				time.Sleep(time.Second)
			}()
		}
		if atomic.LoadInt32(&sem.counter) == atomic.LoadInt32(&sem.semCap)+1 {
			isAcquired := sem.sem.TryAcquire(1)
			assert.EqualValues(false, isAcquired, "semaphore cannot acquire when full")
		}
	}
	for int(atomic.LoadInt32(&sem.counter)) != iters {
	}
	assert.EqualValues(sem.counter, iters, "all jobs done")
}
