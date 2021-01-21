package main

import (
	"fmt"
	"runtime"
	"sync"
)

func dosomeThing(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)

}

const maxUrlCount = 1000
const maxConcurrent = 5

func crawl(name string, queue <-chan int, wg *sync.WaitGroup) chan bool {
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Job: ", name, " quited.")
				break
			case count := <-queue:
				s := fmt.Sprintf("%s is running for url: %d", name, count)
				fmt.Println(s)
				wg.Done()

			}
			runtime.Gosched()
		}
	}()

	return quit
}

func main() {
	wg := new(sync.WaitGroup)
	queueCh := make(chan int, maxUrlCount) // max 1000
	quiters := make([]chan bool, maxConcurrent)

	for i := 1; i <= maxUrlCount; i++ {
		queueCh <- i
		wg.Add(1)
	}

	for i := 1; i <= maxConcurrent; i++ {
		quiters[i-1] = crawl(fmt.Sprintf("Job %d", i), queueCh, wg)
	}

	wg.Wait()

	for i := 0; i < maxConcurrent; i++ {
		quiters[i] <- true
	}

	fmt.Println("DONE")
}
