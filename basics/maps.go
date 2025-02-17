package basics

import "fmt"

func MapsPrint() {
	// в качестве ключа может быть любой сравниваемый тип данных
	var m map[string]string = map[string]string{
		"name": "Tonya",
		"age":  "19",
	}
	fmt.Println(m) // map[age:19 name:Tonya]

	m1 := make(map[string]string, 5) // нужного размера
	fmt.Println(m1)

	mlen := len(m)
	fmt.Println(mlen)

	lastName := m["lastName"]
	fmt.Println(lastName) // возвращается значение по умолчание

	// проверка на существование ключа
	lastName, lastNameExists := m["lastName"]
	fmt.Println(lastName, lastNameExists)

	_, lastNameExists2 := m["lastName"]
	fmt.Println(lastNameExists2) // false

	// удаление ключа
	delete(m, "name")
	fmt.Println(m) // map[age:19]
}
