package main

import "fmt"

func clearList(list []int) []int {
	return nil
}

func main() {
	myList := []int{1, 2, 3, 4, 5}
	fmt.Println("Before clear:", myList)

	myList = clearList(myList)
	fmt.Println("After clear:", myList)
}
