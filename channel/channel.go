package main

import "fmt"

func main() {
	//unbuffered channel of integers
	unbuffered := make(chan int) //0xc00001a1e0
	fmt.Println(unbuffered)
	//buffered channel of strings
	buffered := make(chan string,10) //0xc00005c180
	fmt.Println(buffered)

	//send a string a through of string we should use <- operator
	buffered <- "Gopher"
	fmt.Println(buffered) //0xc00005c180

	//receive a string from the channel
	value := <- buffered
	fmt.Println(value) //Gopher



}
