package set_test

import (
	"heavy-go/collection/set"
	"reflect"
	"sort"
	"testing"
)

func TestSet_Add(t *testing.T) {
	set := set.Of[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	if set.Size() != 3 {
		t.Errorf("Expected set size to be 3, but got %d", set.Size())
	}

	if !set.Contains(1) || !set.Contains(2) || !set.Contains(3) {
		t.Error("Expected set to contain elements 1, 2, 3")
	}
}

func TestSet_Remove(t *testing.T) {
	set := set.Of[int]()
	set.Add(1)
	set.Add(2)
	set.Remove(1)

	if set.Contains(1) {
		t.Error("Expected set to not contain element 1 after removal")
	}

	if set.Size() != 1 {
		t.Errorf("Expected set size to be 1 after removal, but got %d", set.Size())
	}
}

func TestSet_Contains(t *testing.T) {
	set := set.Of[int]()
	set.Add(1)

	if !set.Contains(1) {
		t.Error("Expected set to contain element 1")
	}

	if set.Contains(2) {
		t.Error("Expected set to not contain element 2")
	}
}

func TestSet_Size(t *testing.T) {
	set := set.Of[int]()
	if set.Size() != 0 {
		t.Errorf("Expected empty set size to be 0, but got %d", set.Size())
	}

	set.Add(1)
	set.Add(2)
	set.Add(3)

	if set.Size() != 3 {
		t.Errorf("Expected set size to be 3, but got %d", set.Size())
	}
}

func TestSet_ToSlice(t *testing.T) {
	set := set.Of[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	slice := set.ToSlice()
	expected := []int{1, 2, 3}

	// ToSliceの順序は保証されないため、ソートして比較
	sort.Ints(slice)
	sort.Ints(expected)

	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("Expected slice %v, but got %v", expected, slice)
	}
}
