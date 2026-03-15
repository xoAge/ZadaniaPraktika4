package doublyList

// двусвязный список
type DoublyLinkedList[T any] struct {
	first *item[T]
	last  *item[T]
	size  int
}

// создание нового списка
func New[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// добавление элемента в список
func (l *DoublyLinkedList[T]) Add(v T) {
	newItem := &item[T]{v: v}
	if l.size == 0 {
		l.first = newItem
		l.last = newItem
	} else {
		l.last.next = newItem
		newItem.prev = l.last
		l.last = newItem
	}
	l.size++
}

// получение элемента по индексу
func (l *DoublyLinkedList[T]) Get(idx int) (T, bool) {
	var zero T
	if idx < 0 || idx >= l.size {
		return zero, false
	}
	current := l.first
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current.v, true
}

// удаление элемента по индексу
func (l *DoublyLinkedList[T]) Remove(idx int) bool {
	if idx < 0 || idx >= l.size {
		return false
	}
	current := l.first
	for i := 0; i < idx; i++ {
		current = current.next
	}
	if current.prev != nil {
		current.prev.next = current.next
	} else {
		l.first = current.next
	}
	if current.next != nil {
		current.next.prev = current.prev
	} else {
		l.last = current.prev
	}
	l.size--
	return true
}

// возвращает слайс значений
func (l *DoublyLinkedList[T]) Values() []T {
	result := make([]T, 0, l.size)
	current := l.first
	for current != nil {
		result = append(result, current.v)
		current = current.next
	}
	return result
}
