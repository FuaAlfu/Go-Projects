package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"

)

var (
	name = kingpin.Arg("name", "Name to greet").Required().String()
	age  = kingpin.Flag("age", "Age of the person to greet").Int()
)

func main() {
	kingpin.Parse()

	fmt.Printf("Hello, %s!", *name)
	if *age > 0 {
		fmt.Printf(" You are %d years old.\n", *age)
	} else {
		fmt.Println()
	}
}
