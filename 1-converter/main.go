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
		return "", 0, "", fmt.Errorf("ошибка чтения исходной валюты: %v", err)
	}
	fmt.Println("Какое число?")
	if _, err := fmt.Scan(&number); err != nil {
		return "", 0, "", fmt.Errorf("ошибка чтения числа: %v", err)
	}
	fmt.Println("Принял. Во что будем конвертировать?")
	if _, err := fmt.Scan(&valueOutput); err != nil {
		return "", 0, "", fmt.Errorf("ошибка чтения целевой валюты: %v", err)
	}

	if valueInput == "" || valueOutput == "" {
		return "", 0, "", errors.New("значение не может быть пустым")
	}

	if number <= 0 {
		return "", 0, "", errors.New("значение не может быть меньше или равно 0 или некорректный формат данных")
	}

	// Валидация валют
	if valueInput != "USD" && valueInput != "EUR" && valueInput != "RUB" {
		return "", 0, "", errors.New("неподдерживаемая исходная валюта. Используйте USD, EUR или RUB")
	}
	if valueOutput != "USD" && valueOutput != "EUR" && valueOutput != "RUB" {
		return "", 0, "", errors.New("неподдерживаемая целевая валюта. Используйте USD, EUR или RUB")
	}

	if valueInput == valueOutput {
		return "", 0, "", errors.New("исходная и целевая валюта не могут быть одинаковыми")
	}

	return valueInput, number, valueOutput, nil
}

func convertCurrency(fromCurrency string, toCurrency string, amount float64) float64 {
	// Сначала конвертируем в USD как базовую валюту
	var amountInUSD float64
	switch fromCurrency {
	case "USD":
		amountInUSD = amount
	case "EUR":
		amountInUSD = amount * EUR
	case "RUB":
		amountInUSD = amount * RUB
	}

	// Затем конвертируем из USD в целевую валюту
	switch toCurrency {
	case "USD":
		return amountInUSD
	case "EUR":
		return amountInUSD / EUR
	case "RUB":
		return amountInUSD / RUB
	}
	return 0
}

func calculate() {
	valueInput, number, valueOutput, err := readInput()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	result := convertCurrency(valueInput, valueOutput, number)
	fmt.Printf("%.2f %s = %.2f %s\n", number, valueInput, result, valueOutput)
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
