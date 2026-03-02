package main

import "fmt"

type singlyLinkedList struct {
	first *item
	last  *item
	size  int
}

type item struct {
	v    any
	next *item
}

func newSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{}
}

// add - добавление значения в связный список в конец
func add(l *singlyLinkedList, v any) {
	newItem := &item{v: v}
	if l.first == nil {
		// список пустой
		l.first = newItem
		l.last = newItem
	} else {
		l.last.next = newItem
		l.last = newItem
	}
	l.size++
}

// get - получение значения по индексу из связного списка
func get(l *singlyLinkedList, idx int) any {
	if idx < 0 || idx >= l.size {
		panic("index out of bounds")
	}
	current := l.first
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current.v
}

// remove - удаление значения по индексу из списка
func remove(l *singlyLinkedList, idx int) {
	if idx < 0 || idx >= l.size {
		panic("index out of bounds")
	}
	if idx == 0 {
		// удаляем первый элемент
		l.first = l.first.next
		if l.first == nil {
			// список стал пустым
			l.last = nil
		}
	} else {
		prev := l.first
		for i := 0; i < idx-1; i++ {
			prev = prev.next
		}
		// prev.next - удаляемый элемент
		toRemove := prev.next
		prev.next = toRemove.next
		if toRemove == l.last {
			l.last = prev
		}
	}
	l.size--
}

// values - получение слайса значений из списка
func values(l *singlyLinkedList) []any {
	result := make([]any, 0, l.size)
	current := l.first
	for current != nil {
		result = append(result, current.v)
		current = current.next
	}
	return result
}

func main() {
	// Пример
	list := newSinglyLinkedList()
	add(list, 10)
	add(list, 20)
	add(list, 30)

	fmt.Println("Values:", values(list))
	fmt.Println("Get index 1:", get(list, 1))
	remove(list, 1)
	fmt.Println("Values after remove:", values(list))
}
