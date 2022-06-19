package main

import (
	"fmt"

	"log"

	"rsc.io/quote/v4"

	"playground/greetings"
)

func main() {
	fmt.Println("Bismillah")
	fmt.Println(quote.Go())

	message, err := greetings.Hello("Safi")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
