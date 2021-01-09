package main

import "fmt"

// Creating an interface
type tank interface {

	// Methods
	Area() float64
	Volume() float64
}

type myValue struct {
	radius float64
	height float64
}

// Implementing methods of
// the tank interface
func (m myValue) Area() float64  {
	return 2 * m.radius * m.height + 2 * 3.14 * m.radius * m.radius

}

func (m myValue) Volume() float64  {
	return 3.14 * m.radius * m.radius * m.height
}



func main() {
	// Accessing elements of
	// the tank interface
	var t tank
	t = myValue{10,14}
	fmt.Println("Area of the Tank :",t.Area())
	fmt.Println("Volume of the Tank :",t.Volume())
	
}
