package set

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

func TestSet_New_Add_GetAllSorted(t *testing.T) {
	a := NewVals()
	a.Add("a", "b", "c", "d", "e")
	got := a.GetAll()
	slices.SortFunc(got, func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	want := []string{"a", "b", "c", "d", "e"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("a := New() a.Add() got = %v want = %v", got, want)
	}
}

func TestSet_Delete(t *testing.T) {
	a := NewVals("a", "b", "c", "d", "e")
	a.Delete("c", "d", "e")
	got := a.GetAll()
	slices.SortFunc(got, func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	want := []string{"a", "b"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("a.Delete() got = %v want = %v", got, want)
	}
}

func TestSet_IsPresent_NewVals(t *testing.T) {
	a := NewVals("1", "14", "4214")
	got := a.IsPresent("a")
	want := false
	if got != want {
		t.Errorf("a.IsPresent() got = %v want = %v", got, want)
	}
}

func TestSet_SubstractTwo(t *testing.T) {
	a := NewVals("1", "2", "3", "4", "5")
	b := NewVals("1", "2", "6")
	got, count := SubstractTwo(a, b)
	want := NewVals("3", "4", "5")
	countWant := 2
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SubstractTwo() got = %v want = %v", got, want)
	}
	if count != countWant {
		t.Errorf("SubstractTwo() count = %v countWant = %v", count, countWant)
	}
}

func TestSet_Union(t *testing.T) {
	a := NewVals("1", "2", "3", "4", "5")
	b := NewVals("1", "2", "6", "7", "5", "8", "9")
	got, count := Union(a, b)
	want := NewVals("1", "2", "3", "4", "5", "6", "7", "8", "9")
	countWant := 9
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Union() got = %v want = %v", got, want)
	}
	if count != countWant {
		t.Errorf("Union() count = %v countWant = %v", count, countWant)
	}
}

func TestSet_Intersect(t *testing.T) {
	a := NewVals("1", "2", "3", "4", "5")
	b := NewVals("1", "2", "6", "7", "5", "8", "4")
	got, count := Intersect(a, b)
	want := NewVals("1", "2", "5", "4")
	countWant := 4
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Intersect() got = %v want = %v", got, want)
	}
	if count != countWant {
		t.Errorf("Intersect() count = %v countWant = %v", count, countWant)
	}
}
