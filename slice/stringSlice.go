package main

import (
	"fmt"
	pretty "github.com/inancgumus/prettyslice"
)

func main() {
	source := []string{"Apple","Orange","Plum","Banana","Grape"}
	pretty.Show("source",source)


	//performing on three-index slice
	slice := source[2:3:3]
	//if you attempt to set a cap larger than the available cap
	//you will get a rune time error
	pretty.Show("performing on three-index slice",slice)

	slice = append(slice,"kiwi")
	pretty.Show("appended slice",slice)


	//appending to slice form another slice
	slice1 := []int{1,2}
	slice2 := []int{3,4}
	pretty.Show("new appended slice",append(slice1,slice2...))

	//iterating over a slice using for range
	slice3 := []int{10,20,30,40,50}
	for i,v := range slice3 {
		fmt.Printf("index: %d Value: %d\n",i,v)
	}

	fmt.Println()
	fmt.Println("for loop")
	//if you need more control iterating over a slice
	//you can always use a traditional for loop
	for i := 2 ;i <len(slice3);i++{
		fmt.Printf("index : %d Value : %d\n",i,slice3[i])
	}


	//declaring a multidimensional slice
	slice4 := [][]int{{10},{100,200},{30,40,45}}
	pretty.Show("multidimensional slice",slice4)
	slice4[0] = append(slice4[0],20)
	pretty.Show("appended slice4",slice4)


}
