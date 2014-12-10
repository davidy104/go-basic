package main

import "fmt"

func main() {
	var x int
	fmt.Scanf("%d", &x)

	if x >= 30 {
		fmt.Println("Old!")
	} else if x >= 18 && x < 30 {
		fmt.Println("Young!")
	} else {
		fmt.Println("Sorry, you are too young")
	}

}
