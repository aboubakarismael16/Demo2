package main

import (
	"fmt"
	"sync"
	"time"
)

var bt sync.WaitGroup

func Runner(baton chan int) {
	var newRunner int

	//wait to receive the baton
	runner := <- baton

	//start running around to the track
	fmt.Printf("Runner %d running with Baton\n",runner)

	//new runner to hte line
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To the Line\n",newRunner)

		go Runner(baton)

	}

	//running around the track
	time.Sleep(100 * time.Millisecond)

	//is the race over
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race over \n",runner)
		bt.Done()
		return
	}

	//exchange the baton for the next runner
	fmt.Printf("Runner %d Exchange With Runner %d\n",runner,newRunner)
	baton <- newRunner


}

func main() {

	//create an unbuffered channel
	baton := make(chan int)

	//add a count of one for the last runner
	bt.Add(1)

	//first runner to his mark
	go Runner(baton)

	//start the race
	baton <- 1

	//wait for the race to finish
	bt.Wait()

	
}
