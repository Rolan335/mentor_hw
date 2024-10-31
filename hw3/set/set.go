package set

import (
	"slices"
	"strings"
)

// Stores unique passed values
type Set struct {
	set map[string]struct{}
}

func New() *Set {    
	return &Set{make(map[string]struct{})}        
}    

// Create set with values
func NewVals(input ...string) *Set {
	setInit := make(map[string]struct{}, len(input))
	for _, v := range input {
		setInit[v] = struct{}{}
	}
	return &Set{set: setInit}
}

// returns number of elements added
func (s *Set) Add(input ...string) int {
	count := 0
	for _, v := range input {
		if _, ok := s.set[v]; !ok {
			count++
		}
		s.set[v] = struct{}{}
	}
	return count
}

// return number of elements deleted
func (s *Set) Delete(input ...string) int {
	count := 0
	for _, v := range input {
		if _, ok := s.set[v]; ok {
			count++
			delete(s.set, v)
		}
	}
	return count
}

func (s *Set) IsPresent(k string) bool {
	_, ok := s.set[k]
	return ok
}

func (s *Set) GetAll() []string {
	slice := make([]string, 0, len(s.set))
	for k := range s.set {
		slice = append(slice, k)
	}
	return slice
}

func (s *Set) GetAllSorted() []string {
	slice := s.GetAll()
	slices.SortFunc(slice, func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	return slice
}

// Returns len of new set
func Union(a, b *Set) (*Set, int) {
	vals := append(a.GetAll(), b.GetAll()...)
	newSet := NewVals(vals...)
	return newSet, len(newSet.set)
}

// a - b
// Returns count substracted
func SubstractTwo(a, b *Set) (*Set, int) {
	deleted := a.Delete(b.GetAll()...)
	return a, deleted
}

// Returns count intersected
func Intersect(a, b *Set) (*Set, int) {
	setNew := New()
	count := 0
	for k := range a.set {
		if _, ok := b.set[k]; ok {
			count++
			setNew.Add(k)
		}
	}
	return setNew, count
}
