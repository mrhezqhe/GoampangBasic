package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

//take an array values then write to a file

func main() {

	days := []string {"Sunday", "Monday", "Wednesday", "Thursday", "Friday", "Saturday", "Monday"}
	var data bytes.Buffer
	for _, value := range days {
		data.WriteString(value)
		data.WriteString("\n")
	}

	fmt.Print(data.String())
	err := ioutil.WriteFile("/tmp/lesson4.txt", data.Bytes(), 0644)
	check(err)

	
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
