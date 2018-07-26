package main

import (
	"fmt"
)

//take an array values then access with index

func main() {

	data := []string {"a", "b", "c", "d", "e"}
	fmt.Println(data[0])
	fmt.Println(data[0:3])
	fmt.Println(data[3:4])
	fmt.Println(data[:])

	/* will print
	a
	[a b c]
	[d]
	[a b c d e]
	 */
	
}
