package main

import "fmt"

func myFunc(a interface{})  {
	// Extracting the value of a
	//when we change a.(string) by a.(int) we will get panic
	val := a.(string)
	fmt.Println("Value : ",val)

}

func myFunc2(a interface{})  {
	//when we change a.(string) by a.(int) we will get panic
	// so to overcome this panic we use this syntax val,ok:=
	val,ok := a.(float64)
	fmt.Println(val,ok)

}

func myFunc3(a interface{})  {
	switch a.(type) {
	case int :
		fmt.Println("Type int , Value :",a.(int))
	case string :
		fmt.Println("\nType string, Value :",a.(string))
	case float64:
		fmt.Println("\nType float, Value :",a.(float64))
	default:
		fmt.Println("\nType Not found")

	}

}

func main() {

	var val interface{} = "Hi there"
	myFunc(val)

	//illustrate type assertion
	var a1 interface{} = 98.4
	myFunc2(a1)
	var a2 interface{} = "ismael"
	myFunc2(a2)


	fmt.Println()
	fmt.Println("switch")
	//switch
	myFunc3(12)
	myFunc3("ismael")
	myFunc3(true)

}
