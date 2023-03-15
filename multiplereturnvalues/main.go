package main

import (
	"strings"
)

func getInitials(name string) (string, string) {
	splitName := strings.Split(name, " ")
	strArr := []string{}
	for _, value := range splitName {
		if len(value) > 1 {
			strArr = append(strArr, value[:2])
		} else {
			strArr = append(strArr, "")
		}
	}
	if len(strArr) > 1 {
		return strArr[0], strArr[1]
	}
	return "", ""
}
func main() {
	init, init2 := getInitials("Connor Nusser")
	println(init)
	println(init2)
}
