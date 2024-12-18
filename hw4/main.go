package main

import (
	"fmt"
	"hw4/pi"
	"hw4/semaphore"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func runSem() {
	semaphore := semaphore.NewSemaphoreMu(5)
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	// defer cancel()
	timer := time.Now()
	var wg sync.WaitGroup
	for i := range 10 {
		isAcquired := semaphore.TryAcquire(1)
		if isAcquired {
			wg.Add(1)
			fmt.Println(1)
			go func() {
				defer wg.Done()
				defer semaphore.Release(1)
				fmt.Println("goroutine ", i, " start")
				time.Sleep(time.Second)
				fmt.Println("goroutine ", i, " stop")
			}()
		}
	}
	wg.Wait()
	fmt.Println(time.Since(timer))
}

func runPi() {
	piCalc := pi.NewPiCalculator(12, 100_000_000)
	go piCalc.Calc()
	var res float64
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigs:
		res = piCalc.End()
	case <-piCalc.DoneCh():
		res = piCalc.End()
	}
	fmt.Println(res)

}

func main() {
	runPi()
}
