//Nurassyl Hairulla SIS-2122
/*Write a program that takes a list of integers as
input and returns a new list containing the unique values from the original list.*/
package main

import "fmt"

func removeDuplicates(arr []int) []int {
	unique := make(map[int]bool)
	result := []int{}
	for _, value := range arr {
		if _, ok := unique[value]; !ok {
			result = append(result, value)
			unique[value] = true
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 1, 2, 5, 6, 3}
	uniqueArr := removeDuplicates(arr)
	fmt.Println(uniqueArr)
}

