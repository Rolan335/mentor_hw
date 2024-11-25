package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func calcRow(i int, input int, sigNotice <-chan struct{}, sumCh chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum float64
	for {
		select {
		case <-sigNotice:
			sumCh <- sum
			return
		default:
			sum += math.Pow(float64(-1), float64(i)) / float64(2*i+1)
			i += input
		}
	}
}
func main() {
	input := *flag.Int("goroutines", 1, "enter num of goroutines")
	flag.Parse()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	sigNotice := make(chan struct{}, input)
	sumCh := make(chan float64, input)
	var wg sync.WaitGroup

	for i := 0; i < input; i++ {
		wg.Add(1)
		go calcRow(i, input, sigNotice, sumCh, &wg)
	}

	//блокируем горутину main до получения сигнала signal.notify(). При получении signal.notify отправляем в сигнализирующий канал который
	//передаётся в горутины столько раз, сколько у нас горутин.
	//Логика завершения горутин со строчки 19 и wg.Done()
	<-sig
	for i := 0; i < input; i++ {
		sigNotice <- struct{}{}
	}
	close(sig)
	wg.Wait()

	var sum float64
	for i := range sumCh {
		sum += i * 4
		//Если закрывать канал после цикла, send on closed channel
		if len(sumCh) == 0 {
			close(sumCh)
		}
	}
	fmt.Println(sum)
}
