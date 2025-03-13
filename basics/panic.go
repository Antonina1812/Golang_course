package basics

import "fmt"

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic!", err)
		}
	}()
	fmt.Println("smth")
	panic("smth is bad") // рушится вся программа
}

func PanicPrint() {
	test()
	// smth
	// panic! smth is bad
}
