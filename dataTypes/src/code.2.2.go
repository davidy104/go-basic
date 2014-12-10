package main 

import "fmt"

func main() {
	//default = false
	var b1,b2 bool
	b1 = false
	// b2 = true
	fmt.Println(b1,b2)

	var f1 float32
	f1 = 2.34
	fmt.Println(f1)

	var i1 int
	var i2 int64
	var i3 uint

	fmt.Println(i1,i2,i3)

	var r rune
	r = '*'
	fmt.Println(r)

	var s1 string

	b1 = 5 < 10
	fmt.Println(b1)

	f2 := 3 * (5 + 2) + 2.0 / 15
	fmt.Println(f2)

	s2 := "Hello"
	fmt.Println(s2)

	s1 = "World"
	fmt.Println(s1)

	s3 := s2 +" "+s1
	fmt.Println(s3)
}