package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Prime Numbers between 1 and 100")
	out := make(chan int)
	var wg sync.WaitGroup

	go printer(out)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		// fmt.Printf("range : %d to %d\n", i * 10, (i+1) * 10)
		go arePrimes(i*10, (i+1)*10, out, &wg)
	}
	wg.Wait()
	close(out)
	// panic("testing")
}

func printer(in <-chan int) {
	for {
		select {
		case val, ok := <-in:
			if !ok {
				fmt.Println("terminating printer")
				return
			} else {
				fmt.Printf("%d ", val)
			}
		}
	}
}

func arePrimes(low, high int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	// fmt.Printf("executing : %d to %d\n", low, high)
	for n := low; n < high; n++ { // Test for each number
		if isPrime(n) {
			out <- n
		}
	}

}

func isPrime(n int) bool {
	if n == 0 || n == 1 {
		return false
	}

	for j := 2; j < n; j++ {
		if n%j == 0 {
			return false
		}
	}
	return true
}
