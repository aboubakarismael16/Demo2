package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad = 10
)

var bu sync.WaitGroup

func worker(tasks chan string,worker int)  {
	//Report that we just returned
	defer bu.Done()

	for {
		//wait for work to be assigned
		task,ok := <- tasks
		if !ok {
			//this means the channel is empty and closed
			fmt.Printf("Worker: %d : Shutting Down\n",worker)
			return
		}

		//display we are starting the work
		fmt.Printf("Worker: %d : Started %s\n",worker,task)

		//Randomly wait to simulate work time
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		//display we finished the work
		fmt.Printf("Worker: %d : Completed %s\n",worker,task)
	}

}

func main() {
	//create a buffered channel to manage the task load
	tasks := make(chan string,taskLoad)

	//launch goroutine to handle the work
	bu.Add(numberGoroutines)
	for gr := 1 ; gr <= numberGoroutines ; gr++{
		go worker(tasks,gr)
	}

	//add a bunch of work to get done
	for post := 1 ; post <= taskLoad ; post++ {
		tasks <- fmt.Sprintf("Task: % d",post)
	}

	//close the channel so the goroutines will quit
	//when all the work is done
	close(tasks)

	//wait for all the work to get done
	bu.Wait()

}
