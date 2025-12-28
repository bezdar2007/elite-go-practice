package main

import "fmt"

func compressRepeat(stroka string) {
	hash := make(map[string]int)

	for _, value := range stroka {
		count, ok := hash[string(value)]
		if ok == true {
			hash[string(value)] = count + 1
		} else {
			hash[string(value)] = count
		}
	}

	for key, value := range hash {
		if value >= 0 {
			hash[key] += 1
		}
	}

	for key, value := range hash {
		fmt.Printf("%v%v", key, value)
	}
}

func main() {
	fmt.Printf("%s -> ", "aaabbc")
	compressRepeat("aaabbc")
	fmt.Println()

	fmt.Printf("%s -> ", "abcd")
	compressRepeat("abcd")
	fmt.Println()

	fmt.Printf("%s -> ", "")
	compressRepeat("")
	fmt.Println()

	fmt.Printf("%s -> ", "11122223")
	compressRepeat("11122223")
	fmt.Println()
}
