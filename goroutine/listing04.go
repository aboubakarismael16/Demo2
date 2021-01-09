package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	//allocate a logical processor for every available core
	//runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(1)

	//add count of two ,one for each goroutine
	wg.Add(2)

	//create two goroutines
	fmt.Println("Create Goroutines")

	go printPrime("A")
	go printPrime("B")

	//wait fro the goroutine to finish.
	fmt.Println("Waiting To finish")
	wg.Wait()

	fmt.Println("Terminating Program")


}

//printPrime displays prime numbers for the first 5000 numbers
func printPrime(prefix string)  {
	//schedule the call to Done to tell main we are done
	defer wg.Done()

next :
	for outer := 2 ;outer < 20 ;outer++ {
		for inner := 2 ; inner < outer ; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n",prefix,outer)
	}
	fmt.Println("Completed",prefix)

}
