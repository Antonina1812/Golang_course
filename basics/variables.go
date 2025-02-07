package basics

import "fmt"

func VariablesPrint() {
	var n1 int
	var n2 int = 10
	var n3 = 20
	n4 := 30
	fmt.Println(n1, n2, n3, n4) // 0 10 20 30

	var a, b int = 1, 2
	a, b = 3, 4
	a, c := 5, 6
	fmt.Println(a, b, c) //5 4 6

	var big int64 = 1<<32 - 1
	var unsigned uint = 100
	var bigUnsigned uint64 = 1<<64 - 1
	fmt.Println(big, unsigned, bigUnsigned)

	var pi float32 = 3.1415 //одинарная точность
	var pi2 float64 = 3.141 //двойная точность
	var e = 2.71828
	ch := 1.6666
	fmt.Println(pi, pi2, e, ch)

	var boo bool //по умолчанию - false
	var boo1 bool = true
	var boo2 = true
	boo3 := true
	fmt.Println(boo, boo1, boo2, boo3)

	var c0 complex128 = -1.1 + 7.1i
	var c1 complex64 = 8i
	c2 := -3i + 12
	fmt.Println(c0, c1, c2)
}
