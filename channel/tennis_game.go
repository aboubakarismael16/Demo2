package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func player(name string,court chan int)  {

	defer wg.Done()

	for {
		//wait for the ball to be hit back to us
		ball , ok := <- court
		if !ok {
			//if the channel was closed we won
			fmt.Printf("Player %s Won\n",name)
			return
		}

		//pick a random number ans see if we miss the ball
		n := rand.Intn(100)
		if n %13 == 0 {
			fmt.Printf("Player %s Missed\n",name)

			//close the channel to signal we lost
			close(court)
			return
		}

		//display and then increment the hit count by one
		fmt.Printf("Player %s Hit : %d\n",name,ball)
		ball++

		//Hit he ball back to the opposing player
		court <- ball
	}

}

func init()  {
	rand.Seed(time.Now().UnixNano())

}

func main() {
	//create an unbuffered channel
	court := make(chan int)

	wg.Add(2)

	//launch two player
	go player("Nadal",court)
	go player("Djokovic",court)

	//start the set
	court <- 1

	//wait for the game finished
	wg.Wait()
	
}
