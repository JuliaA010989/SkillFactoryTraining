package main

import (
	"SkillFactoryTraining/phones/electronic"
	"fmt"
)

func main() {

	var typePhone string

	applePhone13Pro := electronic.NewApplePhone("13 Pro")
	samsungGalPhone23 := electronic.NewAndroidPhone("Samsung", "Galaxy S23 Ultra")
	panasonicPhone1611 := electronic.NewRadioPhone("Panasonic", "KX-TG1611", 20)

	var applePhone electronic.Phone = &applePhone13Pro
	var samsungGalPhone electronic.Phone = &samsungGalPhone23
	var panasonicPhone electronic.Phone = &panasonicPhone1611

	var appleOS electronic.Smartphone = &applePhone13Pro
	var samsungOS electronic.Smartphone = &samsungGalPhone23

	var panasonicBut electronic.StationPhone = &panasonicPhone1611

	typePhone = printCharacteristics(applePhone)
	if typePhone == "smartphone" {
		printOS(appleOS)
	}
	typePhone = printCharacteristics(samsungGalPhone)
	if typePhone == "smartphone" {
		printOS(samsungOS)
	}
	typePhone = printCharacteristics(panasonicPhone)
	if typePhone == "station" {
		printButCont(panasonicBut)
	}

}

func printCharacteristics(p electronic.Phone) string {
	fmt.Printf("Бренд: %s, Модель: %s, Тип телефона: %s\n", p.Brand(), p.Model(), p.Type())
	return p.Type()
}

func printOS(sm electronic.Smartphone) {
	fmt.Println("Операционная система:", sm.OS())
}

func printButCont(st electronic.StationPhone) {
	fmt.Println("Количество кнопок:", st.ButtonsCount())
}

func printAddCharacteristic(p electronic.Phone, sm electronic.Smartphone, st electronic.StationPhone) {
	switch p.Type() {
	case "smartphone":
		fmt.Println("Операционная система:", sm.OS())
	case "station":
		fmt.Println("Количество кнопок:", st.ButtonsCount())
	default:
		fmt.Println("Нет дополнительных характеристик")
	}
}
