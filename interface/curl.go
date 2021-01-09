package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

//init is call before main
func init()  {
	if len(os.Args) != 2 {
		fmt.Println("Usage : ./example <url>")
		os.Exit(-1)
	}

}

//main is the entry point for the application
func main() {
	//get response form the web
	r,err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	//Copies from the body to Stdout
	io.Copy(os.Stdout,r.Body)
	if err := r.Body.Close() ; err != nil {
		fmt.Println(err)
	}


	//sample program to show how a bytes.Buffer can also be used
	//with the io.Copy function
	var b bytes.Buffer
	//write a string to the buffer
	b.Write([]byte("Hello"))
	//use fprintf to concatenate a string to hte buffer
	fmt.Fprintf(&b,"world !")
	//write the content of the Buffer to stdout
	io.Copy(os.Stdout,&b)

}
