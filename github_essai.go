package main


import (
	"github.com/joho/godotenv"
	"gopkg.in/gookit/color.v1"
)

func main() {

	godotenv.Load()

	color.Blue.Println("Chad is my country")
	color.Yellow.Println("Chad is my country")
	color.Red.Println("Chad is my country")
}
