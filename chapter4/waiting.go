package chapter4;

import (
	"fmt"
	"sync"
)

func printHello(iteration int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Hello you ;) #%d\n", iteration + 1)
}

// Using waiting groups to demonstrate how to wait for a go routine with out the use of channels
func Waiting() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printHello(i, &wg)
	}
	wg.Wait()
}
