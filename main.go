package main

import "fmt"

func fibCache(n int, cache map[int]int) int {
	if val, ok := cache[n]; ok {
		return val
	}
	cache[n] = fibCache(n-1, cache) + fibCache(n-2, cache)
	return cache[n]
}

func main() {
	var n int
	fmt.Println("Введите значение")
	fmt.Scanln(&n)
	fmt.Println("Число Фибонначи: ", fibCache(n, map[int]int{0: 0, 1: 1}))

}
