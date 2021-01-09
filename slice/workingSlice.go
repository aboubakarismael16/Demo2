package main

import (
	//"fmt"
	"fmt"
	pretty "github.com/inancgumus/prettyslice"
)

func main(){

	// k=5,i he j cap = k-i,j-i
	slice1 := []int{10,20,30,40,50}
	//slice1[1] = 25
	//pretty.Show("slice1",slice1)

	fmt.Println("taking the slice of slice")

	newSlice := slice1[1:3]
	pretty.Show("newSlice",newSlice)

	fmt.Println("Potential consequence of making change to a slice")
	//newSlice[3] = 35
	newSlice[1] = 35
	pretty.Show("New newSlice",newSlice)
	pretty.Show("old slice also takes some change",slice1)
	fmt.Printf("slcie %p\n",&slice1)
	fmt.Printf("newslcie %p\n",&newSlice)

	//using append to add an element to a slice
	newSlice = append(newSlice,60)
	pretty.Show("appended newSlice",newSlice)

	newSlice = append(slice1,60)
	pretty.Show("increase len and cap of slice",newSlice)
	
}
