package main

import "fmt"

//notifier is an interface that defined notification
//type behavior
type notifier interface {
	notify()
}

//user defines a user in the program
type user struct {
	name string
	email string
}

func (u *user) notify()  {
	fmt.Printf("Sending user email to " +
		"to %s<%s>\n",u.name,u.email)

}

//admin defines a admin in the program
type admin struct {
	name string
	email string
}

func (u *admin) notify()  {
	fmt.Printf("Sending admin email to " +
		"to %s<%s>\n",u.name,u.email)

}



//sendNotification accepts values that implement the notifier
//interface and sends notifications.
func sendNotification(n notifier)  {
	n.notify()

}

//duration is a type with a base type of int
type duration int


//format pretty-prints the duration value
func (d *duration) pretty() string  {
	return fmt.Sprintf("Duration: %d",*d)
}

func main() {
	//create a user value and pass it to sendNotification
	ismael := user{"ismael","ismael@gmail.com"}
	sendNotification(&ismael)

	//create a user value and pass it to sendNotification
	aboubakar := admin{"aboubakar","aboubakar@email.com"}
	sendNotification(&aboubakar)

	//cannot call pointer method on duration(45)
	//cannot take the address of duration(45)
	//duration(45).pretty()

}
