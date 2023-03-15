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
	return strArr[0], ""
}

func returnStrings(n string, f func(string) (string, string)) (string, string) {
	return f(n)

}
func main() {
	init1, init2 := returnStrings("Connor Nusser", getInitials)
	println(init1)
	println(init2)
	init1, init2 = returnStrings("Connor Busser", getInitials)
	println(init1)
	println(init2)
	init1, init2 = returnStrings("Connousser", getInitials)
	println(init1)
	println(init2)
}
