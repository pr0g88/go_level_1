package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var nums = make([]int, 0, 10)
	fmt.Println("Введите первое число число")

	for scanner.Scan() {
		fmt.Println("Введите число или Ctrl+с для выхода")
		new, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, new)
	}

	for i := 1; i < len(nums); i++ {
		x := nums[i]
		j := i
		for ; j >= 1 && nums[j-1] > x; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = x
	}
	fmt.Println("Ваш результат: ", nums)
}
