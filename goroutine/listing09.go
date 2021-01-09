package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)


var (
	counter int64

	sy sync.WaitGroup
	mutex sync.Mutex

)

func iCounter(id int )  {
	defer sy.Done()
	for count := 0 ; count <2 ; count++ {

		mutex.Lock()

		{
			value:= counter
			//capture the value of value
			atomic.AddInt64(&counter,1)

			////yield the thread and be value placed back in queue
			runtime.Gosched()
			//
			////increment our local value of count
			value++
			//
			////store the value back into count
			counter =  value
		}
		mutex.Unlock()
	}
}

func main() {

	sy.Add(2)

	go iCounter(1)
	go iCounter(2)

	sy.Wait()

	fmt.Println("Final counter: ",counter)

}
