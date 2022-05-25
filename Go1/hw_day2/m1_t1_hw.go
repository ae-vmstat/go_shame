/*
Задача №1
Вход:
    расстояние(50 - 10000 км),
    расход в литрах (5-25 литров) на 100 км и
    стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*/

package main

import "fmt"

func main() {
	var distance float64
	var consumptionLiters float64
	const gasolineCost float64 = 48

	fmt.Print("Введите расстояние(50 - 10000 км):")
	fmt.Scan(&distance)
	fmt.Print("Введите расход в литрах (5-25 литров):")
	fmt.Scan(&consumptionLiters)

	if 49 < distance && distance < 10001 && 4 < consumptionLiters && consumptionLiters < 26 {
		costOfTravel := (consumptionLiters * gasolineCost) * (distance / 100)
		//fmt.Printf("%d %d\n", distance, consumptionLiters)
		fmt.Println("стоимость поездки при цене бензина 48 руб.:", costOfTravel, "руб")
	} else {
		fmt.Println("Не поедем")
	}
}
