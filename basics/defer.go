package basics

import "fmt"

func get() string {
	fmt.Println("func run")
	return "get smth"
}

func DeferPrint() {
	// множественные defer выполняются в порядке обратном их объявления
	// defer fmt.Println("after")
	// defer fmt.Println(get()) // аргументы отложных функций вычисляются при объявлении блока defer
	// fmt.Println("work")

	//func run
	//work
	//get smth
	//after

	defer fmt.Println("after")
	defer func() {
		fmt.Println(get())
	}()
	fmt.Println("work")

	//work
	//func run
	//get smth
	//after
}
