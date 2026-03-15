package Stack

// структура стека с дженериками
type Stack[T any] struct {
	s    []T
	head int
}

// конструктор стека
func New[T any](size int) *Stack[T] {
	return &Stack[T]{
		s:    make([]T, size),
		head: -1,
	}
}

// добавить элемент в стек
func (st *Stack[T]) Push(v T) {
	if st.head+1 >= len(st.s) {
		panic("Stack overflow")
	}
	st.head++
	st.s[st.head] = v
}

// взять и удалить верхний элемент
func (st *Stack[T]) Pop() T {
	if st.head == -1 {
		panic("Stack is empty")
	}
	val := st.s[st.head]
	var zero T
	st.s[st.head] = zero // очистка
	st.head--
	return val
}

// посмотреть на верхний элемент
func (st *Stack[T]) Peek() T {
	if st.head == -1 {
		panic("Stack is empty")
	}
	return st.s[st.head]
}
