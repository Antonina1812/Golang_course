package basics

import "fmt"

func ArraysPrint() {
	//размер массива - часть типа данных
	//массив размерностью 2 и размерностью 3 - это два разных массива

	var a [3]int
	fmt.Println(a)         // [0 0 0]
	fmt.Printf("%v\n", a)  // [0 0 0]
	fmt.Printf("%#v\n", a) // [3]int{0, 0, 0}

	const size = 2 // для определения размера массива могут быть использованы только константы!
	var a1 [2 * size]bool
	fmt.Println(a1) // [false false false false]

	a2 := [...]int{1, 2, 3, 4} // троеточие автоматически определяет размер массива
	fmt.Println(a2)
}
