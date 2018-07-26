package main

import (
	"fmt"
	"io/ioutil"
	"bytes"
)

//open file and read then append new data to a file

func main() {

	textFile, err := ioutil.ReadFile("/tmp/lesson4.txt")
	check(err)
	fmt.Print(string(textFile))
	fmt.Println("-------------")

	//append new text
	var newData bytes.Buffer
	newData.Write(textFile)
	newData.WriteString("new data is added")
	newData.WriteString("\n")

	err = ioutil.WriteFile("/tmp/lesson4_1.txt", newData.Bytes(), 0644)
	check(err)
	fmt.Print(newData.String())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
