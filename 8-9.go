package main

import "fmt"

func GetByValue(list []int, value int) int {
	for _, item := range list {
		if item == value {
			return item
		}
	}
	return 0
}

func GetAllByValue(list []int, value int) []int {
	result := []int{}
	for _, item := range list {
		if item == value {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	myList := []int{1, 2, 3, 2, 4, 5, 2}
	valueToFind := 2

	firstOccurrence := GetByValue(myList, valueToFind)
	fmt.Println("First occurrence of", valueToFind, "is", firstOccurrence)

	allOccurrences := GetAllByValue(myList, valueToFind)
	fmt.Println("All occurrences of", valueToFind, "are", allOccurrences)
}
