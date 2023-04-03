package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	str, _ := reader.ReadString(',') //ради упрощения примера ошибка игнорируется

	println(str)
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
