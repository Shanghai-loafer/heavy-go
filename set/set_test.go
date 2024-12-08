package set_test

import (
	"github.com/Shanghai-loafer/heavy-go/set"
	"reflect"
	"sort"
	"testing"
)

func TestSet_Add(t *testing.T) {
	items := set.Of[int]()
	items.Add(1)
	items.Add(2)
	items.Add(3)

	if items.Size() != 3 {
		t.Errorf("Expected items size to be 3, but got %d", items.Size())
	}

	if !items.Contains(1) || !items.Contains(2) || !items.Contains(3) {
		t.Error("Expected items to contain elements 1, 2, 3")
	}
}

func TestSet_Remove(t *testing.T) {
	items := set.Of[int]()
	items.Add(1)
	items.Add(2)
	items.Remove(1)

	if items.Contains(1) {
		t.Error("Expected items to not contain element 1 after removal")
	}

	if items.Size() != 1 {
		t.Errorf("Expected items size to be 1 after removal, but got %d", items.Size())
	}
}

func TestSet_Contains(t *testing.T) {
	items := set.Of[int]()
	items.Add(1)

	if !items.Contains(1) {
		t.Error("Expected items to contain element 1")
	}

	if items.Contains(2) {
		t.Error("Expected items to not contain element 2")
	}
}

func TestSet_Size(t *testing.T) {
	items := set.Of[int]()
	if items.Size() != 0 {
		t.Errorf("Expected empty items size to be 0, but got %d", items.Size())
	}

	items.Add(1)
	items.Add(2)
	items.Add(3)

	if items.Size() != 3 {
		t.Errorf("Expected items size to be 3, but got %d", items.Size())
	}
}

func TestSet_ToSlice(t *testing.T) {
	items := set.Of[int]()
	items.Add(1)
	items.Add(2)
	items.Add(3)

	slice := items.ToSlice()
	expected := []int{1, 2, 3}

	// ToSliceの順序は保証されないため、ソートして比較
	sort.Ints(slice)
	sort.Ints(expected)

	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("Expected slice %v, but got %v", expected, slice)
	}
}
