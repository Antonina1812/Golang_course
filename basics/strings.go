package basics

import (
	"fmt"
	"unicode/utf8"
)

func StringsPrint() {
	var s string //пустая строка
	var hello string = "Hello\n\t"
	var world string = `World`
	var str = "Привет"
	str1 := "你好，世界" //UTF-8 из коробки
	fmt.Println(s, hello, world, str, str1)

	//rune - 32-битное целое число

	var Byte byte = '\x27' //(uint8)
	/* Это значение представляет собой символ, код которого в шестнадцатеричном формате равен 27.
	В ASCII и UTF-8 символ с кодом 27 — это одинарная кавычка (').
	Таким образом, Byte будет равен 39 (в десятичной системе), поскольку код символа ' равен 39. */
	var Rune rune = 'A' //(uint32) для UTF-8 символов
	fmt.Println(Byte, Rune)

	helloWorld := "Привет, мир!"
	iamHere := helloWorld + "Я тут!"
	fmt.Println(iamHere)

	//строки не изменяемые!

	byteLen := len(str)                    //длина в байтах
	runeLen := utf8.RuneCountInString(str) //длина в рунах
	fmt.Println(byteLen, runeLen)

	//получение подстроки в байтах
	str2 := str[:4] // Пр
	S := str[0]     // 208
	fmt.Println(str2, S)

	byteStr := []byte(str)     // слайс байт
	strByte := string(byteStr) // строка из слайса байт
	fmt.Println(byteStr, strByte)
}
