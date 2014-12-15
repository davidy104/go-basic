package main

import "fmt"

type Passport struct {
	Photo       []byte
	Name        string
	Surname     string
	DateOfBirth string
}

func main() {
	var p1 Passport
	p2 := Passport{}
	p3 := Passport{
		Photo:       make([]byte, 0, 0),
		Name:        "Rostyslav",
		Surname:     "Dzinko",
		DateOfBirth: "LongAgo",
	}
	fmt.Println(p1, p2, p3)

	p3.DateOfBirth = "Not so long ago"
	fmt.Println(p3.DateOfBirth)

	//note, the difference between pass by value and pointer
	pp3 := &p3
	fmt.Println(pp3)

	pp4 := new(Passport)
	fmt.Println(pp4)
}
