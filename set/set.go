package set

type Set[T comparable] map[T]struct{}

// Of は空のSetを生成します
func Of[T comparable]() Set[T] {
	return make(Set[T])
}

// Add はSetに要素を追加します
func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

// Remove はSetから要素を削除します
func (s Set[T]) Remove(value T) {
	delete(s, value)
}

// Contains はSetに特定の要素が含まれているかを確認します
func (s Set[T]) Contains(value T) bool {
	_, exists := s[value]
	return exists
}

// Size はSetの要素数を返します
func (s Set[T]) Size() int {
	return len(s)
}

// ToSlice はSetの要素をスライスに変換して返します
func (s Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s))
	for key := range s {
		result = append(result, key)
	}
	return result
}
