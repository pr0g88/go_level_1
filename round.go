package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64
	fmt.Println("Введите величину площади окружности")
	fmt.Scanf("%d\n", &s)
	
	b := s/3.14
	fmt.Println(b)
	
	res := math.Sqrt(b)
	fmt.Println(2*res)
	
}