package pi

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func calcRow(wg *sync.WaitGroup, i int, sumCh chan float64) {
	defer wg.Done()
	sum := math.Pow(float64(-1), float64(i)) / float64(2*i+1)
	sumCh <- sum
}

func Pi(input int) float64 {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	var wg sync.WaitGroup
	var sum float64
	var sumCh chan float64
	rowCounter := 0
	ticker := 0
	for {
		select {
		case <-sig:
			close(sig)
			wg.Wait()
			fmt.Println(ticker * input)
			return sum * 4
		default:
			ticker++
			sumCh = make(chan float64, input)
			for i := 0; i < input; i++ {
				wg.Add(1)
				go calcRow(&wg, rowCounter+i, sumCh)
			}
			wg.Wait()
			close(sumCh)
			for v := range sumCh {
				sum += v
			}
			rowCounter += input
		}
	}
}

//Количество итераций не синхронизировано
// func calcRow(i int, input int, sig <-chan os.Signal, sumCh chan<- float64, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	var sum float64
// 	ticker := 0
// 	//считаем свои числа ряда в default до того, как получим сигнал syscall
// 	for {
// 		select {
// 		case <-sig:
// 			sumCh <- sum
// 			fmt.Println(ticker)
// 			return
// 		default:
// 			sum += math.Pow(float64(-1), float64(i)) / float64(2*i+1)
// 			i += input
// 			ticker++
// 		}
// 	}
// }

// func Pi(input int) float64 {
// 	sig := make(chan os.Signal, 1)
// 	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
// 	sumCh := make(chan float64, input)
// 	var wg sync.WaitGroup

// 	for i := 0; i < input; i++ {
// 		wg.Add(1)
// 		go calcRow(i, input, sig, sumCh, &wg)
// 	}
// 	<-sig

// 	close(sig)
// 	wg.Wait()
// 	close(sumCh)

// 	var sum float64
// 	for i := range sumCh {
// 		sum += i
// 	}
// 	sum *= 4
// 	return sum
// }
