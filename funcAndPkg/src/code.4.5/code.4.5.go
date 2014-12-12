//multiple defer would be executed in reverse order
//defer used for resouces management
//making code better with deferred closures
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		return
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}
