package main

import (
	"errors"
	"fmt"
)

var USD float64 = 1.0   // базовая валюта
var EUR float64 = 1.08  // 1 EUR = 1.08 USD
var RUB float64 = 0.011 // 1 RUB = 0.011 USD

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover:", r)
		}
	}()
	fmt.Println("___ Конвертер Валют ___")
	for {
		calculate()
		//
		isRepeatCalc := checkRepeatCalculation()
		if !isRepeatCalc {
			break
		}
	}

}

func readInput() (string, float64, string, error) {
	var valueInput string
	var number float64
	var valueOutput string

	fmt.Println("Выберите что будем конвертировать: \nUSD, EUR, RUB")
	if _, err := fmt.Scan(&valueInput); err != nil {
		panic("Ошибка чтения исходной валюты: " + err.Error())
	}
	fmt.Println("Какое число?")
	if _, err := fmt.Scan(&number); err != nil {
		panic("Ошибка чтения числа: " + err.Error())
	}
	fmt.Println("Принял. Во что будем конвертировать?")
	if _, err := fmt.Scan(&valueOutput); err != nil {
		panic("Ошибка чтения целевой валюты: " + err.Error())
	}

	if valueInput == "" || valueOutput == "" {
		return "", 0, "", errors.New("Значение не может быть пустым")
	}

	if number <= 0 {
		panic("Критическая ошибка: значение не может быть меньше или равно 0 или некорректный формат данных")
	}

	// Валидация валют
	if valueInput != "USD" && valueInput != "EUR" && valueInput != "RUB" {
		panic("Критическая ошибка: неподдерживаемая исходная валюта. Используйте USD, EUR или RUB")
	}
	if valueOutput != "USD" && valueOutput != "EUR" && valueOutput != "RUB" {
		panic("Критическая ошибка: неподдерживаемая целевая валюта. Используйте USD, EUR или RUB")
	}

	if valueInput == valueOutput {
		return "", 0, "", errors.New("Исходная и целевая валюта не могут быть одинаковыми")
	}

	return valueInput, number, valueOutput, nil
}

func calculate() {
	var result float64
	valueInput, number, valueOutput, err := readInput()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	switch {
	case valueInput == "USD" && valueOutput == "EUR":
		result = number / EUR
		fmt.Printf("%.2f USD = %.2f EUR\n", number, result)

	case valueInput == "USD" && valueOutput == "RUB":
		result = number / RUB
		fmt.Printf("%.2f USD = %.2f RUB\n", number, result)

	case valueInput == "EUR" && valueOutput == "USD":
		result = number * EUR
		fmt.Printf("%.2f EUR = %.2f USD\n", number, result)

	case valueInput == "RUB" && valueOutput == "USD":
		result = number * RUB
		fmt.Printf("%.2f RUB = %.2f USD\n", number, result)

	case valueInput == "EUR" && valueOutput == "RUB":
		usd := number * EUR
		result = usd / RUB
		fmt.Printf("%.2f EUR = %.2f RUB\n", number, result)

	case valueInput == "RUB" && valueOutput == "EUR":
		usd := number * RUB
		result = usd / EUR
		fmt.Printf("%.2f RUB = %.2f EUR\n", number, result)

	default:
		fmt.Println("Неподдерживаемая конвертация валют")
	}
}

func checkRepeatCalculation() bool {
	var userChoice string

	fmt.Println("Хотите сделать новый рассчет ? (Y/n)")
	fmt.Scan(&userChoice)
	if userChoice == "Y" || userChoice == "y" {
		return true
	}
	return false
}
