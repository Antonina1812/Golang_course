package async

import (
	"fmt"
	"sync"
)

// map - конкурентно не безопасный тип данных
// обращаясь к нему из разных горутин, которые могут выполняться на разных процессорах можно словить состояние гонки
// map внутри себя - это ссылка на структуру данных, она может копироваться в другое место. При этом в кэше одного процессора может лежать одно значение, в кэше другого - другое. Когда они попытаются обновить это значение в основной памяти, произойдет конфликт

func RacePrint() {
	var counters = map[int]int{}
	mu := &sync.Mutex{} // ссылка на объект
	for i := 0; i < 5; i++ {
		go func(counters map[int]int, th int, mu *sync.Mutex) {
			for j := 0; j < 5; j++ {
				mu.Lock()
				counters[th*10+j] += 1
				mu.Unlock()
			}
		}(counters, i, mu)
	}
	fmt.Scanln()
	mu.Lock()
	fmt.Println("counters result", counters)
	mu.Unlock()
}

// go run -race main.go
