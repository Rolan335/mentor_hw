package lrucache

import "slices"

//Переписать всё на односвязный список.
type LRUCache struct {
	cap      int
	data     map[string]any
	whenUsed []string
}

//Явно задать значения
func New(cap int) *LRUCache {
	return &LRUCache{cap: cap, data: make(map[string]any, cap), whenUsed: make([]string, 0, cap)}
}

//убрать false true
//returns false if key rewrited, true if added
func (l *LRUCache) Set(k string, v any) bool {
	//Check if key needs to be overwritten
	if _, ok := l.data[k]; ok {
		l.data[k] = v
		l.whenUsed = elemToStart(l.whenUsed, k)
		return false
	}

	l.data[k] = v
	if len(l.whenUsed) < l.cap {
		l.whenUsed = append([]string{k}, l.whenUsed...)
		return true
	}
	delete(l.data, l.whenUsed[len(l.whenUsed)-1])
	l.whenUsed = append([]string{k}, l.whenUsed[:len(l.whenUsed)-1]...)
	return true
}

func (l *LRUCache) Get(k string) (any, bool) {
	v, ok := l.data[k]
	if !ok {
		return nil, ok
	}
	l.whenUsed = elemToStart(l.whenUsed, k)
	return v, ok
}

//Не возвращать длину.
//returns map and count
func (l *LRUCache) GetAll() (map[string]any, int) {
	return l.data, len(l.data)
}

func (l *LRUCache) Delete(k string) bool {
	if _, ok := l.data[k]; !ok {
		return false
	}
	delete(l.data, k)
	//copy needed so that cap doesn't increase after deletion
	sliceTemp := append(elemToStart(l.whenUsed, k)[1:], "")
	copy(l.whenUsed, sliceTemp)
	return true
}

func (l *LRUCache) DeleteAll() {
	l.data = make(map[string]any, l.cap)
	l.whenUsed = make([]string, 0, l.cap)
}

func elemToStart(s []string, k string) []string {
	i := slices.Index(s, k)
	return append(append([]string{k}, s[:i]...), s[i+1:]...)
}
