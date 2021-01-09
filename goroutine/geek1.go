package main

import (
	"fmt"
	"time"
)

func display(str string)  {
	for w := 0; w < 10 ; w++ {
		time.Sleep(1*time.Millisecond)
		fmt.Println(str)
	}

}

func main() {
	//calling the goroutine
	//go display("Hello Goroutine")
	//
	////calling the normal function
	//display("Welcome to goroutine")


	fmt.Println()
	//creating a anonymous goroutine
	fmt.Println("Welcome to the main function")

	//anonymous goroutine start with go func(){}()
	go func() {
		fmt.Println("Welcome to Anonymous Goroutine")
	}()

	time.Sleep(1 * time.Millisecond)
	fmt.Println("Goodbaye !! to hte main function")

	
}
