package main

import (
	"fmt"
)

func main() {
	var answer string
	fmt.Scanln(&answer)
	if answer != "h" {
	//if answer != "h\n" {
		fmt.Println("I did not understand !! Got: " + answer)
	} else {
		fmt.Println("GOT IT !")
	}
}
