/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100)
2) Количество (только числа)
3) ФИО покупателя (только буквы)
4) Контактный телефон (10 цифр)
5) Адрес(индекс(ровно 6 цифр), город, улица, дом, квартира)

Эти данные не могут быть пустыми.
Проверить правильность заполнения полей.

реализовать несколько методов у типа "Накладная"

createReader == NewReader
*/

package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
	"bufio"
	"strings"
	"regexp"
	"log"
)

type Winner struct {
	Name		string
	Telephone	string
	Address		 // это тип встроенной структуры
}

type Address struct {
	ZipCode 	string
	City		string
	Street		string
	House		string
	Apartment	string
}


var tmp string
var max int
var min int

func main() {

	min := randInt(1,100)
	max := randInt(1,100)

	if min > max {
		fmt.Println("Победитель не определен, попробуйте позднее.")
		os.Exit(1)
	}

	
	content, err := readLines("/home/egorovae/prize.txt")

	if err != nil {
		log.Fatal(err)
	}

	prizeArray := [100]string{}

	for i, line := range content {
		prizeArray[i] = line
	}

	WinnerArray := [2]string{} 
	
	for i := 0; i < 2; i++ {
		if i  == 0 {
			fmt.Printf("Введите ФИО победителя(только буквы): ")
			tmp = strings.TrimSuffix(input(), "\n")
			WinnerArray[0] = tmp
			if checkLetters(tmp) != true {
				for i := 1; i > 0;  i++ {
					fmt.Println("Неверный ввод: ", tmp)
					fmt.Printf("Введите ФИО победителя:")
					tmp = strings.TrimSuffix(input(), "\n")
					if checkLetters(tmp) == true {
						WinnerArray[0] = tmp
						break
					}
							
				}
			}
		} else {
			fmt.Printf("Введите телефон 10 цифр, не может начинаться с 0: ")
			tmp = strings.TrimSuffix(input(), "\n")
			WinnerArray[1] = tmp
			if checkTelephone(tmp) != true {
				for i := 1; i > 0;  i++ {
					fmt.Println("Неверный ввод: ", tmp)
					fmt.Printf("Введите телефон 10 цифр, не может начинаться с 0: ")
					tmp = strings.TrimSuffix(input(), "\n")
					if checkTelephone(tmp) == true {
						WinnerArray[1] = tmp
						break
					}
							
				}
			}
		}
	}

	WinnerAddressArray := [5]string{}

	for i := 0; i < 2; i++ {
		if i  == 0 {
			fmt.Printf("Введите индекс(6 цифр): ")
			tmp = strings.TrimSuffix(input(), "\n")
			WinnerAddressArray[0] = tmp
			if checkZipCode(tmp) != true {
				for i := 1; i > 0;  i++ {
					fmt.Println("Неверный ввод: ", tmp)
					fmt.Printf("Введите индекс(6 цифр): ")
					tmp = strings.TrimSuffix(input(), "\n")
					if checkZipCode(tmp) == true {
						WinnerAddressArray[0] = tmp
						break
					}
							
				}
			}
		} else {
			str := []string{"город", "улицу", "дом", "квартиру"}
			for i, s := range str{
				fmt.Printf("Введите %s: ",s)
				tmp = strings.TrimSuffix(input(), "\n")
				WinnerAddressArray[i+1] = tmp
			}
		}
	}

	winner := Winner{WinnerArray[0], WinnerArray[1], Address{WinnerAddressArray[0], 
		WinnerAddressArray[1], WinnerAddressArray[2], WinnerAddressArray[3], 
		WinnerAddressArray[4]}} 

	fmt.Println()
	fmt.Println("Победитель:", winner.getName(), "телефон:", winner.getTelephone())
	fmt.Println("Получает:")
	fmt.Println()
	
	for i := min; i < max; i++ {	
		fmt.Printf("%-76s %3d %3s\n", prizeArray[i], randInt(1,100), "шт.") 
	}

	fmt.Println()
	fmt.Println("Доставка будет осуществлена по адресу:")
	fmt.Println("почтовый индекс:", winner.getZipCode())
	fmt.Println("город:", winner.getCity())
	fmt.Println("улица:", winner.getStreet())
	fmt.Println("дом:", winner.getHouse())
	fmt.Println("квартира:", winner.getApartment())

}

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error input", err)
	}
	return string(str)
}

func randInt (min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max - min)
}
 	
func checkLetters (tmp string) bool{
	match, _ := regexp.MatchString(`^[\p{L} ]+$`, tmp)
	return(match)
}

func checkTelephone(tmp string) bool{
	match, _ := regexp.MatchString(`[1-9]\d{9}`, tmp)
	return(match)
}

func checkZipCode(tmp string) bool{
	match, _ := regexp.MatchString(`\d{6}`, tmp)
	return(match)
}


func (b Winner) getName() string {
	return b.Name
}

func (b Winner) getTelephone() string {
	return b.Telephone
}

func (b Address) getZipCode() string {
	return b.ZipCode
}

func (b Address) getCity() string {
	return b.City
}

func (b Address) getStreet() string {
	return b.Street
}

func (b Address) getHouse() string {
	return b.House
}

func (b Address) getApartment() string {
	return b.Apartment
}
