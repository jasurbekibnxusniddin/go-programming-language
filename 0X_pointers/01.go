package main

import "fmt"

func main() {
	var number int

	var ptrNum *int = &number
	
	fmt.Println(ptrNum) 
}