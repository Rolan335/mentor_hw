package lrucache_refactor

type LRUCache struct {
	cap      int
	data     map[string]any
	whenUsed *List
}

func New(cap int) *LRUCache {
	return &LRUCache{cap: cap, data: make(map[string]any, cap), whenUsed: NewList()}
}

func (l *LRUCache) Set(k string, v any) {
	if _, ok := l.data[k]; ok {
		l.data[k] = v
		l.whenUsed.Remove(k)
		l.whenUsed.AddToStart(k)
		return
	}

	l.data[k] = v
	if l.whenUsed.Len < l.cap {
		l.whenUsed.AddToStart(k)
		return
	}
	delete(l.data, l.whenUsed.GetEndVal())
	l.whenUsed.Remove(l.whenUsed.GetEndVal())
	l.whenUsed.AddToStart(k)
}

func (l *LRUCache) Get(k string) (any, bool) {
	v, ok := l.data[k]
	if !ok {
		return nil, ok
	}
	l.whenUsed.Remove(k)
	l.whenUsed.AddToStart(k)
	return v, ok
}

func (l *LRUCache) GetAll() map[string]any {
	return l.data
}

func (l *LRUCache) Delete(k string) bool {
	if _, ok := l.data[k]; !ok {
		return false
	}
	delete(l.data, k)
	l.whenUsed.Remove(k)
	return true
}

func (l *LRUCache) DeleteAll() {
	l.data = make(map[string]any, l.cap)
	l.whenUsed = NewList()
}
