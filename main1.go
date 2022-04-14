package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Введите длину прямоугольника")
	fmt.Scanf("%d\n", &a)

	var b int
	fmt.Println("Введите ширину прямоугольника")
	fmt.Scanf("%d\n", &b)

	fmt.Println("Площадь прямоугольника =", a*b)
}