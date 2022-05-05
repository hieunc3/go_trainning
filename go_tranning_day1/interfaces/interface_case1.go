package main

import "fmt"

//create interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

//find vowels in the specific string
//Mystring implements VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}
func exampleInterface1() {
	name := MyString("John Doe")
	//we assign name which is of type Mystring to v of type VowelsFinder
	//this is possible since Mystring implements the Vowelsfinder interface
	var v VowelsFinder = name
	fmt.Printf("Vowels are %c", v.FindVowels())
}
