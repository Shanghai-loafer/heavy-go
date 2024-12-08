package collection

import (
	"reflect"
	"sort"
)

func Clone[T any](input interface{}) []T {
	v := reflect.ValueOf(input)

	// スライスや配列以外の場合はエラーを発生
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		panic("DeepClone: input must be a slice or array")
	}

	// スライスを受け取るための新しいスライスを作成
	clone := make([]T, v.Len())

	// 配列またはスライスの内容をコピー
	// 配列の場合もスライスに変換してcopy関数でコピー
	copy(clone, v.Slice(0, v.Len()).Interface().([]T))

	return clone
}

// Contains は、指定されたスライス内に item が存在するかを確認する関数
func Contains[T any](items []T, item T, equals func(T, T) bool) bool {
	for _, v := range items {
		if equals(v, item) {
			return true
		}
	}
	return false
}

// Filter は、スライスの要素を条件でフィルタリングする関数
func Filter[T any](items []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Find は、スライス内で条件にマッチする最初の要素を返す関数
// 見つからなかった場合は、要素のゼロ値とfalseを返します
func Find[T any](items []T, predicate func(T) bool) (T, bool) {
	var zero T
	for _, item := range items {
		if predicate(item) {
			return item, true
		}
	}
	return zero, false
}

// Map は、スライスの各要素に対して関数を適用し、変換したスライスを返す関数
func Map[T any, R any](items []T, fn func(T) R) []R {
	result := make([]R, len(items))
	for i, item := range items {
		result[i] = fn(item)
	}
	return result
}

// Reduce は、スライスの要素を累積的に処理し、1つの結果に集約する関数
func Reduce[T any, R any](items []T, initial R, fn func(R, T) R) R {
	result := initial
	for _, item := range items {
		result = fn(result, item)
	}
	return result
}

// Sort は、指定したソート条件でスライスをソートする関数
func Sort[T any](items []T, less func(a, b T) bool) []T {
	sort.Slice(items, func(i, j int) bool {
		return less(items[i], items[j])
	})
	return items
}
