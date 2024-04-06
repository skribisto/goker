package main

import (
	"fmt"
)

/*
SCORE TESTS
score de flop + turn + river always < score of that + player hand
square should always win with correct high card
AKD9999
AAK9999
AKK9999

Full houses
8877222
8887722
...
*/

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
