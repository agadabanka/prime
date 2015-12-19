package main

import (
	"flag"
	"fmt"	
	"math"
	"runtime"
	"sync"
	"bufio"
	"os"
	"strconv"

	ut "github.com/agadabanka/gotools/mathutil"
)

type PrinterInput struct {
	low, high int
	in, stop  <-chan int
}

func main() {
	// outChan := make(chan int)
	// stopChan := make(chan int)
	// lowLimit, highLimit := getLimits()
	// p := &PrinterInput{low: lowLimit, high: highLimit, in: outChan, stop: stopChan}
	// go printer(p)
	// computePrimes(lowLimit, highLimit, outChan) // blocks till completion
	
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if number, err := strconv.Atoi(input.Text()); err != nil {
			fmt.Printf("please input a number. error : %v\n", err)
			continue
		} else {			
			fmt.Printf("Prime Factors for %d: (%s)\n", number, GetPrimeFactors(number).String())
		}		
	}
	// stopChan <- 1
	// panic("testing")
}

func getLimits() (low, high int) {
	argsLow := flag.Int("low", 0, "lower limit")
	argsHigh := flag.Int("high", 1000, "uppper limit")
	flag.Parse()
	low = ut.Max(0, *argsLow)
	high = ut.Max(low, ut.Min(1000, *argsHigh))
	return low, high
}

func computePrimes(low, high int, out chan int) {
	var wg sync.WaitGroup

	batchsize := int(math.Ceil(float64(high-low) / float64(runtime.NumCPU()))) // divide by GOMAXPROCS
	for i := low; i < high; i += batchsize {
		wg.Add(1)
		// fmt.Printf("\nrange : %d to %d\n", i, min(high, i + batchsize))
		go arePrimes(i, ut.Min(high, i+batchsize), out, &wg)
	}
	wg.Wait()
}

func printer(p *PrinterInput) {
	low, high := p.low, p.high
	in, stop := p.in, p.stop
	fmt.Printf("Prime Numbers between %d and %d\n", low, high)
	for {
		select {
		case val, ok := <-in:
			if !ok {
				fmt.Println("in closed: terminating printer")
				return
			} else {
				fmt.Printf("%d ", val)
			}
		case <-stop:
			fmt.Println("stop: terminating printer")
			return
		}
	}
}

func arePrimes(low, high int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
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
