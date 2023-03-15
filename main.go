package main

import (
	"fmt"
)

func sayGreeting(n string) {
	fmt.Printf("Hey welcome to narnia %v \n", n)
}
func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}

}
func main() {
	sayGreeting("hello")

	strings := []string{"Hello", "How", "Are", "You"}
	cycleNames(strings, sayGreeting)

}
