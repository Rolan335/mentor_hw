package semaphore

import (
	"context"
	"sync"
	"sync/atomic"
)

type Semaphore interface {
	Acquire(context.Context, int64) error
	TryAcquire(int64) bool
	Release(int64)
}

type SemaphoreMu struct {
	mu     *sync.Mutex
	locked bool
	len    int64
	cap    int64
}

func NewSemaphoreMu(want int64) *SemaphoreMu {
	return &SemaphoreMu{mu: &sync.Mutex{}, len: 0, cap: want, locked: false}
}

func (s *SemaphoreMu) Acquire(ctx context.Context, want int64) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if atomic.LoadInt64(&want)+atomic.LoadInt64(&s.len) > atomic.LoadInt64(&s.cap) {
			s.locked = true
			s.mu.Lock()
		}
		atomic.AddInt64(&s.len, want)
		return nil
	}
}

func (s *SemaphoreMu) TryAcquire(want int64) bool {
	if atomic.LoadInt64(&want)+atomic.LoadInt64(&s.len) > atomic.LoadInt64(&s.cap) {
		return false
	}
	atomic.AddInt64(&s.len, want)
	return true
}

func (s *SemaphoreMu) Release(want int64) {
	if s.locked {
		s.locked = false
		s.mu.Unlock()
	}
	atomic.AddInt64(&s.len, -want)
}

type SemaphoreChan struct {
	ch chan struct{}
}

func NewSemaphoreChan(cap int64) *SemaphoreChan {
	return &SemaphoreChan{ch: make(chan struct{}, cap)}
}

func (s *SemaphoreChan) Acquire(ctx context.Context, want int64) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		for range want {
			s.ch <- struct{}{}
		}
		return nil
	}
}
func (s *SemaphoreChan) TryAcquire(want int64) bool {
	if int(want)+len(s.ch) > cap(s.ch) {
		return false
	}
	for range want {
		select {
		case s.ch <- struct{}{}:
		default:
			return false
		}
	}
	return true
}

func (s *SemaphoreChan) Release(want int64) {
	for range want {
		select {
		case <-s.ch:
		default:
		}
	}
}

type SemaphoreCond struct {
	mu   sync.Mutex
	cond *sync.Cond
	len  int64
	cap  int64
}

func NewSemaphoreCond(cap int64) *SemaphoreCond {
	new := &SemaphoreCond{mu: sync.Mutex{}, len: 0, cap: cap}
	new.cond = sync.NewCond(&new.mu)
	return new
}

func (s *SemaphoreCond) Acquire(ctx context.Context, want int64) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		s.mu.Lock()
		defer s.mu.Unlock()
		for s.cap < s.len+want {
			s.cond.Wait()
		}
		s.len += want
		return nil
	}
}

func (s *SemaphoreCond) TryAcquire(want int64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.cap < s.len+want {
		return false
	}
	s.len += want
	return true
}

func (s *SemaphoreCond) Release(want int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.len -= want
	s.cond.Broadcast()
}
