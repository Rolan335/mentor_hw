package lrucache_refactor

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

// v1 (slices)
// BenchmarkLRUCache_Set_Get-12              100000              2555 ns/op             153 B/op          3 allocs/op
// BenchmarkLRUCache_Set_Get-12              200000            415074 ns/op          843798 B/op          3 allocs/op
// v2 (linkedList)
// BenchmarkLRUCache_Set_Get-12              100000              2003 ns/op             139 B/op          2 allocs/op
// BenchmarkLRUCache_Set_Get-12              200000           1545357 ns/op              92 B/op          2 allocs/op
// v3 (MoveToStart instead of Remove and addToStart)
// BenchmarkLRUCache_Set_Get-12              200000            480551 ns/op              63 B/op          2 allocs/op
func BenchmarkLRUCache_Set_Get(b *testing.B) {
	cache := New(50000)
	for i := 1; i < b.N; i++ {
		b.StopTimer()
		key := strconv.Itoa(i)
		b.StartTimer()
		cache.Set(key, rand.Int())
		b.StopTimer()
		if i%5 == 0 {
			find := strconv.Itoa(i - 3)
			b.StartTimer()
			cache.Get(find)
		}
	}
}

func TestLRUCache_All(t *testing.T) {
	cache := New(3)
	cache.Set("a", "a")
	cache.Set("b", "b")
	cache.Set("c", "c")
	cache.Set("d", "d")
	want := map[string]any{"b": "b", "c": "c", "d": "d"}
	got := cache.GetAll()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("LRUCache() set d. want = %v got = %v", want, got)
	}
	cache.Get("c")
	cache.Set("e", "e")
	want = map[string]any{"d": "d", "c": "c", "e": "e"}
	got = cache.GetAll()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("LRUCache() get c set e. want = %v got = %v", want, got)
	}
	cache.Set("f", "f")
	want = map[string]any{"c": "c", "e": "e", "f": "f"}
	got = cache.GetAll()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("LRUCache() set f. want = %v got = %v", want, got)
	}
	cache.Delete("c")
	want = map[string]any{"e": "e", "f": "f"}
	got = cache.GetAll()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("LRUCache() delete c. want = %v got = %v", want, got)
	}
	cache.Set("g", "g")
	want = map[string]any{"e": "e", "f": "f", "g": "g"}
	got = cache.GetAll()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("LRUCache() set g. want = %v got = %v", want, got)
	}
	cache.DeleteAll()
	want = map[string]any{}
	got = cache.GetAll()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("LRUCache.deleteAll() want = %v got = %v", want, got)
	}
}

func TestLRUCache_Set(t *testing.T) {
	cache := New(4)
	cache.Set("a", "a")
	cache.Set("b", "b")
	cache.Set("c", "c")
	type args struct {
		k string
		v any
	}
	tests := []struct {
		cacheNew map[string]any
		name     string
		l        *LRUCache
		args     args
	}{
		{
			name:     "add when not full",
			l:        cache,
			args:     args{"d", "d"},
			cacheNew: map[string]any{"a": "a", "b": "b", "c": "c", "d": "d"},
		},
		{
			name:     "add when full. Delete last used elem",
			l:        cache,
			args:     args{"e", "e"},
			cacheNew: map[string]any{"b": "b", "c": "c", "d": "d", "e": "e"},
		},
		{
			name:     "rewrite existed value",
			l:        cache,
			args:     args{"b", "bb"},
			cacheNew: map[string]any{"c": "c", "d": "d", "e": "e", "b": "bb"},
		},
		{
			name:     "add when full. Delete last used elem after rewrite",
			l:        cache,
			args:     args{"f", "f"},
			cacheNew: map[string]any{"d": "d", "e": "e", "b": "bb", "f": "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache.Set(tt.args.k, tt.args.v)
			got := cache.GetAll()
			if !reflect.DeepEqual(got, tt.cacheNew) {
				t.Errorf("LRUCache.Set() cache = %v want = %v", got, tt.cacheNew)
			}
		})
	}
}

func TestLRUCache_Get(t *testing.T) {
	cache := New(3)
	cache.Set("1", 1)
	cache.Set("2", 2)
	cache.Set("3", 3)
	tests := []struct {
		name  string
		l     *LRUCache
		key   string
		want  any
		want1 bool
	}{
		{
			name:  "get not found",
			l:     cache,
			key:   "4",
			want:  nil,
			want1: false,
		},
		{
			name:  "get found",
			l:     cache,
			key:   "3",
			want:  3,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.l.Get(tt.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LRUCache.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LRUCache.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLRUCache_GetAll(t *testing.T) {
	cap := 5
	cache := New(cap)
	cache.Set("1", 5)
	cache.Set("2", 5)
	cache.Set("3", 5)
	cache.Set("4", 5)
	cache.Set("5", 5)
	cache.Set("6", 5)
	wantMap := map[string]any{"6": 5, "2": 5, "3": 5, "4": 5, "5": 5}
	gotMap := cache.GetAll()
	if !reflect.DeepEqual(wantMap, gotMap) {
		t.Errorf("LRUCache.GetAll() wantMap = %v gotMap = %v", wantMap, gotMap)
	}
}

func TestLRUCache_Delete(t *testing.T) {
	cache := New(3)
	cache.Set("1", 1)
	cache.Set("2", 2)
	cache.Set("3", 3)
	tests := []struct {
		want     bool
		name     string
		key      string
		cacheNew map[string]any
		l        *LRUCache
	}{
		{
			name:     "delete existing elem",
			l:        cache,
			key:      "2",
			want:     true,
			cacheNew: map[string]any{"1": 1, "3": 3},
		},
		{
			name:     "delete unexisting elem",
			l:        cache,
			key:      "555",
			want:     false,
			cacheNew: map[string]any{"1": 1, "3": 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDel := cache.Delete(tt.key)
			gotCache := cache.GetAll()
			if gotDel != tt.want {
				t.Errorf("LRUCache.Delete() return = %v want = %v", gotDel, tt.want)
			}
			if !reflect.DeepEqual(gotCache, tt.cacheNew) {
				t.Errorf("LRUCache.Delete() cache = %v want = %v", gotCache, tt.cacheNew)
			}
		})
	}
}

func TestLRUCache_DeleteAll(t *testing.T) {
	cap := 5
	cache := New(cap)
	cache.Set("1", 1)
	cache.Set("1", 2)
	cache.Set("1", 3)
	cache.Set("1", 4)
	cache.Set("1", 5)
	cache.DeleteAll()
	want := New(cap)
	if !reflect.DeepEqual(cache, want) {
		t.Errorf("LRUCache.DeleteAll() got = %v want = %v", cache, want)
	}
}
