package main

import (
	"fmt"
)

func sayGreeting(n string) string {
	return fmt.Sprintf("Hey welcome to narnia %v \n", n)
}
func cycleNames(n []string, f func(string) string) string {
	str1 := ""
	for _, value := range n {
		s := f(value)
		str1 += s
	}
	return str1

}
func main() {
	sayGreeting("hello")

	strings := []string{"Hello", "How", "Are", "You"}
	cycleNames(strings, sayGreeting)

}
