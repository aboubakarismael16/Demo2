package main

import (
	"demo2/counters"
	"fmt"
)

func main() {

	//create a variable of hte unexported type using the exported
	//New function from the package counters.
	counter := counters.New(10)
	fmt.Printf("counter: %d\n",counter)
}
