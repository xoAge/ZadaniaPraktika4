package doublyList

// item представляет элемент двусвязного списка
type item[T any] struct {
	v    T
	next *item[T]
	prev *item[T]
}
