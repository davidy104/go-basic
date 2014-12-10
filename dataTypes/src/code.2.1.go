package main 

import "fmt"

var x int

var (
	g int
	s string
	i int64
	y int64 // package or global level
)

const pi = 3.14

const (
	SuperMagicNumber = 1.432
)

func main() {
	// var y int64
	// y = 14567
	y := pi * 2 //if run function, this y is function level
	x = 10
	fmt.Println("Our int variable after assignment:",x,y)


}