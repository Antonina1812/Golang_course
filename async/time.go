package async

import (
	"fmt"
	"time"
)

func Timer() {
	timer := time.NewTimer(1 * time.Second)

	select {
	case <-timer.C: // .C - это канал
		fmt.Println("timer.C timeout happend")
	case <-time.After(time.Minute):
		fmt.Println("time.After timeout happend")
	}
}

func Ticker() {
	ticker := time.NewTicker(time.Second) // промежуток времени, через который что-то повторяется
	i := 0
	for tickTime := range ticker.C {
		i++
		fmt.Println("step", i, "time", tickTime)
		if i >= 5 {
			ticker.Stop()
			break
		}
	}
	fmt.Println("total", i)
}

func AfterFunc() {
	timer := time.AfterFunc(1*time.Second, Hello) // запускает функцию через промежуток времени
	fmt.Scanln()
	timer.Stop()
}

func Hello() {
	fmt.Println("hello")
}

func TimePrint() {
	//Timer()
	//Ticker()
	AfterFunc()
}
