package basics

import "fmt"

func NamedReturn() (out int) { // именованный вывод
	out = 2
	return
}

func MultipleReturn(i int) (int, error) {
	if i > 2 {
		return 0, fmt.Errorf("error")
	}
	return i, nil // при возвращении нужно указывать все параметры через запятую!
}

func MultipleNamedReturn(ok bool) (res int, err error) {
	if ok {
		err = fmt.Errorf("error")
		res = 14
		return res, err
	}
	res = 2
	return
}

func Sum(in ...int) (res int) {
	fmt.Println(in)
	for _, val := range in {
		res += val
	}
	return
}

type funcType func(string)

func FunctionsPrint() {

	fmt.Println(NamedReturn())
	fmt.Println(MultipleReturn(1))
	fmt.Println(MultipleNamedReturn(true))

	fmt.Println(Sum(1, 2, 3, 4))

	nums := []int{5, 6, 7, 8}
	fmt.Println(Sum(nums...))

	// анонимная функция
	func(i string) {
		fmt.Println("string", i)
	}("tonya")

	// присваивание анонимной функции в переменную
	printer := func(i string) {
		fmt.Println("string", i)
	}
	printer("cool")

	worker := func(callback funcType) {
		callback("as callback")
	}
	worker(printer) // callback = printer (значение, которое передаём в callback передаём в printer)

	// замыкание - это функция, которая ссылается на переменные вне своего тела функции
	prefixer := func(prefix string) funcType { // принимает строку, возвращает ф-ю, которая принимает строку
		return func(i string) { // принимает строку
			fmt.Println(prefix, i)
		}
	}
	successLogger := prefixer("SUCCESS") // prefix == "SUCCESS"
	successLogger("expected behaviour")  // i == "expected behaviour"
}
