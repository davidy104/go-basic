package main

import (
	"fmt"
	"time"
)

func routine(i int) {
	fmt.Println("routine", i)
}

func main() {
	for i := 0; i < 10; i++ {
		go routine(i)
	}
	time.Sleep(time.Second * 1)
}
