package basics

import "fmt"

func ConstantsPrint() {
	const pi = 3.1415

	const (
		str = "Hello"
		e   = 2.71828
	)

	const (
		zero  = iota // iota = 0
		_            // iota = 1 (пропускаем)
		_            // iota = 2 (пропускаем)
		three        // iota = 3
		four         // iota = 4
	)

	const (
		_         = iota             //0
		KB uint64 = 1 << (10 * iota) //1024
		MB                           //1048576
	)

	const (
		s1     = 1000 // нетипизированная
		s2 int = 1000 // типизированная
	)

	fmt.Println(pi, str, e)
	fmt.Println(zero, three, four)
	fmt.Println(KB, MB)
	fmt.Println(s1, s2)
}
