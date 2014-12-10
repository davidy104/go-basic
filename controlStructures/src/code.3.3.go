package main

import "fmt"

func main() {
	// a := []int{1, 2, 3, 4, 5, 8}
	// var sum int

	// for i := 0; i < len(a); i++ {
	// 	sum += a[i]
	// 	// fmt.Println(a[i])
	// }
	// fmt.Println(sum)

	// var x int
	// for x < 10 { // same as while
	// 	fmt.Scanf("%d", &x)
	// }

	// fmt.Println("Greater than 10:", x)

	var i, x int

	for {
		i++
		if i == 10 {
			break
		}
		if i%2 != 0 {
			x += i
		} else {
			continue
		}
	}

	fmt.Println(x)
}
