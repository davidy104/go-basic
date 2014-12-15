package main

import "fmt"

type Vehicle struct {
	numberOfWheels int
}

type Truck struct {
	Vehicle
	CarryWeight int
}

func main() {
	t := new(Truck)
	t.numberOfWheels = 1
	fmt.Println(t.numberOfWheels)
}
