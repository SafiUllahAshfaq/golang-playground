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

	// message, err := greetings.Hello("Safi")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(message)

	names := []string{"Safi", "Hassan", "Ali"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
