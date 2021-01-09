package main

import (
	"fmt"
	pretty "github.com/inancgumus/prettyslice"
)

func main()  {

	nums := []int{1, 9, 5, 6, 4, 8}
	pretty.Show("nums", nums)

	var array1 [3]*string

	array2 := [3]*string{new(string),new(string),new(string)}

	*array2[0] = "red"
	*array2[1] = "blue"
	*array2[2] = "green"

	array1 = array2
	//fmt.Print(array1)

	pretty.Show("array",array1)

	fmt.Println("multidimensional")

	array3 := [4][2]int{{10,11},{20,21},{30,31},{40,41}}
	array3 = [4][2]int{1:{10,11},3:{40,41}}
	array3 = [4][2]int{1:{0:11},3:{1:41}}

	pretty.Show("array3",array3)

	fmt.Println("access to element in two-dimensional array")

	array3[0][0] = 10
	array3[0][1] = 20
	array3[1][1] = 30

	pretty.Show("array3",array3)

	fmt.Println("passing array between function")
	//allocate an array of 8 megabytes
	var array4 [1e6] int

	//pass the address of the array to the function
	foo(&array4)


}

//function foo accepts a pointer to an array of one million integers
func foo(array *[1e6]int)  {

}
