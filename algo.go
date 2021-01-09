package main

import (
	"fmt"
)

func takeCupCake(CupeCake []int) int {
	fmt.Println("Have a CupeCake number: ", CupeCake[0])
	return CupeCake[0]
}

func eatChips(blowOfChips int)  {

	if chip := 0 ; chip < blowOfChips {
		fmt.Println("Have some chips")
		return
	}else {
		fmt.Println("No more chips")
	}

}

func pizzaDelivery(boxeDelivered int)  {

	for pizzaBox := 0 ; pizzaBox < boxeDelivered; pizzaBox++{
		//for pizza := 0 ; pizza <= pizzaBox; pizza ++{
		//	for slice := 0; slice <= pizza ; slice++ {
				fmt.Println("Pizza is Here," +
					"there are  boxes delivered: ",boxeDelivered)
				return
		//	}
		//}
	}
	fmt.Println("OOps! No pizza")
}

func eatMeat(meat int) int  {
	if meat == 0 {
		fmt.Println("Someone ate all the meat")
		return meat
	}
	fmt.Println("Eating meat...")
	return eatMeat( meat-1)

}

func giveForks(diners []int) []int {

	var workForks []int
	var first ,second,third int

	for i := range diners {
		//var max ,maxIndex int

			if diners[i] > first{
				third = second
				second = first
				first = diners[i]
			} else if diners[i] > second {
				third = second
				second = diners[i]
			} else if diners[i] > third{
				third = diners[i]
		}

	}
	workForks = append(workForks,first,second,third)
	fmt.Println(workForks)
	return workForks
}


func main() {
	cupe := []int{12}
	takeCupCake(cupe)
	eatChips(12)
	pizzaDelivery(0)
	eatMeat(3)

	diners := []int{2,88,87,42,16,10,34,1,43,100}
	giveForks(diners)

}
