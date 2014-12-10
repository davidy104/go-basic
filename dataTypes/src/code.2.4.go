package main 

import "fmt"

func main() {
	//initial slice, slice used to store sequence
	s := []int{1,2,3}
	// fmt.Println(s)
	// fmt.Println(s[1])

	s[1] = 15
	// fmt.Println(s)

	//declare and initial slice
	var s1 []int
	s1 = make([]int, 15, 25) //func make([]T, len, cap) []T
	// fmt.Println(s1)

	fmt.Println(s1, "len is ",len(s1), "cap is:",cap(s1))
	s1 = append(s1,2)
	fmt.Println(s1)
	s1 = append(s1,s...) //append another slice
	fmt.Println(s1)

	s2 := s1[1:5:5] //initial slice from exist one, 5 cap
	fmt.Println(s2)

	copy(s2,s1[14:20]) //cp s1 from index 14 to 20 to  s2
	fmt.Println(s2)

	//array, can not be change at runtime, slice can,
	//array is much efficient then slice, the size has to be initiated when created(memory allocation)
	//slice , allocate memory with make func
	a := [3]int{1,2,3}
	fmt.Println(a)
}