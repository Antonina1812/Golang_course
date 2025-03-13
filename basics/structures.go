package basics

import "fmt"

type Person struct {
	Id   int
	Name string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person
	Person  // встроенная структура (все поля Person являются частью структуры Account)
}

func SrtructuresPrint() {
	var acc Account = Account{ // полное объявление
		Id:   1,
		Name: "Tonya",
		Person: Person{
			Name: "Enot",
		},
	}
	fmt.Printf("%#v\n", acc)
	fmt.Println(acc)

	acc.Owner = Person{2, "Tonya"} // короткое объявление

	fmt.Printf("%#v\n", acc)
	fmt.Println(acc.Owner.Id)
	fmt.Println(acc.Name) // приоритет у более верхней (наружней) структуры, output: Tonya
}
