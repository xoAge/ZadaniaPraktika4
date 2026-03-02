package main

import "fmt"

type doublyLinkedList struct {
	first *item
	last  *item
	size  int
}

type item struct {
	v    any
	next *item
	prev *item
}

func newDoublyLinkedList() *doublyLinkedList {
	return &doublyLinkedList{}
}

// add - добавление значения в связный список
func add(l *doublyLinkedList, v any) {
	newItem := &item{v: v}
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

// get - получение значения по индексу из связанного списка
func get(l *doublyLinkedList, idx int) any {
	if idx < 0 || idx >= l.size {
		return nil
	}
	current := l.first
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current.v
}

// remove - удаление значения по индексу из списка
func remove(l *doublyLinkedList, idx int) {
	if idx < 0 || idx >= l.size {
		return
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
}

// values - получение слайса значений из списка
func values(l *doublyLinkedList) []any {
	result := make([]any, 0, l.size)
	current := l.first
	for current != nil {
		result = append(result, current.v)
		current = current.next
	}
	return result
}

func main() {
	list := newDoublyLinkedList()
	add(list, 10)
	add(list, 20)
	add(list, 30)
	fmt.Println("Значения списка:", values(list))
	fmt.Println("Значение по индексу 1:", get(list, 1))
	remove(list, 1)
	fmt.Println("После удаления элемента по индексу 1:", values(list))
}
