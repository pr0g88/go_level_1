package main

import (
	"fmt"
	"math"
)

func main() {

	var s float64
	fmt.Println("Введите величину площади окружности")
	fmt.Scanf("%d\n", &s)

	// Вычисляем длину окружности
	c := 2 * math.Pi * math.Sqrt(s/math.Pi)
	fmt.Println("Длина окружности равна: ", c)

	// Вычисляем диаметр окружности
	d := c / math.Pi
	fmt.Println("Диаметр окружности равен: ", d)

}
