package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	names := []string{"mario", "luigi", "yoshi"}

	for _, value := range names {
		fmt.Println(value)
	}
}
