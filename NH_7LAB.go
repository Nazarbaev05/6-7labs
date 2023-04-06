//Nurassyl Hairulla SIS-2122
/*Write a program that takes a list of integers as
input and returns a new list containing the unique values from the original list.*/
package main

import "fmt"

func main() {
    var arr []int
    var n, x int

    fmt.Print("Enter the number of integers: ")
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        fmt.Printf("Enter integer %d: ", i+1)
        fmt.Scan(&x)
        arr = append(arr, x)
    }

    uniqueArr := []int{}
    for _, value := range arr {
        if !contains(uniqueArr, value) {
            uniqueArr = append(uniqueArr, value)
        }
    }

    fmt.Println("Original array:", arr)
    fmt.Println("Unique array:", uniqueArr)
}

func contains(arr []int, value int) bool {
    for _, element := range arr {
        if element == value {
            return true
        }
    }
    return false
}
