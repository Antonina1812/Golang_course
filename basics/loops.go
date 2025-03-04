package basics

import "fmt"

func LoopsPrint() {
	for { // бесконечный цикл (while true)
		fmt.Println("iteration")
		break
	}

	i := true
	for i {
		fmt.Println("true")
		i = false
	}

	for i := 0; i < 2; i++ {
		fmt.Println("iter", i)
		if i == 1 {
			continue
		}
	}

	// slices
	s := []int{1, 2, 3}
	idx := 0

	for idx < len(s) { // как while
		fmt.Println(idx, s[idx])
		idx++
	}

	for idx := range s {
		fmt.Println(idx, s[idx])
	}

	for idx, val := range s {
		fmt.Println(idx, val)
	}

	// maps
	m := map[int]string{1: "tonya", 2: "cool"}

	for key := range m { // порядок ключей не определен, в разных запусках ключи в памяти могут располагаться по разному
		fmt.Println(key, m[key])
	}

	for key, val := range m {
		fmt.Println(key, val)
	}

	for _, val := range m {
		fmt.Println(val)
	}

	// strings
	str := "hello"             // строка - слайс байт
	for pos, ch := range str { // итерируемся по символам, а не по байтам!
		fmt.Println(ch, pos)
	}
}
