package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func NewMan(n, ln, g string, a, cr int) Man {
	return Man{Name: n, LastName: ln, Age: a, Gender: g, Crimes: cr}
}

func main() {

	var maxCr int
	var foundCr bool
	var fondManName string

	people := make(map[string]Man)
	people["Василий Пупкин"] = NewMan("Василий", "Пупкин", "Муж.", 41, 4)
	people["Екатерина Пупкина"] = NewMan("Екатерина", "Пупкина", "Жен.", 38, 2)
	people["Алла Топорова"] = NewMan("Алла", "Топорова", "Жен.", 28, 1)
	people["Петр Кулешов"] = NewMan("Петр", "Кулешов", "Муж.", 25, 3)
	people["Семен Иванов"] = NewMan("Семен", "Иванов", "Муж.", 50, 7)
	people["Григорий Петров"] = NewMan("Григорий", "Петров", "Муж.", 36, 5)
	people["Прохор Сергеев"] = NewMan("Прохор", "Сергеев", "Муж.", 34, 11)
	people["Инна Топольская"] = NewMan("Инна", "Топольская", "Жен.", 65, 6)

	suspects := make([]string, 5)
	suspects = append(suspects, "Василий Пупкин", "Екатерина Пупкина", "Афанасий Кулешов", "Григорий Петров", "Прохор Сергеев")

	for _, crName := range suspects {
		crMan, ok := people[crName]
		if !ok {
			continue
		}

		foundCr = true
		if itIsMoreCrimes(maxCr, crMan.Crimes) {
			fondManName = crName
			maxCr = crMan.Crimes
		}
	}

	if !foundCr {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
	}

	if foundCr {
		fmt.Printf("Наибольшее количество преступлений совершил %s\n", fondManName)
	}

}

func itIsMoreCrimes(cr1, cr2 int) bool {

	return cr2 > cr1

}
