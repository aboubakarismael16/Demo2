package main

import (
	"fmt"
)

//user defines a user in the program
type user struct {
	name string
	email string
}
func (u *user) notify()  {
	fmt.Printf("Sending user email to " +
		"%s<%s>\n",u.name,u.email)

}

//admin represents an admin user with privileges
type admin struct {
	user //embedded type
	level string
}


//let's change the sample by adding an interface
//notifier is an interface that defined notification
//type behavior
type notifier interface {
	notify()
}

//sendNotification accepts values that implement the notifier
//interface and sends notifications.
func sendNotification(n notifier)  {
	n.notify()

}


//notifier implements a method that can be called via
//a value of type admin
func (u *admin) notify()  {
	fmt.Printf("Sending admin email to " +
		"%s<%s>\n",u.name,u.email)

}

func main() {
	//create an admin user
	ad := admin {
		user : user {
			name: "ismael",
			email: "ismael@email.com",
		},
		level : "Super",
	}

	//we can access the inner type's method directly
	ad.user.notify()

	//the inner type's method is promoted
	//ad.notify()

	//send the admin user a notification.
	//The embedded inner type's implementation of the
	//interface is "promoted" to the outer type
	sendNotification(&ad)

	//the inner type's method is NOT promoted
	ad.notify()

	
}
