package pi

import (
	"fmt"
	"math"
	"os"
	"sync"
)

func FindPi(i int, input int, sig <-chan os.Signal, chSum chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum float64
loop:
	for {
		select {
		case <-sig:
			fmt.Println("recieved signal")
			chSum <- sum
			break loop
		default:
			sum += math.Pow(float64(-1), float64(i)) / float64(2*i+1)
		}
	}
}

// 3.034991s
// 2.8570899s
