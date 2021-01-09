package main

import (
	"fmt"
)

//user defines in the program
type user struct {
	name string
	email string
}

//notify implements a method with a value receiver
func (u user) notify()  {
	fmt.Printf("Sending User Email To %s<%s>\n",u.name,u.email)

}

//changeEmail implements a method with a value receiver
func (u *user)  changeEmail(email string) {
	u.email = email
}


func main() {
	//value of type user can be used to call method
	//declared with a value receiver
	ismael := user{
		name: "Ismael",
		email: "ismael@gmail.com",
	}
	ismael.notify()

	//pointers of type user can be used to call method
	//declared with a value receiver
	fred := &user{
		name: "Fred",
		email: "fred@gmail.com",
	}
	fred.notify()


	(ismael).changeEmail("aboubakar@gmail.com")
	ismael.notify()

	fred.changeEmail("freddy@gmail.com")
	(*fred).notify()

	
}
