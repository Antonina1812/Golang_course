package async

import "fmt"

func Select1() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int)

	ch1 <- 10

	select {
	case val := <-ch1: // читаем из канала
		fmt.Println("ch1", val)
	case ch2 <- 1: // пишем в канал
		fmt.Println("put val to ch2")
	default:
		fmt.Println("default")
	}
}

func Select2() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	ch2 := make(chan int, 2)
	ch2 <- 3

loop:
	for { // цикл while
		select { // каналы для чтения из канала выбираются планировщиком го рандомно
		case v1 := <-ch1:
			fmt.Println("ch1", v1)
		case v2 := <-ch2:
			fmt.Println("ch2", v2)
		default:
			break loop
		}
	}

}

func Select3() {
	cancelCh := make(chan struct{}) // канал отмены, пустая структура (нет полей => ничего нет)
	dataCh := make(chan int)

	go func(cancelCh chan struct{}, dataCh chan int) {
		val := 0
		for {
			select {
			case <-cancelCh: // читаем из канала
				return
			case dataCh <- val: // записываем в канал
				val++
			}
		}
	}(cancelCh, dataCh)

	for curVal := range dataCh {
		fmt.Println("read", curVal)
		if curVal > 3 {
			fmt.Println("send cancel")
			cancelCh <- struct{}{}
			break
		}
	}
}

func SelectPrint() {
	//Select1()
	//Select2()
	Select3()
}
