package main

import "fmt"

func main() {
	word := "가나다"

	for _, value := range word {
		fmt.Println(string(value), value)
	}
}
