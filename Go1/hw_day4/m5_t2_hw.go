/* 
Задача №1
Написать функцию, которая расшифрует строку. 
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через вторую строку.

codeToString(code) -> "???????' 
*/

package main

import (
    "fmt"
	"strconv"
	"os"
)


func main() {

	var code string = "220411112603141304"

	fmt.Println(toCharStr(code))
}
		
	func toCharStr(i string) string {
		numChars := []rune(i)
		var code string
			for i := 0; i < len(numChars); i++ {
			chars := string(numChars[i])
			i++
			charsPlus := string(numChars[i])
			code += numLitToLetter(chars, charsPlus)
		}
		return string(code)
	}

func numLitToLetter (i, iPlus string) string {
	i = i + iPlus
	iInt, err := strconv.Atoi(i)
	if err != nil {
		os.Exit(1)
	}
	if iInt == 26 {
		return string(' ')
	} 
	return string('a' - 00 + iInt)

}
