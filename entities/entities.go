package entities

//User defines a user in the program
type user struct {
	Name string
	Email string
}

//Admin defines a admin in the program
type Admin struct {
	user    //embedded hte type is unexported
	Rights int
}