package collection_test

import (
	"heavy-go/collection"
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	if !collection.Contains(items, 3, func(a, b int) bool { return a == b }) {
		t.Error("Expected Contains to return true, but got false")
	}

	if collection.Contains(items, 6, func(a, b int) bool { return a == b }) {
		t.Error("Expected Contains to return false, but got true")
	}
}

func TestFilter(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4}
	result := collection.Filter(items, func(n int) bool { return n%2 == 0 })

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestFind(t *testing.T) {
	// テストデータ
	numbers := []int{1, 2, 3, 4, 5}

	// 条件に一致する要素が見つかる場合
	item, found := collection.Find(numbers, func(n int) bool { return n > 3 })
	if !found {
		t.Error("Expected to find an item, but none was found")
	}
	if item != 4 {
		t.Errorf("Expected item to be 4, but got %v", item)
	}

	// 条件に一致する要素が見つからない場合
	item, found = collection.Find(numbers, func(n int) bool { return n > 10 })
	if found {
		t.Errorf("Expected not to find any item, but found %v", item)
	}

	// 空のスライスで検索した場合
	var empty []int
	item, found = collection.Find(empty, func(n int) bool { return n > 0 })
	if found {
		t.Errorf("Expected not to find any item in an empty slice, but found %v", item)
	}
}

func TestMap(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4, 6, 8, 10}
	result := collection.Map(items, func(n int) int { return n * 2 })

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestReduce(t *testing.T) {
	// 合計のテスト
	numbers := []int{1, 2, 3, 4, 5}
	sum := collection.Reduce(numbers, 0, func(acc, n int) int {
		return acc + n
	})
	if sum != 15 {
		t.Errorf("Expected sum to be 15, but got %d", sum)
	}

	// 文字列の結合のテスト
	words := []string{"Go", "is", "awesome!"}
	sentence := collection.Reduce(words, "", func(acc, word string) string {
		if acc == "" {
			return word
		}
		return acc + " " + word
	})
	expectedSentence := "Go is awesome!"
	if sentence != expectedSentence {
		t.Errorf("Expected sentence to be '%s', but got '%s'", expectedSentence, sentence)
	}

	// 最大値のテスト
	numbers = []int{3, 1, 4, 1, 5, 9}
	max := collection.Reduce(numbers, numbers[0], func(acc, n int) int {
		if n > acc {
			return n
		}
		return acc
	})
	if max != 9 {
		t.Errorf("Expected max to be 9, but got %d", max)
	}

	// 要素の積のテスト
	product := collection.Reduce(numbers, 1, func(acc, n int) int {
		return acc * n
	})
	expectedProduct := 3 * 1 * 4 * 1 * 5 * 9 // = 540
	if product != expectedProduct {
		t.Errorf("Expected product to be %d, but got %d", expectedProduct, product)
	}
}

func TestSort(t *testing.T) {
	items := []int{5, 3, 4, 1, 2}
	expected := []int{1, 2, 3, 4, 5}
	collection.Sort(items, func(a, b int) bool { return a < b })

	if !reflect.DeepEqual(items, expected) {
		t.Errorf("Expected %v, but got %v", expected, items)
	}

	itemsDesc := []int{1, 2, 3, 4, 5}
	expectedDesc := []int{5, 4, 3, 2, 1}
	collection.Sort(itemsDesc, func(a, b int) bool { return a > b })

	if !reflect.DeepEqual(itemsDesc, expectedDesc) {
		t.Errorf("Expected %v, but got %v", expectedDesc, itemsDesc)
	}
}
