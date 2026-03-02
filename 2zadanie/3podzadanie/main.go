package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var rows, cols int

	// Считываем размеры массива
	fmt.Print("Введите количество строк: ")
	fmt.Scan(&rows)
	fmt.Print("Введите количество столбцов: ")

	fmt.Scan(&cols)
	rand.Seed(time.Now().UnixNano())

	// Создаем двумерный массив
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	used := make(map[int]bool)

	minVal, maxVal := 1, 100

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			var num int
			for {
				num = rand.Intn(maxVal-minVal+1) + minVal
				if !used[num] {
					used[num] = true
					break
				}
			}
			matrix[i][j] = num
		}
	}

	fmt.Println("Случайный двумерный массив с уникальными числами:")
	for _, row := range matrix {
		fmt.Println(row)
	}
}
