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
	calculate()
}

func readInput() (string, float64, string, error) {
	var valueInput string
	var number float64
	var valueOutput string

	fmt.Println("Выберите что будем конвертировать: \nUSD, EUR, RUB")
	fmt.Scan(&valueInput)
	fmt.Println("Какое число?")
	fmt.Scan(&number)
	fmt.Println("Принял. Во что будем конвертировать?")
	fmt.Scan(&valueOutput)

	if valueInput == "" || valueOutput == "" {
		return "", 0, "", errors.New("Значение не может быть пустым")
	}

	if number <= 0 {
		return "", 0, "", errors.New("Значение не может быть меньше 0")
	}

	return valueInput, number, valueOutput, nil
}

func calculate() {
	var result float64
	what_convert, target, inConver, err := readInput()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	switch {
	case what_convert == "USD" && inConver == "EUR":
		// USD -> EUR: делим на курс EUR
		result = target / EUR
		fmt.Printf("%.2f USD = %.2f EUR\n", target, result)

	case what_convert == "USD" && inConver == "RUB":
		// USD -> RUB: делим на курс RUB
		result = target / RUB
		fmt.Printf("%.2f USD = %.2f RUB\n", target, result)

	case what_convert == "EUR" && inConver == "USD":
		// EUR -> USD: умножаем на курс EUR
		result = target * EUR
		fmt.Printf("%.2f EUR = %.2f USD\n", target, result)

	case what_convert == "RUB" && inConver == "USD":
		// RUB -> USD: умножаем на курс RUB
		result = target * RUB
		fmt.Printf("%.2f RUB = %.2f USD\n", target, result)

	case what_convert == "EUR" && inConver == "RUB":
		// EUR -> RUB: сначала в USD, потом в RUB
		usd := target * EUR
		result = usd / RUB
		fmt.Printf("%.2f EUR = %.2f RUB\n", target, result)

	case what_convert == "RUB" && inConver == "EUR":
		// RUB -> EUR: сначала в USD, потом в EUR
		usd := target * RUB
		result = usd / EUR
		fmt.Printf("%.2f RUB = %.2f EUR\n", target, result)

	default:
		fmt.Println("Неподдерживаемая конвертация валют")
	}
}
