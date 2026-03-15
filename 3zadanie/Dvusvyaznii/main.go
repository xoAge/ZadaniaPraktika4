package main

import (
	"Dvusvyazni/doublyList"
	"fmt"
)

func main() {
	list := doublyList.New[int]()

	// Добавляем элементы
	list.Add(10)
	list.Add(20)
	list.Add(30)

	// Выводим все значения
	fmt.Println("Значения списка:", list.Values())

	// Получаем значение по индексу
	if v, ok := list.Get(1); ok {
		fmt.Println("Значение по индексу 1:", v)
	} else {
		fmt.Println("Индекс вне диапазона")
	}

	// Удаляем элемент по индексу
	list.Remove(1)
	fmt.Println("После удаления элемента по индексу 1:", list.Values())
}
