// https://www.codewars.com/kata/5168bb5dfe9a00b126000018/train/go

package main

import (
	"fmt"
)

func Solution(word string) string {
	var runes = []rune(word)

	for i := 0; i < len(word)/2; i++ {
		temp := runes[i]
		runes[i] = runes[len(word)-i-1]
		runes[len(word)-i-1] = temp
	}

	return string(runes)
}

func main() {
	fmt.Println(Solution("World"))
}
