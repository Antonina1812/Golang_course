package basics

import "fmt"

func PointersPrint() {
	a := 2
	b := &a
	*b = 3 // a = 3
	c := &a

	d := new(int) // d = 0
	*d = 10
	*c = *d // a = 10 -> b = 10
	*d = 20

	c = d                    // c указывает туда же куда и d
	*c = 12                  // d = 12
	fmt.Println(&a, b, c, d) // 10 10 12 12
}
