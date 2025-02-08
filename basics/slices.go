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
}
