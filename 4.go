package main

import "fmt"

func removeByIndex(list []int, index int) []int {
	if index >= 0 && index < len(list) {
		list = append(list[:index], list[index+1:]...)
	}
	return list
}

func main() {
	myList := []int{1, 2, 3, 4, 5}
	indexToRemove := 3
	myList = removeByIndex(myList, indexToRemove)
	fmt.Println("Updated list:", myList)
}
