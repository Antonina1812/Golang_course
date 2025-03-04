package basics

import "fmt"

func NamedReturn() (out int) { // именованный вывод
	out = 2
	return
}

func MultipleReturn(i int) (int, error) {
	if i > 2 {
		return 0, fmt.Errorf("error")
	} else {
		return i, nil // при возвращении нужно указывать все параметры через запятую!
	}
}

func MultipleNamedReturn(ok bool) (res int, err error) {
	if ok {
		err = fmt.Errorf("error")
		return res, err
	} else {
		res = 2
		return
	}
}

func FunctionsPrint() {
	NamedReturn()
	//MultipleReturn(5)
	//MultipleNamedReturn(true)
}
