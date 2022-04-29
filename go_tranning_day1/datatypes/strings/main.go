package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//strings type
	mysliceByte := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}
	mysliceRune := []rune{0x0047, 0x0065, 0x0065, 0x006b, 0x0073}
	myString1 := string(mysliceByte)
	myString2 := string(mysliceRune)

	fmt.Println("Create string with slice of byte: ", myString1)
	fmt.Println("Create string with slice of rune: ", myString2)

	myString3 := "ｿｹｯﾄを作成する"
	myString3Runes := []rune(myString3)
	myString4 := "abcdefghi"

	fmt.Println("String1: ", myString3)
	fmt.Println("Length1: ", utf8.RuneCountInString(myString3))
	printBytes(myString3)
	printCharRune(myString3Runes)
	printCharRune(myString3Runes[1:])

	fmt.Println()

	fmt.Println("String2: ", myString4)
	fmt.Println("Length2: ", len(myString4))
	printBytes(myString4)
	printChars(myString4)
	fmt.Println("After substring from 0:", myString4[1:])
}

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()
}

func printCharRune(runes []rune) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Println()
}
