package main

import (
	"Odnosvyazni/singleList"
	"fmt"
)

func main() {
	list := singleList.New[int]()

	list.Add(10)
	list.Add(20)
	list.Add(30)

	fmt.Println("Values:", list.Values())
	fmt.Println("Get index 1:", list.Get(1))
	list.Remove(1)
	fmt.Println("Values after remove:", list.Values())
}
