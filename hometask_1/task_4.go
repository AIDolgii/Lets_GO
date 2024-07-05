package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Scanln(&num1, &num2, &num3)

	fmt.Println(max(num1, num2, num3))
}