package main

import (
	"fmt"
	"unicode/utf8"
)

const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

func main() {
	fmt.Println(sample)
	fmt.Printf("%x\n", sample) // %x (hexadecimal) format
	fmt.Printf("% x\n", sample) // use the "space" flag in that format
	fmt.Printf("%q\n", sample) // quote
	fmt.Printf("%+q\n", sample)

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Println()

	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}

//Go source code is always UTF-8.
//A string holds arbitrary bytes.
//A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
//Those sequences represent Unicode code points, called runes.
//No guarantee is made in Go that characters in strings are normalized.