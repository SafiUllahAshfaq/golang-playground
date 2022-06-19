package main

import (
	"fmt"

	"rsc.io/quote/v4"

	"playground/greetings"
)

func main() {
	fmt.Println("Bismillah")
	fmt.Println(quote.Go())

	message := greetings.Hello("Safi")
	fmt.Println(message)
}
