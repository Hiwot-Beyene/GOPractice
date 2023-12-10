package main

import "fmt"
func main() {	
	var welcome string = "Welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")
	}