/*
Написать 3 функции.
Даны координаты трех точек(x1, y1, x2, y2, x3, y3), значения(целые) которых >= 0.
Первая функция проверяет, что можно построить треугольник по заданным точкам
Вторая функция вычисляет площадь треугольника.
Третья функция должна определить, является ли треугольник прямоугольным.
*/

package main

import (
"fmt"
"strconv"
"os"
//"math"
)

var coordinatPx1, coordinatPy1, coordinatPx2, coordinatPy2, coordinatPx3, coordinatPy3 string

//}

func main() {

	fmt.Print("Введите координаты точек целые числа >= 0. (x1, y1, x2, y2, x3, y3), разделенных пробелом (прим. 1 1 5 10 5 1):")
	fmt.Scanln(&coordinatPx1, &coordinatPy1, &coordinatPx2, &coordinatPy2, &coordinatPx3, &coordinatPy3)

	inputCheck(coordinatPx1, coordinatPy1, coordinatPx2, coordinatPy2, coordinatPx3, coordinatPy3)

	coordinatPx1 := InputCoordinatPointToFloat(coordinatPx1)
	coordinatPy1 := InputCoordinatPointToFloat(coordinatPy1)
	coordinatPx2 := InputCoordinatPointToFloat(coordinatPx2)
	coordinatPy2 := InputCoordinatPointToFloat(coordinatPy2)
	coordinatPx3 := InputCoordinatPointToFloat(coordinatPx3)
	coordinatPy3 := InputCoordinatPointToFloat(coordinatPy3)

	oneLine(coordinatPx1, coordinatPy1, coordinatPx2, coordinatPy2, coordinatPx3, coordinatPy3)
	areaOfTriangle(coordinatPx1, coordinatPy1, coordinatPx2, coordinatPy2, coordinatPx3, coordinatPy3)
	rightTriangle(coordinatPx1, coordinatPy1, coordinatPx2, coordinatPy2, coordinatPx3, coordinatPy3)
}

func inputCheck(args ... string){
	exitFlag := 0
	for _, v := range args {
		vInt, err := strconv.Atoi(v)
		if err != nil || vInt < 0 {
			fmt.Println("Неверная координата:", v)
			exitFlag++
			}
	}
	if exitFlag !=0 {
	fmt.Println("Ошибка ввода")
	os.Exit(1)
	}
}

func InputCoordinatPointToFloat(x string) (float64){
	vFloat, err := strconv.ParseFloat(x, 64)
	if err != nil {
		fmt.Println("Неверная координата:", x)
	}
	
	return vFloat
}	

func oneLine (coordinatPx1, coordinatPy1, coordinatPx2,
	coordinatPy2, coordinatPx3, coordinatPy3 float64) {

	if ((coordinatPx3 - coordinatPx1) / (coordinatPx2 - coordinatPx1)) ==
	((coordinatPy3 - coordinatPy1) / (coordinatPy2 - coordinatPy1)) {
		fmt.Println ("Не треугольник")
		os.Exit(2)
	}

}

func areaOfTriangle(coordinatPx1, coordinatPy1, coordinatPx2,
		coordinatPy2, coordinatPx3, coordinatPy3 float64){

	S := (((coordinatPx1-coordinatPx3)*(coordinatPy2-coordinatPy3))-
		((coordinatPx2-coordinatPx3)*(coordinatPy1-coordinatPy3)))
	if S < 0 {
		S = (S * (-1))*0.5
	} else {
		S = (S * 0.5)
	}
	fmt.Println("Площадь треугольника:", S)

}

func rightTriangle(coordinatPx1, coordinatPy1, coordinatPx2,
	coordinatPy2, coordinatPx3, coordinatPy3 float64){

	AB := int((coordinatPx2 - coordinatPx1) * (coordinatPx2 - coordinatPx1) + (coordinatPy2 - coordinatPy1) * (coordinatPy2 - coordinatPy1))
	BC := int((coordinatPx3 - coordinatPx2) * (coordinatPx3 - coordinatPx2) + (coordinatPy3 - coordinatPy2) * (coordinatPy3 - coordinatPy2))
	AC := int((coordinatPx3 - coordinatPx1) * (coordinatPx3 - coordinatPx1) + (coordinatPy3 - coordinatPy1) * (coordinatPy3 - coordinatPy1))

	if  AB + BC == AC ||
		AB + AC == BC ||
		BC + AC == AB {
		fmt.Println("Прямоуголный треугольник")
	}	
}
