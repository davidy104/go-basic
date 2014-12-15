//behavio
package main

import "fmt"

type Passport struct {
	photo       []byte
	name        string
	surname     string
	dateOfBirth string
}

//Value receiver func (use copy to process, not change the original value)
func (p Passport) GetFullName() string {
	//sprintf return a string, rather than print a string
	return fmt.Sprintf("%s %s", p.name, p.surname)
}

//Pointer receiver func (change the original value)
func (p *Passport) ChangeFullName(name string, surname string) {
	p.name = name
	p.surname = surname
}

func main() {
	p1 := new(Passport)
	p1.ChangeFullName("Go", "Is Awesome")
	fmt.Println(p1.GetFullName())
}
