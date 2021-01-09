package main

import (
	"fmt"
	"sync"
)

func task(msg string,wg *sync.WaitGroup)  {
	for i :=0 ; i<5 ; i++ {
		fmt.Println(msg)
	}
	wg.Done()

}

func hello(channel chan bool){
	fmt.Println("Hello Channel function")
	channel <- true

}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go task("first",&wg)
	wg.Add(1)
	go task("second",&wg)
	wg.Wait()

	ch := make(chan bool) // create a boolean type channel

	//run the following function as a goroutine
	go func() {
		fmt.Println("Hello channel")
		ch <- true // pass a boolean value to hte channel
	}()
	<- ch // block until ch receive a boolean value

	ch1 := make(chan bool)
	go hello(ch1)
	<- ch1
	fmt.Println("main function")
}
