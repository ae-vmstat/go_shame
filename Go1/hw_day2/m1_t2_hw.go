/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 100, выход: 001
*/

package main

import "fmt"

func main() {
	var number int
	fmt.Print("Введите трехзначное число: ")
	fmt.Scan(&number)

	numberReverse := (((number%10)*10)+((number%100)/10))*10 + (number / 100)

	fmt.Println(numberReverse)
}
