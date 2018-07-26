package main

import (
	"bufio"
	"os"
	"fmt"
)

//take input as argument then display, until user type "q" will quit

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for text != "q" {
		fmt.Print("enter your text $ ")
		scanner.Scan()
		text = scanner.Text()
		if text != "q" {
			fmt.Println("Your text was <", text, ">")
		}
	}
	
}
