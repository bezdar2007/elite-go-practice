package main

import (
	"fmt"
	"strings"
)

func f(stroka string) int {

	if stroka == "" {
		return 0
	}

	simbols := ",.!?;:'()"
	var result string = stroka
	for i := 0; i < len(simbols); i++ {
		result = strings.ReplaceAll(result, string(simbols[i]), "")
	}
	lowerStroka := strings.ToLower(result)
	stroka1 := strings.Split(lowerStroka, " ")

	hash := make(map[string]int)

	for _, value := range stroka1 {
		val, ok := hash[value]
		if ok == true {
			hash[value] = val + 1
		} else {
			hash[value] = 1
		}
	}

	return len(hash)
}

func main() {
	fmt.Printf("%s -> %v\n", "Hello, world! Hello", f("Hello, world! Hello"))
	fmt.Printf("%s -> %v\n", "", f(""))
	fmt.Printf("%s -> %v\n", "Go is great. Go, go, GO!", f("Go is great. Go, go, GO!"))
}
