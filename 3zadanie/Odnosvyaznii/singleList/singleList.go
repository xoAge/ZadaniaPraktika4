package singleList

// элемент односвязного списка
type Item[T any] struct {
	V    T
	Next *Item[T]
}

// структура односвязного списка с дженериками
type SinglyLinkedList[T any] struct {
	First *Item[T]
	Last  *Item[T]
	Size  int
}

// конструктор списка
func New[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

// добавление элемента в конец списка
func (l *SinglyLinkedList[T]) Add(v T) {
	newItem := &Item[T]{V: v}
	if l.First == nil {
		l.First = newItem
		l.Last = newItem
	} else {
		l.Last.Next = newItem
		l.Last = newItem
	}
	l.Size++
}

// получение элемента по индексу
func (l *SinglyLinkedList[T]) Get(idx int) T {
	if idx < 0 || idx >= l.Size {
		panic("index out of bounds")
	}
	current := l.First
	for i := 0; i < idx; i++ {
		current = current.Next
	}
	return current.V
}

// удаление элемента по индексу
func (l *SinglyLinkedList[T]) Remove(idx int) {
	if idx < 0 || idx >= l.Size {
		panic("index out of bounds")
	}
	if idx == 0 {
		l.First = l.First.Next
		if l.First == nil {
			l.Last = nil
		}
	} else {
		prev := l.First
		for i := 0; i < idx-1; i++ {
			prev = prev.Next
		}
		toRemove := prev.Next
		prev.Next = toRemove.Next
		if toRemove == l.Last {
			l.Last = prev
		}
	}
	l.Size--
}

// получить слайс значений
func (l *SinglyLinkedList[T]) Values() []T {
	result := make([]T, 0, l.Size)
	current := l.First
	for current != nil {
		result = append(result, current.V)
		current = current.Next
	}
	return result
}
