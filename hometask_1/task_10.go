package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) GetSquare() float64 {
	return r.width * r.height
}

func main() {
	var w, h float64
	fmt.Scanln(&w, &h)
	rect := Rectangle{width: w, height: h}
	fmt.Println(rect.GetSquare())
}