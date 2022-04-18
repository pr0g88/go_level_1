package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Введите трехзначное число")
	fmt.Scanf("%d\n", &a)

	sot := a / 100
	des := (a % 100) / 10
	edi := a % 10

	if sot == 0 {
		fmt.Println("В вашем числе сотен: 1", sot, ", десятков:", des, ", единиц:", edi)
	} else {
		fmt.Println("В вашем числе сотен: ", sot, ", десятков:", des, ", единиц:", edi)
	}

}
