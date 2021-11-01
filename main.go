package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func read_file(s string) []string {
	// read in file
	content, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	// replace newlines with space and random character and split using the random character
	text := strings.Replace(string(content), "\n", " %", -1)
	words := strings.FieldsFunc(text, func(r rune) bool { return strings.ContainsRune("%", r) })
	return words
}

func ascii_art() {

	art := read_file("ascii-art")
	letters := []rune(os.Args[1])
	for j := 0; j < 8; j++ {
		for i, letter := range letters {
			fmt.Print(art[((int(letter)-32)*9)+j])
			if i == len(letters)-1 {
				fmt.Print("\n")
			}
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		panic("use one arg")
	} else {
		ascii_art()
	}
}
