package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("C:\\Program Files\\Go\\src\\ReadFiles\\test.txt") // absolute path
	// data,err:=ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
	//Passing the file path as a command line flag
	fptr := flag.String("fpath", "test.txt", "file path to read from") //function returns the address of the string variable that stores the value of the flag
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)
}
