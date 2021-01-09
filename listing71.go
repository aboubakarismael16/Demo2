package main

import (
	"demo2/entities"
	"fmt"
)

func main() {
	u := entities.Admin{
		Rights: 18006150105,
	}

	//set the exported fields from the unexported
	//inner type
	u.Name = "Ismael"
	u.Email = "ismael@email.com"
	fmt.Printf("User : %v\n",u)

}
