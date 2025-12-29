package main

import "fmt"

func safeDel(a, b int) (int, bool) {

	if b == 0 {
		return 0, false
	} else {
		result := a / b
		return result, true
	}
}

func main() {
	fmt.Println(safeDel(10, 2))
	fmt.Println(safeDel(5, 0))
	fmt.Println(safeDel(0, 5))
}
