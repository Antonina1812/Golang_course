package basics

import "fmt"

type Persona struct {
	Id   int
	Name string
}

func (p Persona) EditName(name string) { // не изменяет оригинальную структуру (передается копия типа)
	p.Name = name
}

func (p *Persona) UpdateName(name string) { // изменяет оригинальную структуру (передается адрес)
	p.Name = name
}

type Profile struct {
	Id   int
	Name string
	Persona
}

func MethodsPrint() {
	var pers Persona
	//pers := new(Persona)

	//pers.EditName("tonya") // не имеет смысла
	pers.UpdateName("Tonya")
	(&pers).UpdateName("Tonechka")
	// строчки 22 и 23 "одинаковы"

	fmt.Println(pers.Name) // вывода при EditName не будет, а при UpdateName будет

	var prof Profile = Profile{
		Name: "tonya",
		Persona: Persona{
			Id:   2,
			Name: "Enot",
		},
	}

	prof.Persona.UpdateName("LoveGo")

	fmt.Println(prof.Persona.Name) // LoveGo
	fmt.Println(prof.Name)         //tonya

	prof.UpdateName("LoveTonya")

	fmt.Println(prof.Persona.Name) // LoveTonya
	fmt.Println(prof.Name)         //tonya
}
