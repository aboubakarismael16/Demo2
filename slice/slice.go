package main

import (
	"fmt"
	pretty "github.com/inancgumus/prettyslice"
)

func main() {
	//slice :=[]int{2,5}
	slice := make([]string,5)
	pretty.Show("slice ",slice)

	slice1 := make([]string,3,5)
	pretty.Show("slice1 ",slice1)

	slice01 := make([]int64,3,5)
	pretty.Show("int3",slice01)
	slice01 = append(slice01,1,2,3,4,5,6)
	pretty.Show("appended int3",slice01)

	fmt.Println("declaring a slice with a slice literal")

	slice2 := []string{"red","blue","green","yellow","pink"}
	pretty.Show("slice2",slice2)

	//use make to create an empty slice of integers
	slice3 := make([]int,0)
	//also can use slice literal to create an empty
	slice3 = []int{}
	pretty.Show("slice3 ",slice3)


	//passing slice between functions
	slice5 := make([]int,6)
	result := foo(slice5)

	pretty.Show("slice5",result)

}

func foo(slice []int)  []int {
	return slice
}
