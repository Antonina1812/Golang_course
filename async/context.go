package async

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Контекст в Go — это объект, который передаётся в функции и горутины для управления временем выполнения.
// Он помогает контролировать тайм-ауты, отменять долгие запросы и синхронизировать работу горутин

func worker(ctx context.Context, workerNum int, out chan<- int) {
	waitTime := time.Duration(rand.Intn(100)+10) * time.Millisecond
	fmt.Println(workerNum, "sleep", waitTime)
	select {
	case <-ctx.Done():
		return
	case <-time.After(waitTime):
		fmt.Println("worker", workerNum, "done")
		out <- workerNum
	}
}

func TimeOut() {
	workTime := 50 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), workTime)
	defer cancel()

	result := make(chan int, 1)

	for i := 0; i <= 10; i++ {
		go worker(ctx, i, result)
	}

	totalFound := 0

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case foundBy := <-result:
			totalFound++
			fmt.Println("result found by", foundBy)
		}
	}

	fmt.Println("total found", totalFound)
	time.Sleep(time.Second)
}

func ContextPrint() {
	// ctx, finish := context.WithCancel(context.Background()) // контекст и функция отмены
	// result := make(chan int, 1)

	// for i := 0; i <= 10; i++ {
	// 	go worker(ctx, i, result)
	// }

	// foundBy := <-result
	// fmt.Println("result found by", foundBy)
	// finish()

	// time.Sleep(time.Second)

	TimeOut()
}
