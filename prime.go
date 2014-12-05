package main

import (
	"log"
	"os"
	"runtime"
)

func GeneratePrimes(out chan<- int) {
	primeSource := make(chan int)
	go generate(primeSource)

	for {
		prime := <-primeSource

		out <- prime

		newPrimeSource := make(chan int)
		go filter(primeSource, newPrimeSource, prime)
		primeSource = newPrimeSource
		//output while testing
		log.Println(prime)
	}
}

func generate(out chan<- int) {
	i := 2
	for {
		out <- i
		i++
	}
}

func filter(in <-chan int, out chan<- int, prime int) {
	for i := range in {
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)

	out := make(chan int)
	go GeneratePrimes(out)
}
