package main

import "fmt"

func main() {

	firstNumber, err := getNumber("first")
	if err != nil {
		panic(err)
	}
	operation, err := getOperation()
	if err != nil {
		panic(err)
	}
	secondNumber, err := getNumber("second")
	if err != nil {
		panic(err)
	}

	total, err := getTotal(operation, firstNumber, secondNumber)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Результат выражения: %f", total)
	fmt.Println()

}

func getNumber(position string) (float64, error) {
	var number float64
	switch position {
	case "first":
		fmt.Print("Введите первое число: ")
	case "second":
		fmt.Print("Введите второе число: ")
	default:
		panic("Некорректное значение позиции")
	}
	_, err := fmt.Scanln(&number)

	return number, err
}

func getOperation() (string, error) {
	var operation string
	fmt.Print("Введите операцию +, -, * или /: ")
	_, err := fmt.Scanln(&operation)

	return operation, err
}

func getTotal(operation string, firstNumber float64, secondNumber float64) (float64, error) {
	var total float64
	switch operation {
	case "+":
		total = addition(firstNumber, secondNumber)
	case "-":
		total = subtraction(firstNumber, secondNumber)
	case "*":
		total = multiplication(firstNumber, secondNumber)
	case "/":
		if secondNumber == 0 {
			err := fmt.Errorf("недопустимаое значение для вычитаемого: 0")
			return 0, err
		}
		total = division(firstNumber, secondNumber)
	default:
		err := fmt.Errorf("недопустимая операция")
		return 0, err
	}

	return total, nil
}

func addition(firstNumber float64, secondNumber float64) float64 {
	return firstNumber + secondNumber
}

func subtraction(firstNumber float64, secondNumber float64) float64 {
	return firstNumber - secondNumber
}

func multiplication(firstNumber float64, secondNumber float64) float64 {
	return firstNumber * secondNumber
}

func division(firstNumber float64, secondNumber float64) float64 {
	return firstNumber / secondNumber
}
