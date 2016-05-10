package chapter4;

import (
	"fmt"
	"sync"
)

const iterations = 10

func printHello(iteration int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Hello you ;) #%d\n", iteration+1)
}

func Waiting() {
	var wg sync.WaitGroup
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go printHello(i, &wg)
	}
	wg.Wait()
}
