package main

import (
	"fmt"
	"strings"
)

type VowelsFinder  interface {
	FindVowels() []rune
}

type MyString string

func (ms MyString) FindVowels() []rune {
	const vowelsString = "ueoai"
	var vowels []rune
	for _, c := range ms {
		if strings.Contains(vowelsString, string(c)) {
			vowels = append(vowels, c)
		}
	}
	return vowels
}

func main() {
	var i VowelsFinder = MyString("Nguyen Manh")
	vowels := i.FindVowels()
	for i, v := range vowels {
		fmt.Printf("vowels %d is %c \n", i, v)
	}
}
