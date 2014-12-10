package main

import "fmt"

//key is unique
func main() {
	m := map[int]int{1: 2, 2: 3}
	fmt.Println(m)
	fmt.Println(m[1])

	m[1] = 15
	fmt.Println(m[1])

	delete(m, 1)
	fmt.Println(m)

	var m2 map[int]string
	m2 = make(map[int]string, 5) //initial 5 cap with make func
	fmt.Println(m2)

	m2[1] = "hello"
	m2[2] = "world"
	m2[1] = "world1"
	fmt.Println(m2)

	m2 = make(map[int]string)
	m2[1] = "hello"
	fmt.Println(m2)
}
