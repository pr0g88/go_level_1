package main

import (
	"./config"
	"fmt"
)

func main() {
	fmt.Println("Считанная конфигурация из file.json: ")
	fmt.Println(config.Valid)
	fmt.Println(config.ReadJson)

}
