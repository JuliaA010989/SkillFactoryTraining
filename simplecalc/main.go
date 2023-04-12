package main

import (
	"SkillFactoryTraining/simplecalc/calc"
	"fmt"
	"log"
)

func main() {

	firstNumber, err := valNumber("first")
	if err != nil {
		log.Println("Ошибка при вводе первого значения: ", err)
		return
	}
	operation, err := valOperation()
	if err != nil {
		log.Println("Ошибка при вводе операции: ", err)
		return
	}
	secondNumber, err := valNumber("second")
	if err != nil {
		log.Println("Ошибка при вводе второго значения: ", err)
		return
	}

	calculator := calc.NewCalculator()

	total, err := calculator.Calculate(operation, firstNumber, secondNumber)
	if err != nil {
		log.Println("не удалось вычислить выражение, причина:", err)
		return
	}

	fmt.Printf("Результат выражения: %f", total)
	fmt.Println()

}

func valNumber(position string) (float64, error) {
	var number float64
	switch position {
	case "first":
		fmt.Print("Введите первое число: ")
	case "second":
		fmt.Print("Введите второе число: ")
	default:
		err := fmt.Errorf("недопустимая операция")
		return 0, err
	}
	_, err := fmt.Scanln(&number)

	return number, err
}

func valOperation() (string, error) {
	var operation string
	fmt.Print("Введите операцию +, -, * или /: ")
	_, err := fmt.Scanln(&operation)

	return operation, err
}
