package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//allocate 1 logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)

	//sy is used to wait fort he program to finish
	//add a count of two,one fro each goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutine")




	//declare an anonymous function and create a goroutine
	go func() {
		//schedule the call to Done to tell main we are done
		defer wg.Done()

		//display the alphabet three times
		for count := 0 ; count < 3; count++{
			for char := 'a';char < 'a' + 26 ; char++ {
				fmt.Printf("%c ",char)
			}
		}
	}()


	go func() {
		//schedule the call to Done to tell main we are done
		defer wg.Done()

		//display the alphabet three times
		for count := 0 ; count < 3; count++{
			for char := 'A';char < 'A' + 26 ; char++ {
				fmt.Printf("%c ",char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Waiting To finished")

	}()

	//fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}
