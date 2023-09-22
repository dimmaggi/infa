package main

import (
	"container/list"
	"fmt"
)

func main() {

	myList := list.New()

	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	myList.PushBack(4)
	myList.PushBack(5)

	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	length := myList.Len()

	fmt.Println("Length of the list:", length)
}
