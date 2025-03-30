package async

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	iteretionsNum = 7
	goroutinesNum = 5
)

func doSmth(in int) {
	for j := 0; j < iteretionsNum; j++ {
		fmt.Print(formatFunc(in, j))
		runtime.Gosched() // планировщик задач
	}
}

func formatFunc(in, j int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		"th", in,
		"iter", j, strings.Repeat("■", j))
}

func GoroutinesPrint() {
	// из горутины нельзя вернуть значение, для этого используются каналы
	for i := 0; i < goroutinesNum; i++ {
		go doSmth(i)
	}
	fmt.Scanln()
}
