/*
Задача №2
Вход:
Пользователь должен ввести правильный пароль, состоящий из:
цифр,
букв латинского алфавита(строчные и прописные) и 
специальных символов  special = "_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.
*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "_!@#$%^&"

Выход:
Написать, что ввели правильный пароль.

Пример: 
хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8 
*/

package main

import (
	"fmt"
	"regexp"
)

var password string

func main() {

fmt.Println("Необходимо придумать пароль от 8 до 15 символов, в пароле должны быть:") 
fmt.Println("необходимо использовать латинские буквы верхнего и нижнего регистра,")
fmt.Println("минимум одна цифра от 0 до 9, минимум один спец. символ: _!@#$%^&")

for i := 5; i > 0; i-- {

	fmt.Printf("Введите пароль: ")
	fmt.Scanln(&password)
	passwd := passwordCheck(password)
		if passwd == password {
			fmt.Println("хороший пароль: ", passwd)
			break
		} else {
			pea := i - 1

			fmt.Println("плохой пароль:", passwd)
			if pea == 0 {
				fmt.Println("не удалось создать пароль, попробуйте позднее")
				break
			}
			fmt.Println("осталось: ", pea, " попыток ввода пароля")
		}
	}
}

func passwordCheck(passwd string) (string){
	passLen := len(passwd)
	if passLen < 8 {
		passwd += " " + "|" + "меньше 8 символов" + "|"
	}
	if passLen > 15 {
		passwd += " " + "|" + "больше 15 символов" + "|"
	}
	passwdDigitsCheck := digitsCheck(passwd)
	 if passwd != passwdDigitsCheck {
			passwd += " " + "|" + passwdDigitsCheck + "|"
	 }
	 passwdLowercaseCheck := lowerCaseCheck(passwd)
	 if passwd != passwdLowercaseCheck {
			passwd += " " + "|" + passwdLowercaseCheck + "|" 
	 }
	passwdUppercasecaseCheck := upperCaseCheck(passwd)
	if passwd != passwdUppercasecaseCheck {
		passwd += " " + "|" + passwdUppercasecaseCheck + "|" 
	}
	passwdSpecialCheck := specialCheck(passwd)
	if passwd != passwdSpecialCheck {
		passwd += " " + "|" + passwdSpecialCheck + "|" 
	}
	return string(passwd)
}

func digitsCheck(passwd string) (string){
	match, _:=regexp.MatchString("[0-9]+", password) 
	if match ==  true {
		return string(passwd)
	} else {
		return string("не содержит цифр")
	}
}

func lowerCaseCheck(passwd string) (string){
	match, _:=regexp.MatchString("[a-z]+", password) 
	if match ==  true {
		return string(passwd)
	} else {
		return string("не содержит букв нижнего регистра")
	}
}

func upperCaseCheck(passwd string) (string){
	match, _:=regexp.MatchString("[A-Z]+", password) 
	if match ==  true {
		return string(passwd)
	} else {
		return string("не содержит букв верхнего регистра")
	}
}

func specialCheck(passwd string) (string){
	match, _:=regexp.MatchString("[_!@#$%^&]+", password) 
	if match ==  true {
		return string(passwd)
	} else {
		return string("не содержит спец. символов")
	}
}
