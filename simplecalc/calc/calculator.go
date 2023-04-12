package calc

import "fmt"

type calculator struct{}

func NewCalculator() calculator {
	return calculator{}
}

func (c *calculator) Calculate(operation string, firstNumber float64, secondNumber float64) (float64, error) {

	var total float64

	switch operation {
	case "+":
		total = c.addition(firstNumber, secondNumber)
	case "-":
		total = c.subtraction(firstNumber, secondNumber)
	case "*":
		total = c.multiplication(firstNumber, secondNumber)
	case "/":
		if secondNumber == 0 {
			err := fmt.Errorf("недопустимое значение для вычитаемого: 0")
			return 0, err
		}
		total = c.division(firstNumber, secondNumber)
	default:
		err := fmt.Errorf("недопустимая операция")
		return 0, err
	}

	return total, nil
}

func (c *calculator) addition(firstNumber float64, secondNumber float64) float64 {
	return firstNumber + secondNumber
}

func (c *calculator) subtraction(firstNumber float64, secondNumber float64) float64 {
	return firstNumber - secondNumber
}

func (c *calculator) multiplication(firstNumber float64, secondNumber float64) float64 {
	return firstNumber * secondNumber
}

func (c *calculator) division(firstNumber float64, secondNumber float64) float64 {
	return firstNumber / secondNumber
}
