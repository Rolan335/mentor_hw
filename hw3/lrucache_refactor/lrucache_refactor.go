package lrucache_refactor

type LRUCache struct {
	cap      int
	data     map[string]any
	whenUsed *list
}

func New(cap int) *LRUCache {
	return &LRUCache{cap: cap, data: make(map[string]any, cap), whenUsed: newList()}
}

func (l *LRUCache) Set(k string, v any) {
	if _, ok := l.data[k]; ok {
		l.data[k] = v
		l.whenUsed.remove(k)
		l.whenUsed.addToStart(k)
		return
	}

	l.data[k] = v
	if l.whenUsed.len < l.cap {
		l.whenUsed.addToStart(k)
		return
	}
	delete(l.data, l.whenUsed.getEndVal())
	l.whenUsed.remove(l.whenUsed.getEndVal())
	l.whenUsed.addToStart(k)
}

func (l *LRUCache) Get(k string) (any, bool) {
	v, ok := l.data[k]
	if !ok {
		return nil, ok
	}
	l.whenUsed.moveToStart(k)
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
	l.whenUsed.remove(k)
	return true
}

func (l *LRUCache) DeleteAll() {
	l.data = make(map[string]any, l.cap)
	l.whenUsed = newList()
}
