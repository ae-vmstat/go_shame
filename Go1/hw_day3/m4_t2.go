/*
Написать функцию.
Входные аргументы функции: количество бутылок от 0 до 200.
Функция должна вернуть количество и слово бутыл<?> с правильным окончанием.
Пример :
In: 5
Out: 5 бутылок
    
In: 1
Out: 1 бутылка
    
In: 22
Out: 22 бутылки
*/

package main

import (
"fmt"
"strconv"
"os"
)

var  bottles string

func main() {

	fmt.Print("Введите к-во бутылок от 0 до 200 (прим. 3):")
	fmt.Scanln(&bottles)
	
	bottles := inputCheck(bottles)

	if bottles%10 == 0 || bottles%10 > 4 || bottles > 10 && bottles <20 || bottles == 111 { 
		fmt.Println(bottles, "бутылок")		 
	}

	if bottles % 10 > 1 && bottles  < 5 || 
		bottles > 21 && (bottles % 10) > 1 && (bottles % 10) < 5 &&  (bottles % 10) != 1 {
		fmt.Println(bottles, "бутылоки")
	}

	if bottles % 10 == 1 && bottles != 111 && bottles != 11 {
		fmt.Println(bottles, "бутылока")
	}

}

func inputCheck(bottles string) (int){
	bInt, err := strconv.Atoi(bottles)
		if err != nil || bInt < 0 || bInt > 200 {
			fmt.Println("Ошибка ввода: ", bottles)
			os.Exit(1)
	}
	return bInt
}	