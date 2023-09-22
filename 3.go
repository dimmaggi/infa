package main

import (
	"container/list"
	"fmt"
)

func main() {

	myList := list.New()

	myList.PushBack("Ананас")
	myList.PushBack("Зелень")
	myList.PushBack("Корень")
	myList.PushBack("МИ-8")
	myList.PushBack("Королева Елизавета")

	newElement := "кристалы"
	myList.PushBack(newElement)

	index := 0
	for e := myList.Front(); e != nil; e = e.Next() {
		if e.Value == newElement {
			break
		}
		index++
	}

	fmt.Println("Index of", newElement, "is", index)
}
