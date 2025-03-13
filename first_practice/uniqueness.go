package practice

import (
	"bufio"
	"fmt"
	"io"
	_ "os"
)

func Scanner(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	var prev string
	for in.Scan() {
		txt := in.Text() // txt - это строка
		if txt == prev {
			continue
		}
		if txt < prev {
			return fmt.Errorf("file is not sorted")
		}
		prev = txt
		fmt.Fprintln(output, txt)
	}
	return nil
	// запуск: cat practice/data.txt | go run main.go
}
