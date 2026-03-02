package main

import "fmt"

type stack struct {
	s    []any // слайс в котором хранятся значения в стеке
	head int   // индекс головы стека
}

// newStack создает новый стек заданного размера
func newStack(size int) *stack {
	return &stack{
		s:    make([]any, size),
		head: -1,
	}
}

// push - добавление в стек значения
func push(s *stack, v any) {
	if s.head+1 >= len(s.s) {
		panic("Stack overflow")
	}
	s.head++
	s.s[s.head] = v
}

// pop - получение значения из стека и его удаление из вершины
func pop(s *stack) any {
	if s.head == -1 {
		panic("Stack is empty")
	}
	val := s.s[s.head]
	s.s[s.head] = nil
	s.head--
	return val
}

// peek - просмотр значения на вершине стека
func peek(s *stack) any {
	if s.head == -1 {
		panic("Stack is empty")
	}
	return s.s[s.head]
}

func main() {
	// Пример использования
	s := newStack(10)
	push(s, 5)
	push(s, 10)
	fmt.Println("Top element:", peek(s))
	fmt.Println("Popped element:", pop(s))
	fmt.Println("Top element after pop:", peek(s))
}
