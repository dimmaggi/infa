package main

import "fmt"

func RemoveByValue(list []int, value int) []int {
	for i := 0; i < len(list); i++ {
		if list[i] == value {
			list = append(list[:i], list[i+1:]...)
			break
		}
	}
	return list
}

func RemoveAllByValue(list []int, value int) []int {
	result := []int{}
	for _, item := range list {
		if item != value {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	myList := []int{1, 2, 3, 2, 4, 5, 2}
	valueToRemove := 2

	newList := RemoveByValue(myList, valueToRemove)
	fmt.Println("Список после RemoveByValue:", newList)

	newList = RemoveAllByValue(myList, valueToRemove)
	fmt.Println("Список после RemoveAllByValue:", newList)
}
