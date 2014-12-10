package main

import (
	"fmt"
	// "strings"
)

func main() {
	/*
		var direction string
		fmt.Scanf("%s", &direction)

		switch direction = strings.ToLower(direction); direction {
		case "n":
			fmt.Println("Go north!")
		case "w":
			fmt.Println("Go west!")
		case "e":
			fmt.Println("Go east!")
		case "s":
			fmt.Println("Go south!")
		default:
		}
	*/

	var x int
	fmt.Scanf("%d", &x)

	switch {
	case x >= 30:
		fmt.Println("Old!")
	case x >= 18 && x < 30:
		fmt.Println("Young")
	case x < 18:
		fmt.Println("Kid")
		fallthrough // go default then
	default:
		fmt.Println("Sorry, you are too young")
	}
}
