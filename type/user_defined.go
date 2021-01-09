package main

import (
	"fmt"
	pretty "github.com/inancgumus/prettyslice"
)

type User struct {
	name        string
	email       string
	matriculate int
	privileged  bool
}


//declaring fields based on other struct types
type admin struct {
	person User //embedding
	level string
}

//declaring int64 type
type Duration int64

func main() {
	var bill User
	fmt.Println(bill.name)

	//declare a variable of type user and initialize all the fields
	ismael := User{
		name:        "ismael",
		email:       "ismael@gmail.com",
		matriculate: 1800045,
		privileged:  true,
	}
	pretty.Show("", ismael.name)
	pretty.Show("", ismael.email)

	fred := admin{
		person: User{
			name: "Fred",
			email: "fred@gmail.com",
			matriculate: 1800046,
			privileged: false,
		},
		level: "super",
	}
	pretty.Show("",fred.person.privileged)
	pretty.Show("",fred.level)

}
