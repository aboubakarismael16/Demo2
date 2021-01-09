package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	at sync.WaitGroup
)

func doWork(name string)  {
	defer at.Done()

	for {
		fmt.Printf("Doing %s Work\n",name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("shutting %s Down\n",name)
			break
		}
	}

}

func main() {

	at.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)

	fmt.Println("Shutting Now")
	atomic.StoreInt64(&shutdown,1)

	at.Wait()
	
}
