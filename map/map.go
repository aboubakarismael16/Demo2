package main

import (
	"fmt"
	pretty "github.com/inancgumus/prettyslice"
)

func main() {
	dict := make(map[string]int)
	pretty.Show("map",dict)
	dict1 := map[string]string{"Red":"#da1337","Orange":"#e95a22"}
	pretty.Show("mapping dict1",dict1)

	//declaring a map that stores slices of string
	dict2 := map[int][]string{}
	pretty.Show("map that stores slices of string",dict2)

	//working with maps
	color := map[string]string{}
	color["Red"] = "#da1337"


	// you can create a nil map by just declaring the map
	var colors map[string]string
	pretty.Show("create a nil map by just declaring the map",colors)


	//retrieve the value for hte key "Blue"
	value ,exists := color["Blue"]
	// did key exists
	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("value doesn't exist")
	}

	//another way to retrieve the value for hte key "Blue"
	value = color["blue"]
	if value != "" {
		fmt.Println(value)
	} else {
		fmt.Println("value doesn't exist")
	}


	fmt.Println()
	//create a map of color and color hex codes
	color1 := map[string]string{
		"AliceBlue":"#f0f8ff",
		"Coral":"#ff7F50",
		"DarkGray":"#a9a9a9",
		"ForestGreen":"#228b22",
	}
	for key,value := range color1 {
		fmt.Printf("Key: %s Value: %s\n",key,value)

	}

	fmt.Println()
	fmt.Println("after remove key/value ")
	//passing map between functions
	removeColor(color1,"Coral")
	for key,value := range color1 {
		fmt.Printf("Key: %s Value: %s\n",key,value)

	}

	
}

func removeColor(colors map[string]string,key string)  {
	delete(colors,key)

}