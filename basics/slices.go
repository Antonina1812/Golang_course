package basics

import "fmt"

func SlicesPrint() {
	var sl []int              // len = 0, cap = 0
	sl1 := []int{}            // len = 0, cap = 0
	sl2 := []int{10}          // len = 1, cap = 1
	sl3 := make([]int, 0)     // len = 0, cap = 0
	sl4 := make([]int, 5)     // (слайс проинициализирован 5ю элементами) len = 5, cap = 5
	sl5 := make([]int, 5, 10) // len = 5, cap = 10
	fmt.Println(sl, sl1, sl2, sl3, sl4, sl5)

	var s []int
	s = append(s, 1, 2) // len = 2, cap = 2
	s = append(s, 3)    // len = 3, cap = 4
	fmt.Println(s)

	s1 := make([]int, 3) // [0 0 0]
	s = append(s, s1...) // len = 6, cap = 8
	fmt.Println(s)

	var slen, scap = len(s), cap(s)
	fmt.Println(slen, scap) // len = 6, cap = 8

	// срезы указывают на ту же область памяти
	buf := []int{1, 2, 3, 4, 5}
	b1 := buf[1:4] // [2 3 4]
	b2 := buf[:2]  // [1 2]
	b3 := buf[2:]  // [3 4 5]
	fmt.Println(b1, b2, b3)

	newBuf := buf[:]
	newBuf[0] = 10 // buf = [10 2 3 4 5] (та же область памяти)
	fmt.Println(buf, newBuf)

	newBuf = append(newBuf, 6) // другая область памяти
	newBuf[0] = 1
	fmt.Println(buf, newBuf)

	// копирование одного слайса в другой
	newBuf1 := make([]int, len(buf))
	copy(newBuf1, buf)
	fmt.Println(newBuf1)

	nums := []int{1, 2, 3, 4}
	copy(nums[1:3], []int{7, 8})
	fmt.Println(nums)
}
