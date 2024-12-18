package fanIn

import "sync"

//Fan_in
func Merge(in ...<-chan string) <-chan string {
	resCh := make(chan string, len(in))

	var wg sync.WaitGroup

	for _, ch := range in {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				resCh <- v
			}
		}()
	}
	go func() {
		wg.Wait()
		close(resCh)
	}()

	return resCh
}

//// Broadcaster
// func Split(mainCh <-chan string, chCount int) []chan string {
// 	chArr := make([]chan string, 0, chCount)
// 	for range chCount {
// 		newCh := make(chan string)
// 		chArr = append(chArr, newCh)
// 	}
// 	go func() {
// 		defer func() {
// 			for _, vArr := range chArr {
// 				close(vArr)
// 			}
// 		}()
// 		for vMain := range mainCh {
// 			for _, vArr := range chArr {
// 				vArr <- vMain
// 			}
// 		}
// 	}()
// 	return chArr
// }

func Split(mainCh <-chan string, chCount int) []chan string {
	chArr := make([]chan string, 0, chCount)
	for range chCount {
		newCh := make(chan string)
		chArr = append(chArr, newCh)
	}
	go func() {
		defer func() {
			for _, vArr := range chArr {
				close(vArr)
			}
		}()
		chArrCounter := 0
		for vMain := range mainCh {
			chArr[chArrCounter] <- vMain
			chArrCounter++
			if chArrCounter == len(chArr) {
				chArrCounter = 0
			}
		}
	}()
	return chArr
}
