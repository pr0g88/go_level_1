package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	var variant int
	fmt.Print("Введите 1 если хотите выполнить простую арифметическую операцию, введите 2 если ходите выполнить: возведение в степень, вычислить факториал: ")
	fmt.Scanln(&variant)

	if variant == 1 {
		one()
	} else if variant == 2 {
		two()
	}
}

func one() {
	var a, b, res float32
	var op string
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)
	fmt.Print("Введите арифметическую операцию (+, -, *, /): ")
	fmt.Scanln(&op)
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя!")
			os.Exit(1)
		} else {
			res = a / b
		}

	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}

func two() {
	var variant2 float32
	var n float64
	var res2 float64
	fmt.Print("Введите 1 если хотите возвести в степень, введите 2 если хотите вычислить факториал : ")
	fmt.Scanln(&variant2)

	if variant2 == 1 {
		step()
	} else if variant2 == 2 {
		fmt.Println("Введите число")
		fmt.Scanln(&n)
		factorial(n)
		res2 = factorial(n)
		fmt.Println(res2)
	}
}

func step() {
	var one1, two2, res1 float32
	fmt.Println("Введите число")
	fmt.Scanln(&one1)
	fmt.Println("Введите степень")
	fmt.Scanln(&two2)
	res1 = float32(math.Pow(float64(one1), float64(two2)))
	fmt.Printf("Результат выполнения операции: %.2f\n", res1)
}

func factorial(n float64) float64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
