package main

import "fmt"

func getElementByIndex(list []int, index int) (int, bool) {
	if index >= 0 && index < len(list) {
		return list[index], true
	}
	return 0, false
}

func main() {
	myList := []int{1, 2, 3, 4, 5}
	index := 3
	element, found := getElementByIndex(myList, index)
	if found {
		fmt.Println("Element at index", index, "is", element)
	} else {
		fmt.Println("0 -->", "false")
	}
}
