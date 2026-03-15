package main

import (
	"fmt"
	"stack/Stack"
)

func main() {
	s := Stack.New[int](10)
	s.Push(5)
	s.Push(10)
	fmt.Println("Top element:", s.Peek())
	fmt.Println("Popped element:", s.Pop())
	fmt.Println("Top element after pop:", s.Peek())
}
