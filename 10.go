package main

import "fmt"

func getAllElements(list []int) ([]int, bool) {
	if len(list) > 0 {
		return list, true
	}
	return nil, false
}

func main() {
	myList := []int{1, 2, 3, 4, 5}
	allElements, found := getAllElements(myList)
	if found {
		fmt.Println("All elements:", allElements)
	} else {
		fmt.Println("List is empty")
	}
}
