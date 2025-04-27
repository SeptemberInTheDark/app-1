package main

import "fmt"

var USD float64
var EUR float64
var RUB float64

func main() {
    fmt.Println("___ Конвертер Валют ___")
}


func readInput() (float64, float64, float64) {
    fmt.Println("Введите сумму в USD:")
	fmt.Scan(&USD)
    fmt.Println("Введите сумму в EUR:")
    fmt.Scan(&EUR)
    fmt.Println("Введите сумму в RUB:")    
	fmt.Scan(&RUB)

    return USD, EUR, RUB
}


func calculate(summa float64, original_val float64, target_val float64) {
}