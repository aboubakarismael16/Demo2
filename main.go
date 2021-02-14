package main

import "fmt"

func main() {
	// for i :=0; i < 10; i++ {
	// 	fmt.Printf("%d\n",i)
	// }
	i := 0
I :
		fmt.Printf("%d\n",i)
		i++
		if i < 10 {
			goto I
		}
}