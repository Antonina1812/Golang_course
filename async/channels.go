package async

import "fmt"

func ChannelsWithLoop() {
	in := make(chan int)

	go func(out chan<- int) { // out chan - в канал можно и писать и читать, out <-chan - из канала можно только читать, out chan<- - в канал можно только писать
		for i := 0; i <= 4; i++ {
			fmt.Println("before", i)
			out <- i
			fmt.Println("after", i)
		}
		close(out) // закрытие канала завершает цикл
		fmt.Println("finish")
	}(in)

	for i := range in {
		fmt.Println("\tget", i)
	}
}

func ChannelsPrint() {
	ChannelsWithLoop()

	ch := make(chan int, 1) // chan - канал (позволяет безопасно обмениваться даннми между горутинами)
	// make по-умолчанию создаёт небуфферизированный канал (плохо, может возникнуть deadlock, т е горутина будет ждать, когда из канала вычитают значение, но не дождётся)
	// его можно сделать буфферизированным, добавив число через запятую (сколько ещё значений мы хотим хранить)

	go func(in chan int) {
		val := <-in // вычитывание из канала (стрелка слева от канала)
		fmt.Println("go get from chan", val)
		fmt.Println("go after from chan")
	}(ch)

	ch <- 42 // кладём в канал (стрелка справа от канала)
	ch <- 100

	fmt.Println("go after put into chan")
	fmt.Scanln()
}
