package main


import (
	"github.com/joho/godotenv"
	"gopkg.in/gookit/color.v1"
)

func main() {

	godotenv.Load()

	color.Blue.Println("Chad my country")
	color.Yellow.Println("Chad my country")
	color.Red.Println("Chad my country")
}
