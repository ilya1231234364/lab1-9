package main

import "fmt"

func main() {
	//Инициализация переменных
	var a int16 = 2
	var b int16 = 3
	var c int64 = 10

	fmt.Println("a + b        = ", a+b)
	fmt.Println("a - b        = ", a-b)
	fmt.Println("a * b        = ", a*b)
	fmt.Println("int(a / b)   = ", int(a/b), "\n")

	c--
	fmt.Println("c--     = ", c)

	//Задание.
	//1. Выпонить выражение: fmt.Println("c--     = ", c--)
}
