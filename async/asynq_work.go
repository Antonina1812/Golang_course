package async

import (
	"fmt"
	"time"
)

func getComments() chan string {
	result := make(chan string, 1) // канал дожен быть буфферизированным (хотя бы на одно значение), иначе будет утечка памяти
	go func(out chan<- string) {
		time.Sleep(2 * time.Second)
		fmt.Println("asynq operation ready") // потом выполнится это
		out <- "32 comment"
	}(result)
	return result
}

func getPage() {
	resultCh := getComments()

	time.Sleep(1 * time.Second)
	fmt.Println("get related articles") // сначала выполнится это

	commentsData := <-resultCh
	fmt.Println("main goroutine: ", commentsData) // в конце выполнится это
}

func AsynqWorkPrint() {
	for i := 0; i < 3; i++ {
		getPage()
	}
}
