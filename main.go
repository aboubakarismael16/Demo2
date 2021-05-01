package main

import "fmt"

func average(xs []float64) (ave float64) {
	sum := 0.0
	switch len(xs) {
	case 0:
		ave = 0
	default:
		for _,v := range xs {
			sum += v
		}
		ave = sum / float64(len(xs))
		fmt.Println(ave)
	}

	return
}

func order(a,b int) (int,int) {
	if a > b {
		return b,a
	}
	return a,b
}

type stack struct {
	i int
	data [10]int
}

func (s *stack) push(k int) {
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}

func main() {
	
	// list := []float64{10,20,30,40,50}
	// average(list)
	// // fmt.Println(average(list))
	// a , b := 7,70
	// fmt.Println(order(a,b))

	// var i int
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%v\n",i)
	// }
	// fmt.Printf("%v\n",i)

	var s stack
	s.push(25)
	s.push(14)
	fmt.Printf("Stack %v\n",s)
	s.pop()
	fmt.Printf("Stack %v\n",s)

}