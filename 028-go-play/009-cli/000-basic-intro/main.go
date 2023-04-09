package main

import (
	"flag"
	"fmt"

)

func main() {
	// Define command-line arguments
	name := flag.String("name", "World", "A name to say hello to")
	age := flag.Int("age", 0, "An age to print")

	// Parse command-line arguments
	flag.Parse()

	// Print output
	fmt.Printf("Hello %s!\n", *name)
	if *age > 0 {
		fmt.Printf("You are %d years old.\n", *age)
	}
}
