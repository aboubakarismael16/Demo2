package main

import (
	"fmt"
	"time"
)



type Employee struct {
	ID int
	name string
	Address string
	DoB time.Time
	Position string
	Salary int
	Manager int
}


func main() {
	const salary = 2000
	//make a variable,then call the type struct
	var ismael Employee
	ismael.name = "Ismael"
	ismael.Position = "Mid-level"
	//fmt.Println("")
	ismael.Salary = salary + (salary*0.2)
	position := &ismael.Position

	//the dot notation also works with a pointer to a struct
	*position = "Senior" + *position
	fmt.Println(ismael.name)
	fmt.Println(ismael.Salary)
	fmt.Println(ismael.Position)

	isamel2 := Employee{
		name: "isamel2",
	}
	fmt.Println(isamel2.name)



}
