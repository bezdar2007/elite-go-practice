package main

import "fmt"

func main() {

	var text []byte = []byte("Golang day here")

	fmt.Printf("Первый символ: %v\n", string(text[0]))

	fmt.Printf("Предпоследний символ: %v\n", string(text[len(text)-1]))

	fmt.Printf("Первые пять символов %v\n", string(text[:5]))

	fmt.Printf("Вся строка, кроме последних двух символов %v\n", string(text[:len(text)-2]))

	chet := make([]byte, 0, len(text)/2)
	for i := 0; i < len(text); i++ {
		if i%2 == 0 {
			chet = append(chet, text[i])
		}
	}
	fmt.Printf("Все символы с чётными индексами: %v\n", string(chet))

	nechet := make([]byte, 0, len(text)/2)
	for i := 0; i < len(text); i++ {
		if i%2 != 0 {
			nechet = append(nechet, text[i])
		}
	}
	fmt.Printf("Все символы с нечётными индексами: %v\n", string(nechet))

	reverse := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		reverse[i] = text[len(text)-1-i]
	}
	fmt.Printf("Все символы в обратном порядке: %v\n", string(reverse))

	reverseSpace := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		reverseSpace[i] = text[len(text)-i-1]
	}

	newReverseSpace := make([]byte, 0, len(reverseSpace))
	for j := 0; j < len(reverseSpace); j++ {
		if j%2 == 0 {
			newReverseSpace = append(newReverseSpace, reverseSpace[j])
		}
	}
	fmt.Printf("Все символы строки через один в обратном порядке: %v\n", string(newReverseSpace))

	slice := [][]byte{chet, nechet, reverse, newReverseSpace}
	fmt.Print("Cрез, состоящий из созданных ранее срезов через пробел: ")
	for i := 0; i < len(slice); i++ {
		fmt.Print(string(slice[i]))
	}
	fmt.Println()

	//пустой массив в который добавляем длины каждого среза
	var lenSlices []int
	for _, value := range slice {
		lenSlices = append(lenSlices, len(value))
	}

	//суммируем длины каждого среза
	summ := 0
	for i := 0; i < len(lenSlices); i++ {
		summ += lenSlices[i]
	}
	fmt.Printf("Длина среза, состоящего из созданных ранее срезов через пробел: %v", summ)
}
