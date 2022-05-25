package main

import (
	"./conf"
	"fmt"
)

func main() {
	fmt.Println(conf.Flag())
	fmt.Println("Результы валидации ссылки: ")
	conf.Valid()
}
