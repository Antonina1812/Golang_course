package async

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

// waitgroup нужен вместо fmt.Scanln()

func StartWorker(in int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < 5; j++ {
		fmt.Println(FormatWork(in, j))
		runtime.Gosched() // runtime.Gosched() позволяет текущей горутине временно освободить процессор, позволяя другим горутинам запуститься
	}
}

func FormatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		"th", in,
		"iter", j, strings.Repeat("■", j))
}

func WaitGroupPrint() {
	wg := &sync.WaitGroup{}
	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1) // добавляем перед заупском горутины
		// wg.Add надо вызывать в той горутине, которая порождает воркеров
		// иначе другая горутина может не успеть запуститься и выполнится Wait
		go StartWorker(i, wg)
	}
	time.Sleep(time.Millisecond)

	wg.Wait()
	fmt.Println(11111)
}
