package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	cards := []string{"first", "second", "third"}
	for _, card := range cards {
		fmt.Println(card)
	}
}
