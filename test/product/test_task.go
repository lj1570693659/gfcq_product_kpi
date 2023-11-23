package main

import "fmt"

func main() {
	twoSendUser := []string{"jiao"}
	User := []string{"lulu"}
	User = append(User, twoSendUser...)
	fmt.Println(User)
}
