/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1, 9, 2
Выход: 1, 2, 9
*/

package main

import "fmt"

func main() {
	var numberFirst int
	var numberSecond int
	var numberThird int
	var numberTmp int

	fmt.Print("Введите три числа, разделенных пробелом (прим. 1 9 2):")
	fmt.Scanln(&numberFirst, &numberSecond, &numberThird)

	if numberFirst > numberSecond {
		numberTmp = numberFirst
		numberFirst = numberSecond
		numberSecond = numberTmp
	}

	if numberFirst > numberThird {
		numberTmp = numberFirst
		numberFirst = numberThird
		numberThird = numberTmp

	}
	if numberSecond > numberThird {
		numberTmp = numberSecond
		numberSecond = numberThird
		numberThird = numberTmp
	}

	fmt.Printf("%d %d %d\n", numberFirst, numberSecond, numberThird)

}
