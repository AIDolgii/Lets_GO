package main

import "fmt"

func sum(array []int) int {
    total := 0
    for _, value := range array {
        total += value
    }
    return total
}

func main() {
	var array_size int
	fmt.Scanln(&array_size)
	var array = make([]int, array_size)
	for el := 0; el < array_size; el++ {
		fmt.Scan(&array[el])
	}

	fmt.Println(sum(array))
}