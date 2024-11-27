package pi

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

//runtime.NumCPU() = 12
//1 goroutine 30 sec calc
//251366897 iter
//3.141592657567563
//6 goroutines 30 sec
//140674779 iter
//3.141592652412406
//12 goroutines 30 sec
//128643665 iter
//3.141592652939999

var logicalProcessors int = runtime.NumCPU()

//Можно было изначально задать максимальное значение итераций как количество членов ряда которые нужно вычислить

// Сделано на костыле. Проблема - не синхронизированно количество итераций при кол-ве горутин >1
func calcRow2(i int, input int, sig <-chan os.Signal, sumCh chan<- float64, wg *sync.WaitGroup, maxIter *int) {
	defer wg.Done()
	var sum float64
	ticker := 0
	waitedIters := 0
	//считаем свои числа ряда в default до того, как получим сигнал syscall
	for {
		select {
		case <-sig:
			//когда получен сигнала sigint в другой горутине заканчивается инкрементирование maxIter. Горутина догоняет счётчик для равного кол-ва итераций
			for ; ticker <= *maxIter; ticker++ {
				sum += math.Pow(float64(-1), float64(i)) / float64(2*i+1)
				i += input
			}
			fmt.Println("iters done: ", ticker)
			fmt.Println("waited iters: ", waitedIters)
			sumCh <- sum
			return
		default:
			//если не выступаем за счётчик, считаем ряд, иначе ждём инкрементации maxIter
			if ticker <= *maxIter {
				sum += math.Pow(float64(-1), float64(i)) / float64(2*i+1)
				i += input
				ticker++
			} else {
				waitedIters++
			}
		}
	}
}

func Pi2(input int) float64 {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	sumCh := make(chan float64, input)
	maxIter := 1_000_000
	var wg sync.WaitGroup
	for i := 0; i < input; i++ {
		wg.Add(1)
		go calcRow2(i, input, sig, sumCh, &wg, &maxIter)
	}
	//костыль для контроля итераций. Чтобы при сигнале sigint у горутин было одинаковое количество итераций.
	go func(maxIter *int) {
		for ticker := 0; ; ticker++ {
			select {
			case <-sig:
				//почему при переносе закрытия sig сюда, программа работаем нестабильно
				return
			default:
				// magic formula, каждую 12 итерацию увеличиваем счётчик на 1. Баланс между тем чтобы не тормозить горутины и чтобы счётчик
				// не убегал далеко вперёд ????
				if ticker%logicalProcessors == 0 {
					*maxIter += 1
				}
			}
		}
	}(&maxIter)
	<-sig
	close(sig)
	wg.Wait()
	close(sumCh)

	var sum float64
	for i := range sumCh {
		sum += i
	}
	sum *= 4
	return sum
}
