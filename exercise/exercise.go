package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	anonim()
	doSomething(func(number int) { fmt.Printf("number: %d", number) })

}

func mounths() {
	const (
		January = iota + 1
		February
		March
		April
		May
	)
	fmt.Println(January, February, March, April, May)
}

func anonim() {
	number := 10
	_ = number
	func() {
		number := 20
		fmt.Println(number)
	}()
}

func readconsole() {
	reader := bufio.NewReader(os.Stdin)

	str, _ := reader.ReadString(',') //ради упрощения примера ошибка игнорируется

	println(str)
}

func doSomething(anonymousFunc func(number int)) {
	anonymousFunc(100)
}
