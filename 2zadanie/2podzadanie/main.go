package main

import "fmt"

func main() {

	values := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	str := "XI"
	r := []rune(str)
	out := 0

	for i := 0; i < len(r); i++ {
		if i+1 < len(r) && values[r[i]] < values[r[i+1]] {
			out -= values[r[i]]
		} else {
			out += values[r[i]]
		}
	}

	fmt.Println(out)
}
