/*
Задача №4. Шахматная доска
Вход: размер шахматной доски, от 0 до 20
Выход: вывести на экран эту доску, заполняя поля нулями и единицами

Пример:
Вход: 5
Выход:
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
*/

package main

import (
"fmt"
"strconv"
"os"
"strings"
)

var intChessBoard int = 0
var inputChessBoard string

func main() {

	fmt.Print("Введите размер шахматной доски, от 0 до 20: ")
	fmt.Scanln(&inputChessBoard)

	intChessBoard := inputCheck(inputChessBoard)
	chessBoardF(intChessBoard)

}

func inputCheck(inputChessBoard string) (int){

	intChessBoard, err := strconv.Atoi(inputChessBoard)
		if err != nil || intChessBoard < 0 || intChessBoard > 200 {
			fmt.Println("Ошибка ввода: ", inputChessBoard)
			os.Exit(1)
		}
		if intChessBoard == 0 {
			fmt.Println("ноль клеток")
		}
	return intChessBoard

}

func chessBoardF(intChessBoard int) {
	
	var chessBoardArray [][]string
	
	for i := 0; i < intChessBoard; i++ {
		chessBoardArray = append(chessBoardArray, []string{})
		for j := 0; j < intChessBoard; j++ {
			if len(chessBoardArray) == 1 {
				if (j+1)%2 == 0 {
					chessBoardArray[i] = append(chessBoardArray[i], "1")	
				} else {
					chessBoardArray[i] = append(chessBoardArray[i], "0")
				}
			}
			if len(chessBoardArray) > 1 {
				if chessBoardArray[i-1][j] == "0" {
					chessBoardArray[i] = append(chessBoardArray[i], "1")	
				} else {
					chessBoardArray[i] = append(chessBoardArray[i], "0")
				}
			}
		}
	}

	for i := 0; i < len(chessBoardArray); i++ {
	chessBoardArrayStr := strings.Join(chessBoardArray[i], " ")
		fmt.Println(chessBoardArrayStr)
	}
}